package service

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"syscall"
	"testing"
)

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

	for _, tc := range []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"connection refused", connRefused, true},
		{"no such host", noSuchHost, true},
		{"timeout", timedOut, true},
		{"bare ECONNREFUSED", syscall.ECONNREFUSED, true},
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

// timeoutError satisfies net.Error reporting a timeout, which is how the http
// client surfaces a deadline reached before any reply.
type timeoutError struct{}

func (timeoutError) Error() string   { return "i/o timeout" }
func (timeoutError) Timeout() bool   { return true }
func (timeoutError) Temporary() bool { return true }
