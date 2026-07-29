package service

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"syscall"
	"testing"

	abci "github.com/gnolang/gno/tm2/pkg/bft/abci/types"
	tm2errors "github.com/gnolang/gno/tm2/pkg/errors"
	"github.com/gnolang/gno/tm2/pkg/std"
	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
)

// realmRejection builds the error a refused call actually arrives as.
//
// A realm panic is recovered by the VM, degraded to an abci.StringError by
// ABCIErrorOrStringError — the typed error is gone by then, only its text is
// left — and wrapped by gnoclient with the response Log. Written out in full
// because what is being tested is that the classification still finds the
// StringError through that wrapping: tm2's errors.Wrap keeps the cause reachable
// through Unwrap, but prints only the cause, which is why the trace text below
// never reaches anyone.
func realmRejection(reason string) error {
	return tm2errors.Wrapf(abci.StringError(reason),
		"deliver transaction failed: log:%s", "VM panic: "+reason+"\nStacktrace:\n…")
}

// The RPC client hands back transport failures wrapped several layers deep, so
// these cases are built the way the real ones arrive rather than as bare
// sentinel errors: a *url.Error around a *net.OpError around the syscall.
func TestIsRemoteUnreachable(t *testing.T) {
	connRefused := &url.Error{
		Op:  "Post",
		URL: "http://10.0.2.2:26658",
		Err: &net.OpError{
			Op:  "dial",
			Net: "tcp",
			Err: &os.SyscallError{Syscall: "connect", Err: syscall.ECONNREFUSED},
		},
	}

	noSuchHost := &url.Error{
		Op:  "Post",
		URL: "http://nowhere.invalid:26657",
		Err: &net.DNSError{Err: "no such host", Name: "nowhere.invalid", IsNotFound: true},
	}

	timedOut := &url.Error{
		Op:  "Post",
		URL: "http://10.0.2.2:26657",
		Err: &net.OpError{Op: "dial", Net: "tcp", Err: timeoutError{}},
	}

	// A request cut short by its context's deadline. Recognised because the http
	// client hands it back as *url.Error, not because the deadline was matched —
	// see the bare case below.
	deadlineExceeded := &url.Error{
		Op:  "Post",
		URL: "http://10.0.2.2:26657",
		Err: context.DeadlineExceeded,
	}

	for _, tc := range []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"connection refused", connRefused, true},
		{"no such host", noSuchHost, true},
		{"timeout", timedOut, true},
		{"request deadline exceeded", deadlineExceeded, true},
		{"bare ECONNREFUSED", syscall.ECONNREFUSED, true},
		// context.DeadlineExceeded satisfies net.Error with Timeout() true, so
		// this is the case a `net.Error` test would wrongly call unreachable:
		// every handler goes through getGrpcError, keybase ones included, and a
		// deadline that was not a request's says nothing about the node.
		{"bare deadline exceeded", context.DeadlineExceeded, false},
		// A reply from the node is not a transport failure, however it reads.
		{"realm rejection", errors.New("unknown request: no such function"), false},
		{"wrapped realm rejection", fmt.Errorf("call failed: %w", errors.New("out of gas")), false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := isRemoteUnreachable(tc.err); got != tc.want {
				t.Errorf("isRemoteUnreachable(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}

// A refusal the chain could not classify still has to reach the client as
// something better than free text, or every caller is left reading prose to find
// out whether the transaction was executed at all.
func TestGetGrpcErrorClassifiesChainRefusals(t *testing.T) {
	got := getGrpcError(realmRejection("thread body is required"))

	if code := api_gen.FirstCode(got); code != api_gen.ErrCode_ErrChainRejected {
		t.Errorf("code = %v, want ErrChainRejected", code)
	}

	// Classifying must not consume the reason: it is the whole value of this
	// failure, and withErrDetails reads it back out of the error.
	var rejected abci.StringError
	if !errors.As(got, &rejected) {
		t.Fatalf("the abci.StringError is no longer reachable in %v", got)
	}
	if rejected.Error() != "thread body is required" {
		t.Errorf("reason = %q, want the realm's own message", rejected.Error())
	}
}

// A typed chain error keeps its own code: the catch-all is the last branch, not
// a replacement for the ones that say more.
func TestGetGrpcErrorPrefersTypedCodes(t *testing.T) {
	if code := api_gen.FirstCode(getGrpcError(std.ErrOutOfGas("out of gas"))); code != api_gen.ErrCode_ErrOutOfGas {
		t.Errorf("code = %v, want ErrOutOfGas", code)
	}
}

// A failure that never reached the chain is not a refusal by it, however it is
// worded, and must keep the code that says a retry is worth trying.
func TestGetGrpcErrorLeavesTransportFailuresAlone(t *testing.T) {
	unreachable := &url.Error{
		Op:  "Post",
		URL: "http://10.0.2.2:26657",
		Err: &net.OpError{Op: "dial", Net: "tcp", Err: &os.SyscallError{Syscall: "connect", Err: syscall.ECONNREFUSED}},
	}

	if code := api_gen.FirstCode(getGrpcError(unreachable)); code != api_gen.ErrCode_ErrRemoteUnreachable {
		t.Errorf("code = %v, want ErrRemoteUnreachable", code)
	}
}

// Nothing else gains a code it cannot justify: an error from inside this service
// is not the chain refusing anything.
func TestGetGrpcErrorLeavesUnrecognisedErrorsUncoded(t *testing.T) {
	plain := errors.New("something else went wrong")
	if got := getGrpcError(plain); !errors.Is(got, plain) {
		t.Errorf("getGrpcError(%v) = %v, want it unchanged", plain, got)
	}
}

// timeoutError satisfies net.Error reporting a timeout, which is how the http
// client surfaces a deadline reached before any reply.
type timeoutError struct{}

func (timeoutError) Error() string   { return "i/o timeout" }
func (timeoutError) Timeout() bool   { return true }
func (timeoutError) Temporary() bool { return true }
