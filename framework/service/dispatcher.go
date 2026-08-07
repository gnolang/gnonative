package gnonative

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/gnolang/gnonative/v4/service"
)

// The dispatcher provides a connect/gRPC-free path from the mobile bridge to the plain-Go
// service.GnoNativeApi. Requests and responses are encoded with protojson in BOTH directions
// (camelCase, int64 as string, bytes as base64 string, zero values omitted). This differs from
// the legacy connect path in service_client.go, which decodes requests with protojson but
// encodes responses with encoding/json (snake_case). Never mix the two decoders on one payload.

// unaryHandler decodes a protojson request, calls the plain method, and returns the protojson response bytes.
type unaryHandler func(ctx context.Context, jsonMessage string) ([]byte, error)

// streamHandler decodes a protojson request and drives the plain streaming method, delivering each
// response as protojson bytes through send. It returns when the method returns (nil or error).
type streamHandler func(ctx context.Context, jsonMessage string, send func([]byte) error) error

// unary wraps a plain unary method into a unaryHandler.
func unary[Req any, PReq interface {
	*Req
	proto.Message
}, Res any, PRes interface {
	*Res
	proto.Message
}](fn func(context.Context, PReq) (PRes, error)) unaryHandler {
	return func(ctx context.Context, jsonMessage string) ([]byte, error) {
		req := PReq(new(Req))
		if err := protojson.Unmarshal([]byte(jsonMessage), req); err != nil {
			return nil, errors.Wrap(err, "unable to unmarshal request")
		}
		res, err := fn(ctx, req)
		if err != nil {
			return nil, err
		}
		out, err := protojson.Marshal(res)
		if err != nil {
			return nil, errors.Wrap(err, "unable to marshal response")
		}
		return out, nil
	}
}

// streaming wraps a plain server-streaming method into a streamHandler.
func streaming[Req any, PReq interface {
	*Req
	proto.Message
}, Res any, PRes interface {
	*Res
	proto.Message
}](fn func(context.Context, PReq, func(PRes) error) error) streamHandler {
	return func(ctx context.Context, jsonMessage string, send func([]byte) error) error {
		req := PReq(new(Req))
		if err := protojson.Unmarshal([]byte(jsonMessage), req); err != nil {
			return errors.Wrap(err, "unable to unmarshal request")
		}
		return fn(ctx, req, func(res PRes) error {
			out, err := protojson.Marshal(res)
			if err != nil {
				return errors.Wrap(err, "unable to marshal response")
			}
			return send(out)
		})
	}
}

// newUnaryHandlers returns the map of unary method name -> handler for the given service.
// Keep this in sync with the proto service; the completeness test in dispatcher_test.go guards drift.
func newUnaryHandlers(svc service.GnoNativeApi) map[string]unaryHandler {
	return map[string]unaryHandler{
		"SetRemote":                 unary(svc.SetRemote),
		"GetRemote":                 unary(svc.GetRemote),
		"SetChainID":                unary(svc.SetChainID),
		"GetChainID":                unary(svc.GetChainID),
		"GenerateRecoveryPhrase":    unary(svc.GenerateRecoveryPhrase),
		"ListKeyInfo":               unary(svc.ListKeyInfo),
		"HasKeyByName":              unary(svc.HasKeyByName),
		"HasKeyByAddress":           unary(svc.HasKeyByAddress),
		"HasKeyByNameOrAddress":     unary(svc.HasKeyByNameOrAddress),
		"GetKeyInfoByName":          unary(svc.GetKeyInfoByName),
		"GetKeyInfoByAddress":       unary(svc.GetKeyInfoByAddress),
		"GetKeyInfoByNameOrAddress": unary(svc.GetKeyInfoByNameOrAddress),
		"CreateAccount":             unary(svc.CreateAccount),
		"CreateLedger":              unary(svc.CreateLedger),
		"ActivateAccount":           unary(svc.ActivateAccount),
		"SetPassword":               unary(svc.SetPassword),
		"RenameKey":                 unary(svc.RenameKey),
		"RotatePassword":            unary(svc.RotatePassword),
		"GetActivatedAccount":       unary(svc.GetActivatedAccount),
		"QueryAccount":              unary(svc.QueryAccount),
		"QuerySessionAccount":       unary(svc.QuerySessionAccount),
		"DeleteAccount":             unary(svc.DeleteAccount),
		"Query":                     unary(svc.Query),
		"Render":                    unary(svc.Render),
		"QEval":                     unary(svc.QEval),
		"MakeCallTx":                unary(svc.MakeCallTx),
		"MakeSendTx":                unary(svc.MakeSendTx),
		"MakeRunTx":                 unary(svc.MakeRunTx),
		"MakeCreateSessionTx":       unary(svc.MakeCreateSessionTx),
		"MakeRevokeSessionTx":       unary(svc.MakeRevokeSessionTx),
		"MakeRevokeAllSessionsTx":   unary(svc.MakeRevokeAllSessionsTx),
		"EstimateGas":               unary(svc.EstimateGas),
		"EstimateTxFees":            unary(svc.EstimateTxFees),
		"SignTx":                    unary(svc.SignTx),
		"AddressToBech32":           unary(svc.AddressToBech32),
		"AddressFromBech32":         unary(svc.AddressFromBech32),
		"AddressFromMnemonic":       unary(svc.AddressFromMnemonic),
		"ValidateMnemonicWord":      unary(svc.ValidateMnemonicWord),
		"ValidateMnemonicPhrase":    unary(svc.ValidateMnemonicPhrase),
		"PubKeyBytesFromBech32":     unary(svc.PubKeyBytesFromBech32),
		"Hello":                     unary(svc.Hello),
	}
}

// newStreamHandlers returns the map of server-streaming method name -> handler for the given service.
func newStreamHandlers(svc service.GnoNativeApi) map[string]streamHandler {
	return map[string]streamHandler{
		"Call":              streaming(svc.Call),
		"Send":              streaming(svc.Send),
		"Run":               streaming(svc.Run),
		"CreateSession":     streaming(svc.CreateSession),
		"RevokeSession":     streaming(svc.RevokeSession),
		"RevokeAllSessions": streaming(svc.RevokeAllSessions),
		"BroadcastTxCommit": streaming(svc.BroadcastTxCommit),
		"HelloStream":       streaming(svc.HelloStream),
	}
}
