// This file implements thin connect wrappers over the plain-Go GnoNativeApi methods in api.go.
// Registering connectHandler (instead of *gnoNativeService directly) keeps the gRPC/connect wire
// behavior identical while the business logic lives in connect-free functions.

package service

import (
	"context"

	"connectrpc.com/connect"

	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
	"github.com/gnolang/gnonative/v4/api/gen/go/_goconnect"
)

type connectHandler struct{ svc *gnoNativeService }

var _ _goconnect.GnoNativeServiceHandler = (*connectHandler)(nil)

func (h *connectHandler) SetRemote(ctx context.Context, req *connect.Request[api_gen.SetRemoteRequest]) (*connect.Response[api_gen.SetRemoteResponse], error) {
	res, err := h.svc.SetRemote(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) GetRemote(ctx context.Context, req *connect.Request[api_gen.GetRemoteRequest]) (*connect.Response[api_gen.GetRemoteResponse], error) {
	res, err := h.svc.GetRemote(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) SetChainID(ctx context.Context, req *connect.Request[api_gen.SetChainIDRequest]) (*connect.Response[api_gen.SetChainIDResponse], error) {
	res, err := h.svc.SetChainID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) GetChainID(ctx context.Context, req *connect.Request[api_gen.GetChainIDRequest]) (*connect.Response[api_gen.GetChainIDResponse], error) {
	res, err := h.svc.GetChainID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) GenerateRecoveryPhrase(ctx context.Context, req *connect.Request[api_gen.GenerateRecoveryPhraseRequest]) (*connect.Response[api_gen.GenerateRecoveryPhraseResponse], error) {
	res, err := h.svc.GenerateRecoveryPhrase(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) ListKeyInfo(ctx context.Context, req *connect.Request[api_gen.ListKeyInfoRequest]) (*connect.Response[api_gen.ListKeyInfoResponse], error) {
	res, err := h.svc.ListKeyInfo(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) HasKeyByName(ctx context.Context, req *connect.Request[api_gen.HasKeyByNameRequest]) (*connect.Response[api_gen.HasKeyByNameResponse], error) {
	res, err := h.svc.HasKeyByName(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) HasKeyByAddress(ctx context.Context, req *connect.Request[api_gen.HasKeyByAddressRequest]) (*connect.Response[api_gen.HasKeyByAddressResponse], error) {
	res, err := h.svc.HasKeyByAddress(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) HasKeyByNameOrAddress(ctx context.Context, req *connect.Request[api_gen.HasKeyByNameOrAddressRequest]) (*connect.Response[api_gen.HasKeyByNameOrAddressResponse], error) {
	res, err := h.svc.HasKeyByNameOrAddress(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) GetKeyInfoByName(ctx context.Context, req *connect.Request[api_gen.GetKeyInfoByNameRequest]) (*connect.Response[api_gen.GetKeyInfoByNameResponse], error) {
	res, err := h.svc.GetKeyInfoByName(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) GetKeyInfoByAddress(ctx context.Context, req *connect.Request[api_gen.GetKeyInfoByAddressRequest]) (*connect.Response[api_gen.GetKeyInfoByAddressResponse], error) {
	res, err := h.svc.GetKeyInfoByAddress(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) GetKeyInfoByNameOrAddress(ctx context.Context, req *connect.Request[api_gen.GetKeyInfoByNameOrAddressRequest]) (*connect.Response[api_gen.GetKeyInfoByNameOrAddressResponse], error) {
	res, err := h.svc.GetKeyInfoByNameOrAddress(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) CreateAccount(ctx context.Context, req *connect.Request[api_gen.CreateAccountRequest]) (*connect.Response[api_gen.CreateAccountResponse], error) {
	res, err := h.svc.CreateAccount(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) CreateLedger(ctx context.Context, req *connect.Request[api_gen.CreateLedgerRequest]) (*connect.Response[api_gen.CreateLedgerResponse], error) {
	res, err := h.svc.CreateLedger(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) ActivateAccount(ctx context.Context, req *connect.Request[api_gen.ActivateAccountRequest]) (*connect.Response[api_gen.ActivateAccountResponse], error) {
	res, err := h.svc.ActivateAccount(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) SetPassword(ctx context.Context, req *connect.Request[api_gen.SetPasswordRequest]) (*connect.Response[api_gen.SetPasswordResponse], error) {
	res, err := h.svc.SetPassword(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) RenameKey(ctx context.Context, req *connect.Request[api_gen.RenameKeyRequest]) (*connect.Response[api_gen.RenameKeyResponse], error) {
	res, err := h.svc.RenameKey(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) RotatePassword(ctx context.Context, req *connect.Request[api_gen.RotatePasswordRequest]) (*connect.Response[api_gen.RotatePasswordResponse], error) {
	res, err := h.svc.RotatePassword(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) GetActivatedAccount(ctx context.Context, req *connect.Request[api_gen.GetActivatedAccountRequest]) (*connect.Response[api_gen.GetActivatedAccountResponse], error) {
	res, err := h.svc.GetActivatedAccount(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) QueryAccount(ctx context.Context, req *connect.Request[api_gen.QueryAccountRequest]) (*connect.Response[api_gen.QueryAccountResponse], error) {
	res, err := h.svc.QueryAccount(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) QuerySessionAccount(ctx context.Context, req *connect.Request[api_gen.QuerySessionAccountRequest]) (*connect.Response[api_gen.QuerySessionAccountResponse], error) {
	res, err := h.svc.QuerySessionAccount(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) DeleteAccount(ctx context.Context, req *connect.Request[api_gen.DeleteAccountRequest]) (*connect.Response[api_gen.DeleteAccountResponse], error) {
	res, err := h.svc.DeleteAccount(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) Query(ctx context.Context, req *connect.Request[api_gen.QueryRequest]) (*connect.Response[api_gen.QueryResponse], error) {
	res, err := h.svc.Query(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) Render(ctx context.Context, req *connect.Request[api_gen.RenderRequest]) (*connect.Response[api_gen.RenderResponse], error) {
	res, err := h.svc.Render(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) QEval(ctx context.Context, req *connect.Request[api_gen.QEvalRequest]) (*connect.Response[api_gen.QEvalResponse], error) {
	res, err := h.svc.QEval(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) MakeCallTx(ctx context.Context, req *connect.Request[api_gen.MakeCallTxRequest]) (*connect.Response[api_gen.MakeTxResponse], error) {
	res, err := h.svc.MakeCallTx(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) MakeSendTx(ctx context.Context, req *connect.Request[api_gen.MakeSendTxRequest]) (*connect.Response[api_gen.MakeTxResponse], error) {
	res, err := h.svc.MakeSendTx(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) MakeRunTx(ctx context.Context, req *connect.Request[api_gen.MakeRunTxRequest]) (*connect.Response[api_gen.MakeTxResponse], error) {
	res, err := h.svc.MakeRunTx(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) MakeCreateSessionTx(ctx context.Context, req *connect.Request[api_gen.CreateSessionRequest]) (*connect.Response[api_gen.MakeTxResponse], error) {
	res, err := h.svc.MakeCreateSessionTx(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) MakeRevokeSessionTx(ctx context.Context, req *connect.Request[api_gen.RevokeSessionRequest]) (*connect.Response[api_gen.MakeTxResponse], error) {
	res, err := h.svc.MakeRevokeSessionTx(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) MakeRevokeAllSessionsTx(ctx context.Context, req *connect.Request[api_gen.RevokeAllSessionsRequest]) (*connect.Response[api_gen.MakeTxResponse], error) {
	res, err := h.svc.MakeRevokeAllSessionsTx(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) EstimateGas(ctx context.Context, req *connect.Request[api_gen.EstimateGasRequest]) (*connect.Response[api_gen.EstimateGasResponse], error) {
	res, err := h.svc.EstimateGas(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) EstimateTxFees(ctx context.Context, req *connect.Request[api_gen.EstimateTxFeesRequest]) (*connect.Response[api_gen.EstimateTxFeesResponse], error) {
	res, err := h.svc.EstimateTxFees(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) SignTx(ctx context.Context, req *connect.Request[api_gen.SignTxRequest]) (*connect.Response[api_gen.SignTxResponse], error) {
	res, err := h.svc.SignTx(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) AddressToBech32(ctx context.Context, req *connect.Request[api_gen.AddressToBech32Request]) (*connect.Response[api_gen.AddressToBech32Response], error) {
	res, err := h.svc.AddressToBech32(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) AddressFromBech32(ctx context.Context, req *connect.Request[api_gen.AddressFromBech32Request]) (*connect.Response[api_gen.AddressFromBech32Response], error) {
	res, err := h.svc.AddressFromBech32(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) AddressFromMnemonic(ctx context.Context, req *connect.Request[api_gen.AddressFromMnemonicRequest]) (*connect.Response[api_gen.AddressFromMnemonicResponse], error) {
	res, err := h.svc.AddressFromMnemonic(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) ValidateMnemonicWord(ctx context.Context, req *connect.Request[api_gen.ValidateMnemonicWordRequest]) (*connect.Response[api_gen.ValidateMnemonicWordResponse], error) {
	res, err := h.svc.ValidateMnemonicWord(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) ValidateMnemonicPhrase(ctx context.Context, req *connect.Request[api_gen.ValidateMnemonicPhraseRequest]) (*connect.Response[api_gen.ValidateMnemonicPhraseResponse], error) {
	res, err := h.svc.ValidateMnemonicPhrase(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) PubKeyBytesFromBech32(ctx context.Context, req *connect.Request[api_gen.PubKeyBytesFromBech32Request]) (*connect.Response[api_gen.PubKeyBytesFromBech32Response], error) {
	res, err := h.svc.PubKeyBytesFromBech32(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) Hello(ctx context.Context, req *connect.Request[api_gen.HelloRequest]) (*connect.Response[api_gen.HelloResponse], error) {
	res, err := h.svc.Hello(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h *connectHandler) Call(ctx context.Context, req *connect.Request[api_gen.CallRequest], stream *connect.ServerStream[api_gen.CallResponse]) error {
	return h.svc.Call(ctx, req.Msg, stream.Send)
}

func (h *connectHandler) Send(ctx context.Context, req *connect.Request[api_gen.SendRequest], stream *connect.ServerStream[api_gen.SendResponse]) error {
	return h.svc.Send(ctx, req.Msg, stream.Send)
}

func (h *connectHandler) Run(ctx context.Context, req *connect.Request[api_gen.RunRequest], stream *connect.ServerStream[api_gen.RunResponse]) error {
	return h.svc.Run(ctx, req.Msg, stream.Send)
}

func (h *connectHandler) CreateSession(ctx context.Context, req *connect.Request[api_gen.CreateSessionRequest], stream *connect.ServerStream[api_gen.CreateSessionResponse]) error {
	return h.svc.CreateSession(ctx, req.Msg, stream.Send)
}

func (h *connectHandler) RevokeSession(ctx context.Context, req *connect.Request[api_gen.RevokeSessionRequest], stream *connect.ServerStream[api_gen.RevokeSessionResponse]) error {
	return h.svc.RevokeSession(ctx, req.Msg, stream.Send)
}

func (h *connectHandler) RevokeAllSessions(ctx context.Context, req *connect.Request[api_gen.RevokeAllSessionsRequest], stream *connect.ServerStream[api_gen.RevokeAllSessionsResponse]) error {
	return h.svc.RevokeAllSessions(ctx, req.Msg, stream.Send)
}

func (h *connectHandler) BroadcastTxCommit(ctx context.Context, req *connect.Request[api_gen.BroadcastTxCommitRequest], stream *connect.ServerStream[api_gen.BroadcastTxCommitResponse]) error {
	return h.svc.BroadcastTxCommit(ctx, req.Msg, stream.Send)
}

func (h *connectHandler) HelloStream(ctx context.Context, req *connect.Request[api_gen.HelloStreamRequest], stream *connect.ServerStream[api_gen.HelloStreamResponse]) error {
	return h.svc.HelloStream(ctx, req.Msg, stream.Send)
}
