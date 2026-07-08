package gnonative

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/gnolang/gnonative/v5/service"
)

// The dispatcher provides a connect/gRPC-free path from the mobile bridge to the plain-Go
// service.GnoNativeApi. Requests and responses are encoded with encoding/json in BOTH directions.
// The api package's wire types reproduce the protojson dialect (camelCase, int64 as string, bytes
// as base64 string, zero values omitted), which is what the @gnolang/gnonative TS layer expects.

// unaryHandler decodes a JSON request, calls the plain method, and returns the JSON response bytes.
type unaryHandler func(ctx context.Context, jsonMessage string) ([]byte, error)

// streamHandler decodes a JSON request and drives the plain streaming method, delivering each
// response as JSON bytes through send. It returns when the method returns (nil or error).
type streamHandler func(ctx context.Context, jsonMessage string, send func([]byte) error) error

// unary wraps a plain unary method into a unaryHandler.
func unary[Req any, Res any](fn func(context.Context, *Req) (*Res, error)) unaryHandler {
	return func(ctx context.Context, jsonMessage string) ([]byte, error) {
		req := new(Req)
		if err := json.Unmarshal([]byte(jsonMessage), req); err != nil {
			return nil, errors.Wrap(err, "unable to unmarshal request")
		}
		res, err := fn(ctx, req)
		if err != nil {
			return nil, err
		}
		out, err := json.Marshal(res)
		if err != nil {
			return nil, errors.Wrap(err, "unable to marshal response")
		}
		return out, nil
	}
}

// streaming wraps a plain server-streaming method into a streamHandler.
func streaming[Req any, Res any](fn func(context.Context, *Req, func(*Res) error) error) streamHandler {
	return func(ctx context.Context, jsonMessage string, send func([]byte) error) error {
		req := new(Req)
		if err := json.Unmarshal([]byte(jsonMessage), req); err != nil {
			return errors.Wrap(err, "unable to unmarshal request")
		}
		return fn(ctx, req, func(res *Res) error {
			out, err := json.Marshal(res)
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
