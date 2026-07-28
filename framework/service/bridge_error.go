package gnonative

import (
	"encoding/json"
	stderrors "errors"

	"connectrpc.com/connect"
	"github.com/pkg/errors"

	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
)

// BridgeErrorPrefix marks a rejection message that carries a JSON envelope
// rather than plain text.
//
// The mobile bridge can only hand a string to the host platform — the Android
// side says so outright ("Only the reject()'s message argument will be thrown to
// React Native") — so a connect error's details cannot cross it as structured
// data. Without them the ErrCode survives only as the "ErrOutOfGas(#211)" shape
// that ErrCode.Error() happens to render, leaving every client to recover it by
// matching a substring.
//
// The envelope carries the detail explicitly instead, and this marks it: a
// reader must be able to tell an envelope from an ordinary message without
// speculatively parsing every rejection as JSON.
//
// It carries no version. The Go framework and the TypeScript package that reads
// it ship in the same npm release, so the two can never disagree about the
// format; a version would be a number that is never once observed to differ.
const BridgeErrorPrefix = "gnonative-error:"

// bridgeErrorEnvelope is the payload following BridgeErrorPrefix.
//
// It mirrors the ErrDetails published to the buf registry, so a reader decodes
// it into its own generated types rather than a hand-copied shape, and keeps the
// underlying error text alongside for logs.
type bridgeErrorEnvelope struct {
	Detail bridgeErrorDetail `json:"detail"`
	Error  string            `json:"error"`
}

type bridgeErrorDetail struct {
	Code    api_gen.ErrCode `json:"code"`
	Message string          `json:"message"`
}

// bridgeError renders err for the bridge, preserving its ErrDetails when it has
// one. `context` describes which bridge call failed, as the plain wrapping did.
//
// Every path that rejects a promise has to go through this, not just the unary
// one: a streaming call carries the result of a broadcast, so its failures are
// the ones a user is most likely to see.
//
// An error without a code is returned as before: an envelope naming no code
// would tell a reader nothing it could not already see, and would cost the plain
// message that is all there is to show.
func bridgeError(err error, context string) error {
	wrapped := errors.Wrap(err, context)

	detail := errDetailOf(err)
	if detail == nil {
		return wrapped
	}

	payload, marshalErr := json.Marshal(bridgeErrorEnvelope{
		Detail: bridgeErrorDetail{Code: detail.Code, Message: detail.Message},
		Error:  wrapped.Error(),
	})
	if marshalErr != nil {
		// The failure still has to reach the caller; losing the detail degrades
		// the message, dropping the error entirely would hide the problem.
		return wrapped
	}

	return errors.New(BridgeErrorPrefix + string(payload))
}

// errDetailOf reads the ErrDetails a connect error carries, or nil.
//
// It has to come from the details rather than from the Go error chain: this
// error came back over connect from the in-process server, so the typed chain
// api_gen.Codes walks is gone by the time it arrives here.
func errDetailOf(err error) *api_gen.ErrDetails {
	var connectErr *connect.Error
	if !stderrors.As(err, &connectErr) {
		return nil
	}

	for _, detail := range connectErr.Details() {
		value, valueErr := detail.Value()
		if valueErr != nil {
			continue
		}
		if details, ok := value.(*api_gen.ErrDetails); ok {
			return details
		}
	}

	return nil
}
