package service

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
)

// errDetailsInterceptor attaches the ErrCode of a failing call to the response
// as an ErrDetails message.
//
// ErrDetails and ErrCode.Grpc() have been in the schema from the start, but
// nothing ever called them, so the code only ever reached a client inside the
// error *text* — ErrCode.Error() renders as "ErrOutOfGas(#211)". Every client
// therefore had to recover it by matching that shape against a string, which is
// what GRPCError.errCode() does in the TypeScript package. A client in a
// language without a hand-written helper had no way to it at all, even though
// the schema describing it is published to the buf registry.
//
// Attaching the detail here rather than at each return site means it covers
// handlers that never go through getGrpcError, and cannot be forgotten by one
// added later.
//
// It has to wrap streaming handlers as well as unary ones. connect's
// UnaryInterceptorFunc covers only unary calls, and BroadcastTxCommit — whose
// failures are the ones a user is most likely to meet, since they carry the
// result of a transaction — is server-streaming.
type errDetails struct{}

func errDetailsInterceptor() connect.Interceptor { return errDetails{} }

func (errDetails) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		res, err := next(ctx, req)
		return res, withErrDetails(err)
	}
}

func (errDetails) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(ctx context.Context, conn connect.StreamingHandlerConn) error {
		return withErrDetails(next(ctx, conn))
	}
}

// The client side runs in-process here and its errors already carry whatever the
// handler attached, so there is nothing to add.
func (errDetails) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return next
}

// withErrDetails returns err as a *connect.Error carrying its ErrCode and the
// default text for it, or err unchanged when nothing classified it.
//
// The outermost code wins. api_gen.Codes reports the whole chain, but no call
// site wraps one ErrCode in another, so in practice there is exactly one; taking
// the first keeps the classification nearest the caller if that ever changes.
func withErrDetails(err error) error {
	if err == nil {
		return nil
	}

	codes := api_gen.Codes(err)
	if len(codes) == 0 {
		// Nothing classified it, so there is nothing to add. Returning err
		// untouched keeps connect's own errors (deadline, cancellation) intact.
		return err
	}
	code := codes[0]

	// A handler that already built its own connect error keeps it; only the
	// detail is added, so an explicitly chosen status code is not overwritten.
	var connectErr *connect.Error
	if !errors.As(err, &connectErr) {
		connectErr = connect.NewError(connect.CodeUnknown, err)
	}

	for _, existing := range connectErr.Details() {
		if existing.Type() == string((&api_gen.ErrDetails{}).ProtoReflect().Descriptor().FullName()) {
			// Already carried — adding a second would make the client choose.
			return connectErr
		}
	}

	detail, detailErr := connect.NewErrorDetail(&api_gen.ErrDetails{
		Code:    code,
		Message: messageForErrCode(code),
	})
	if detailErr != nil {
		// The error itself still has to reach the client; losing the detail is
		// a degradation, not a reason to drop it.
		return connectErr
	}
	connectErr.AddDetail(detail)

	return connectErr
}
