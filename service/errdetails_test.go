package service

import (
	"context"
	"errors"
	"testing"

	"connectrpc.com/connect"
	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
)

func detailOf(t *testing.T, err error) *api_gen.ErrDetails {
	t.Helper()

	var connectErr *connect.Error
	if !errors.As(err, &connectErr) {
		t.Fatalf("expected a *connect.Error, got %T", err)
	}

	for _, detail := range connectErr.Details() {
		value, valueErr := detail.Value()
		if valueErr != nil {
			t.Fatalf("detail.Value(): %v", valueErr)
		}
		if details, ok := value.(*api_gen.ErrDetails); ok {
			return details
		}
	}
	return nil
}

func TestWithErrDetailsAttachesCodeAndMessage(t *testing.T) {
	got := detailOf(t, withErrDetails(api_gen.ErrCode_ErrOutOfGas))

	if got.Code != api_gen.ErrCode_ErrOutOfGas {
		t.Errorf("code = %v, want ErrOutOfGas", got.Code)
	}
	// The wording travels with the code so no client has to invent it.
	if got.Message != errCodeMessages[api_gen.ErrCode_ErrOutOfGas] {
		t.Errorf("message = %q, want the default for the code", got.Message)
	}
	if got.Message == "" {
		t.Error("ErrOutOfGas should carry a default message")
	}
}

// A code only a developer can act on carries no invented text, and the empty
// message is how a client tells that apart from a real one.
func TestWithErrDetailsLeavesDeveloperCodesUnworded(t *testing.T) {
	got := detailOf(t, withErrDetails(api_gen.ErrCode_ErrNotImplemented))

	if got.Code != api_gen.ErrCode_ErrNotImplemented {
		t.Errorf("code = %v, want ErrNotImplemented", got.Code)
	}
	if got.Message != "" {
		t.Errorf("message = %q, want empty", got.Message)
	}
}

// No call site nests coded errors today, but if one ever does the outermost
// classification is the one nearest the caller, so it is the one reported.
func TestWithErrDetailsReportsTheOutermostCode(t *testing.T) {
	inner := api_gen.ErrCode_ErrUnauthorized.Wrap(errors.New("boom"))
	got := detailOf(t, withErrDetails(api_gen.ErrCode_ErrTxDecode.Wrap(inner)))

	if got.Code != api_gen.ErrCode_ErrTxDecode {
		t.Errorf("code = %v, want ErrTxDecode", got.Code)
	}
}

// An error nothing classified has no code to add, and must not be dressed up as
// a connect error just to carry an empty detail.
func TestWithErrDetailsLeavesUncodedErrorsAlone(t *testing.T) {
	plain := errors.New("something else went wrong")
	if got := withErrDetails(plain); !errors.Is(got, plain) {
		t.Errorf("withErrDetails(%v) = %v, want it unchanged", plain, got)
	}
	var connectErr *connect.Error
	if errors.As(withErrDetails(plain), &connectErr) {
		t.Error("an uncoded error was converted into a *connect.Error")
	}
}

// A handler that chose its own status code keeps it: only the detail is added.
func TestWithErrDetailsPreservesAnExplicitCode(t *testing.T) {
	original := connect.NewError(connect.CodePermissionDenied, api_gen.ErrCode_ErrUnauthorized)

	var connectErr *connect.Error
	if !errors.As(withErrDetails(original), &connectErr) {
		t.Fatal("expected a *connect.Error")
	}
	if connectErr.Code() != connect.CodePermissionDenied {
		t.Errorf("code = %v, want CodePermissionDenied", connectErr.Code())
	}
	if got := detailOf(t, connectErr); got.Code != api_gen.ErrCode_ErrUnauthorized {
		t.Errorf("code = %v, want ErrUnauthorized", got.Code)
	}
}

// Running twice must not leave the client choosing between two ErrDetails.
func TestWithErrDetailsIsIdempotent(t *testing.T) {
	once := withErrDetails(api_gen.ErrCode_ErrOutOfGas)
	twice := withErrDetails(once)

	var connectErr *connect.Error
	if !errors.As(twice, &connectErr) {
		t.Fatal("expected a *connect.Error")
	}

	details := 0
	for _, detail := range connectErr.Details() {
		value, valueErr := detail.Value()
		if valueErr != nil {
			continue
		}
		if _, ok := value.(*api_gen.ErrDetails); ok {
			details++
		}
	}
	if details != 1 {
		t.Errorf("ErrDetails count = %d, want 1", details)
	}
}

// BroadcastTxCommit is server-streaming, so a unary-only interceptor would never
// see the failure that carries a transaction's result.
func TestInterceptorWrapsStreamingHandlers(t *testing.T) {
	var interceptor connect.Interceptor = errDetailsInterceptor()

	wrapped := interceptor.WrapStreamingHandler(
		func(context.Context, connect.StreamingHandlerConn) error { return api_gen.ErrCode_ErrOutOfGas },
	)

	got := detailOf(t, wrapped(context.Background(), nil))
	if got.Code != api_gen.ErrCode_ErrOutOfGas {
		t.Errorf("code = %v, want ErrOutOfGas", got.Code)
	}
	if got.Message == "" {
		t.Error("a streaming failure should carry its default message too")
	}
}

func TestWithErrDetailsNil(t *testing.T) {
	if got := withErrDetails(nil); got != nil {
		t.Errorf("withErrDetails(nil) = %v, want nil", got)
	}
}
