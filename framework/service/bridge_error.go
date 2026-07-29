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
// The mobile bridge can only hand a string to the host platform ("Only the
// reject()'s message argument will be thrown to React Native"), so a connect
// error's details cannot cross as structured data, and the ErrCode would survive
// only as the "ErrOutOfGas(#211)" shape ErrCode.Error() renders. The envelope
// carries the detail instead, and this marks it, so a reader tells one from an
// ordinary message without parsing every rejection as JSON.
//
// No version: the Go framework and the TypeScript package that reads it ship in
// the same npm release, so the two cannot disagree about the format.
const BridgeErrorPrefix = "gnonative-error:"

// bridgeErrorEnvelope is the payload following BridgeErrorPrefix. It mirrors the
// published ErrDetails, so a reader decodes it into its own generated types, and
// keeps the underlying error text alongside for logs.
type bridgeErrorEnvelope struct {
	Detail bridgeErrorDetail `json:"detail"`
	Error  string            `json:"error"`
}

type bridgeErrorDetail struct {
	Code    api_gen.ErrCode `json:"code"`
	Message string          `json:"message"`
}

// bridgeError renders err for the bridge, preserving its ErrDetails when it has
// one. `context` names which bridge call failed, as the plain wrapping did.
//
// Every path that rejects a promise goes through this, not only the unary one: a
// streaming call carries the result of a broadcast.
//
// An error without a code is returned as before — an envelope naming no code
// would cost the plain message and add nothing.
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
		// Losing the detail degrades the message; dropping the error would hide it.
		return wrapped
	}

	return errors.New(BridgeErrorPrefix + string(payload))
}

// errDetailOf reads the ErrDetails a connect error carries, or nil. It has to
// come from the details: the error crossed connect from the in-process server,
// so the typed chain api_gen.Codes walks is gone by now.
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
