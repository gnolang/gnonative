package gnonative

import (
	"encoding/json"
	"errors"
	"os"
	"regexp"
	"strings"
	"testing"

	"connectrpc.com/connect"
	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
)

// envelopeOf decodes what the bridge would hand the host platform, failing the
// test if the rejection is not an envelope at all.
func envelopeOf(t *testing.T, err error) bridgeErrorEnvelope {
	t.Helper()

	payload, found := strings.CutPrefix(err.Error(), BridgeErrorPrefix)
	if !found {
		t.Fatalf("message %q does not start with %q", err.Error(), BridgeErrorPrefix)
	}

	var envelope bridgeErrorEnvelope
	if unmarshalErr := json.Unmarshal([]byte(payload), &envelope); unmarshalErr != nil {
		t.Fatalf("envelope is not valid JSON: %v", unmarshalErr)
	}
	return envelope
}

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
	envelope := envelopeOf(t, bridgeError(codedConnectError(t, api_gen.ErrCode_ErrOutOfGas, "The transaction ran out of gas."), "invoke bridge method error"))

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
	envelope := envelopeOf(t, bridgeError(codedConnectError(t, api_gen.ErrCode_ErrOutOfGas, "The transaction ran out of gas."), "stream receive error"))

	if envelope.Detail.Code != api_gen.ErrCode_ErrOutOfGas {
		t.Errorf("code = %v, want ErrOutOfGas", envelope.Detail.Code)
	}
	// The context still names which bridge call failed.
	if !strings.Contains(envelope.Error, "stream receive error") {
		t.Errorf("error = %q, want it to name the failing call", envelope.Error)
	}
}

// The status the handler chose is carried too, so a bridge client branches on
// the same code a networked one would instead of seeing Unknown everywhere.
func TestBridgeErrorCarriesTheConnectCode(t *testing.T) {
	err := connect.NewError(connect.CodePermissionDenied, api_gen.ErrCode_ErrUnauthorized)
	detail, detailErr := connect.NewErrorDetail(&api_gen.ErrDetails{Code: api_gen.ErrCode_ErrUnauthorized})
	if detailErr != nil {
		t.Fatalf("NewErrorDetail: %v", detailErr)
	}
	err.AddDetail(detail)

	envelope := envelopeOf(t, bridgeError(err, "invoke bridge method error"))

	if envelope.ConnectCode != int(connect.CodePermissionDenied) {
		t.Errorf("connectCode = %d, want %d", envelope.ConnectCode, connect.CodePermissionDenied)
	}
}

// The prefix is written out twice, once here and once in
// expo/src/grpc/bridge_error.ts, and nothing at build time checks that the two
// agree. A silent disagreement stops every envelope being decoded and returns
// clients to reading raw text.
//
// So this is the check: the constant here is the source of truth, and the
// TypeScript declaration is read out of the file and compared against it. A
// change on either side fails until the other follows. Matching the declaration
// with a pattern rather than the whole line keeps reformatting — quotes, spacing
// — from failing the test for no reason.
func TestTypeScriptDeclaresTheSameBridgeErrorPrefix(t *testing.T) {
	const path = "../../expo/src/grpc/bridge_error.ts"

	source, err := os.ReadFile(path)
	if err != nil {
		// Not skipped: a guard that quietly stops running is the thing being
		// guarded against.
		t.Fatalf("reading %s: %v", path, err)
	}

	declaration := regexp.MustCompile(`BRIDGE_ERROR_PREFIX\s*=\s*['"]([^'"]*)['"]`).FindSubmatch(source)
	if declaration == nil {
		t.Fatalf("%s no longer declares BRIDGE_ERROR_PREFIX — the envelope marker has to be shared with it", path)
	}

	if inTypeScript := string(declaration[1]); inTypeScript != BridgeErrorPrefix {
		t.Errorf("BridgeErrorPrefix = %q, but %s declares %q — update both", BridgeErrorPrefix, path, inTypeScript)
	}
}

func TestErrDetailOfConnectErrorWithoutDetails(t *testing.T) {
	if got := errDetailOf(connect.NewError(connect.CodeUnknown, errors.New("plain"))); got != nil {
		t.Errorf("errDetailOf = %v, want nil", got)
	}
}
