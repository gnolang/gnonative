package service

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
)

const testProcedure = "/land.gno.gnonative.v1.Test/Fail"

// The detail is only useful if it survives the wire, and a server-streaming call
// carries it differently from a unary one — it goes in the end-of-stream frame
// rather than the response trailers. BroadcastTxCommit is server-streaming, so
// that is the path a failed transaction takes.
//
// ErrDetails doubles as the request and response type here; any proto message
// would do.
func TestErrDetailsSurvivesServerStreaming(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle(testProcedure, connect.NewServerStreamHandler(
		testProcedure,
		func(context.Context, *connect.Request[api_gen.ErrDetails], *connect.ServerStream[api_gen.ErrDetails]) error {
			return api_gen.ErrCode_ErrOutOfGas
		},
		connect.WithInterceptors(errDetailsInterceptor()),
	))

	server := httptest.NewServer(mux)
	defer server.Close()

	client := connect.NewClient[api_gen.ErrDetails, api_gen.ErrDetails](server.Client(), server.URL+testProcedure)

	stream, err := client.CallServerStream(context.Background(), connect.NewRequest(&api_gen.ErrDetails{}))
	if err != nil {
		t.Fatalf("CallServerStream: %v", err)
	}
	defer stream.Close()

	if stream.Receive() {
		t.Fatal("expected the stream to fail rather than yield a message")
	}

	got := detailOf(t, stream.Err())
	if got.Code != api_gen.ErrCode_ErrOutOfGas {
		t.Errorf("code = %v, want ErrOutOfGas", got.Code)
	}
	if got.Message != errCodeMessages[api_gen.ErrCode_ErrOutOfGas] {
		t.Errorf("message = %q, want the default for the code", got.Message)
	}
}

// The same for a unary call, so a regression in either shape is caught here
// rather than on a device.
func TestErrDetailsSurvivesUnary(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle(testProcedure, connect.NewUnaryHandler(
		testProcedure,
		func(context.Context, *connect.Request[api_gen.ErrDetails]) (*connect.Response[api_gen.ErrDetails], error) {
			return nil, api_gen.ErrCode_ErrUnauthorized
		},
		connect.WithInterceptors(errDetailsInterceptor()),
	))

	server := httptest.NewServer(mux)
	defer server.Close()

	client := connect.NewClient[api_gen.ErrDetails, api_gen.ErrDetails](server.Client(), server.URL+testProcedure)

	_, err := client.CallUnary(context.Background(), connect.NewRequest(&api_gen.ErrDetails{}))
	if err == nil {
		t.Fatal("expected the call to fail")
	}

	if got := detailOf(t, err); got.Code != api_gen.ErrCode_ErrUnauthorized {
		t.Errorf("code = %v, want ErrUnauthorized", got.Code)
	}
}
