package service

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
)

// errDetailsInterceptor attaches the ErrCode of a failing call to the response
// as an ErrDetails message, so a client reads the code instead of matching the
// "ErrOutOfGas(#211)" shape ErrCode.Error() renders in the message text.
//
// Attached here rather than at each return site: it covers handlers that never
// go through getGrpcError, and cannot be forgotten by one added later.
//
// It implements connect.Interceptor rather than using UnaryInterceptorFunc,
// which covers unary calls only: BroadcastTxCommit is server-streaming, and its
// failures carry the result of a transaction.
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
// text for it, or err unchanged when nothing classified it.
//
// The outermost code wins: no call site nests ErrCodes today, so taking the
// first keeps the classification nearest the caller if one ever does.
func withErrDetails(err error) error {
	if err == nil {
		return nil
	}

	codes := api_gen.Codes(err)
	if len(codes) == 0 {
		// Untouched, so connect's own errors (deadline, cancellation) survive.
		return err
	}
	code := codes[0]

	// A handler that built its own connect error keeps it, status code included.
	var connectErr *connect.Error
	if !errors.As(err, &connectErr) {
		connectErr = connect.NewError(connect.CodeUnknown, err)
	}

	for _, existing := range connectErr.Details() {
		if existing.Type() == string((&api_gen.ErrDetails{}).ProtoReflect().Descriptor().FullName()) {
			// Already carried — a second would make the client choose.
			return connectErr
		}
	}

	detail, detailErr := connect.NewErrorDetail(&api_gen.ErrDetails{
		Code:    code,
		Message: messageForError(err, code),
	})
	if detailErr != nil {
		// Losing the detail degrades the failure; dropping it would hide it.
		return connectErr
	}
	connectErr.AddDetail(detail)

	return connectErr
}
