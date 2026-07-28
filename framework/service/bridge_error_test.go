package gnonative

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"connectrpc.com/connect"
	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
)

// codedConnectError builds the error shape that actually reaches the bridge: a
// connect error carrying ErrDetails, as the service interceptor produces.
func codedConnectError(t *testing.T, code api_gen.ErrCode, message string) error {
	t.Helper()

	err := connect.NewError(connect.CodeUnknown, errors.New("unable to call RPC method abci_query"))
	detail, detailErr := connect.NewErrorDetail(&api_gen.ErrDetails{Code: code, Message: message})
	if detailErr != nil {
		t.Fatalf("NewErrorDetail: %v", detailErr)
	}
	err.AddDetail(detail)
	return err
}

func TestBridgeErrorCarriesCodeAndMessage(t *testing.T) {
	got := bridgeError(codedConnectError(t, api_gen.ErrCode_ErrOutOfGas, "The transaction ran out of gas."), "invoke bridge method error")

	payload, found := strings.CutPrefix(got.Error(), BridgeErrorPrefix)
	if !found {
		t.Fatalf("message %q does not start with %q", got.Error(), BridgeErrorPrefix)
	}

	var envelope bridgeErrorEnvelope
	if err := json.Unmarshal([]byte(payload), &envelope); err != nil {
		t.Fatalf("envelope is not valid JSON: %v", err)
	}

	if envelope.Detail.Code != api_gen.ErrCode_ErrOutOfGas {
		t.Errorf("code = %v, want ErrOutOfGas", envelope.Detail.Code)
	}
	// The wording crosses with the code, so a client need not carry its own.
	if envelope.Detail.Message != "The transaction ran out of gas." {
		t.Errorf("message = %q, want the default for the code", envelope.Detail.Message)
	}
	// The underlying reason still has to be readable for logs.
	if !strings.Contains(envelope.Error, "abci_query") {
		t.Errorf("error = %q, want it to keep the original reason", envelope.Error)
	}
}

// An error with no codes gains nothing from an envelope, and would lose the
// plain message that is all there is to show.
func TestBridgeErrorLeavesUncodedErrorsPlain(t *testing.T) {
	got := bridgeError(errors.New("something else went wrong"), "invoke bridge method error")

	if strings.HasPrefix(got.Error(), BridgeErrorPrefix) {
		t.Errorf("message %q should not be an envelope", got.Error())
	}
	if !strings.Contains(got.Error(), "something else went wrong") {
		t.Errorf("message = %q, want it to keep the original reason", got.Error())
	}
}

// A connect error can arrive without ErrDetails; that is not a code of zero.
func TestBridgeErrorConnectErrorWithoutDetails(t *testing.T) {
	got := bridgeError(connect.NewError(connect.CodeUnavailable, errors.New("no details here")), "invoke bridge method error")

	if strings.HasPrefix(got.Error(), BridgeErrorPrefix) {
		t.Errorf("message %q should not be an envelope", got.Error())
	}
}

// A broadcast result arrives over a stream, so the streaming paths have to
// envelope their failures as well; only the unary one did at first.
func TestBridgeErrorEnvelopesStreamFailures(t *testing.T) {
	got := bridgeError(codedConnectError(t, api_gen.ErrCode_ErrOutOfGas, "The transaction ran out of gas."), "stream receive error")

	payload, found := strings.CutPrefix(got.Error(), BridgeErrorPrefix)
	if !found {
		t.Fatalf("message %q does not start with %q", got.Error(), BridgeErrorPrefix)
	}

	var envelope bridgeErrorEnvelope
	if err := json.Unmarshal([]byte(payload), &envelope); err != nil {
		t.Fatalf("envelope is not valid JSON: %v", err)
	}
	if envelope.Detail.Code != api_gen.ErrCode_ErrOutOfGas {
		t.Errorf("code = %v, want ErrOutOfGas", envelope.Detail.Code)
	}
	// The context still names which bridge call failed.
	if !strings.Contains(envelope.Error, "stream receive error") {
		t.Errorf("error = %q, want it to name the failing call", envelope.Error)
	}
}

// The prefix is written out twice, once here and once in
// expo/src/grpc/bridge_error.ts, and nothing at build time checks that the two
// agree. A silent disagreement stops every envelope being decoded and returns
// clients to reading raw text, so pin the literal: changing it here fails until
// the TypeScript side is changed to match.
func TestBridgeErrorPrefixMatchesTheTypeScriptConstant(t *testing.T) {
	const inTypeScript = "gnonative-error:" // expo/src/grpc/bridge_error.ts

	if BridgeErrorPrefix != inTypeScript {
		t.Errorf("BridgeErrorPrefix = %q, but expo/src/grpc/bridge_error.ts uses %q — update both",
			BridgeErrorPrefix, inTypeScript)
	}
}

func TestErrDetailOfPlainError(t *testing.T) {
	if got := errDetailOf(errors.New("plain")); got != nil {
		t.Errorf("errDetailOf = %v, want nil", got)
	}
}
