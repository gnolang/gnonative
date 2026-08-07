package service

import (
	"context"

	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
)

// GnoNativeApi is the plain-Go (connect/gRPC-free) surface of the GnoNative service.
// The connect handlers in api_connect.go and the mobile bridge dispatcher in
// framework/service/ both call these methods. Unary methods return a response and
// error; server-streaming methods take a send callback and return an error.
//
// Keep this interface in sync with the proto service (api/rpc.proto). The dispatcher
// completeness test in framework/service/dispatcher_test.go guards against drift.
type GnoNativeApi interface {
	// Unary
	SetRemote(ctx context.Context, req *api_gen.SetRemoteRequest) (*api_gen.SetRemoteResponse, error)
	GetRemote(ctx context.Context, req *api_gen.GetRemoteRequest) (*api_gen.GetRemoteResponse, error)
	SetChainID(ctx context.Context, req *api_gen.SetChainIDRequest) (*api_gen.SetChainIDResponse, error)
	GetChainID(ctx context.Context, req *api_gen.GetChainIDRequest) (*api_gen.GetChainIDResponse, error)
	GenerateRecoveryPhrase(ctx context.Context, req *api_gen.GenerateRecoveryPhraseRequest) (*api_gen.GenerateRecoveryPhraseResponse, error)
	ListKeyInfo(ctx context.Context, req *api_gen.ListKeyInfoRequest) (*api_gen.ListKeyInfoResponse, error)
	HasKeyByName(ctx context.Context, req *api_gen.HasKeyByNameRequest) (*api_gen.HasKeyByNameResponse, error)
	HasKeyByAddress(ctx context.Context, req *api_gen.HasKeyByAddressRequest) (*api_gen.HasKeyByAddressResponse, error)
	HasKeyByNameOrAddress(ctx context.Context, req *api_gen.HasKeyByNameOrAddressRequest) (*api_gen.HasKeyByNameOrAddressResponse, error)
	GetKeyInfoByName(ctx context.Context, req *api_gen.GetKeyInfoByNameRequest) (*api_gen.GetKeyInfoByNameResponse, error)
	GetKeyInfoByAddress(ctx context.Context, req *api_gen.GetKeyInfoByAddressRequest) (*api_gen.GetKeyInfoByAddressResponse, error)
	GetKeyInfoByNameOrAddress(ctx context.Context, req *api_gen.GetKeyInfoByNameOrAddressRequest) (*api_gen.GetKeyInfoByNameOrAddressResponse, error)
	CreateAccount(ctx context.Context, req *api_gen.CreateAccountRequest) (*api_gen.CreateAccountResponse, error)
	CreateLedger(ctx context.Context, req *api_gen.CreateLedgerRequest) (*api_gen.CreateLedgerResponse, error)
	ActivateAccount(ctx context.Context, req *api_gen.ActivateAccountRequest) (*api_gen.ActivateAccountResponse, error)
	SetPassword(ctx context.Context, req *api_gen.SetPasswordRequest) (*api_gen.SetPasswordResponse, error)
	RenameKey(ctx context.Context, req *api_gen.RenameKeyRequest) (*api_gen.RenameKeyResponse, error)
	RotatePassword(ctx context.Context, req *api_gen.RotatePasswordRequest) (*api_gen.RotatePasswordResponse, error)
	GetActivatedAccount(ctx context.Context, req *api_gen.GetActivatedAccountRequest) (*api_gen.GetActivatedAccountResponse, error)
	QueryAccount(ctx context.Context, req *api_gen.QueryAccountRequest) (*api_gen.QueryAccountResponse, error)
	QuerySessionAccount(ctx context.Context, req *api_gen.QuerySessionAccountRequest) (*api_gen.QuerySessionAccountResponse, error)
	DeleteAccount(ctx context.Context, req *api_gen.DeleteAccountRequest) (*api_gen.DeleteAccountResponse, error)
	Query(ctx context.Context, req *api_gen.QueryRequest) (*api_gen.QueryResponse, error)
	Render(ctx context.Context, req *api_gen.RenderRequest) (*api_gen.RenderResponse, error)
	QEval(ctx context.Context, req *api_gen.QEvalRequest) (*api_gen.QEvalResponse, error)
	MakeCallTx(ctx context.Context, req *api_gen.MakeCallTxRequest) (*api_gen.MakeTxResponse, error)
	MakeSendTx(ctx context.Context, req *api_gen.MakeSendTxRequest) (*api_gen.MakeTxResponse, error)
	MakeRunTx(ctx context.Context, req *api_gen.MakeRunTxRequest) (*api_gen.MakeTxResponse, error)
	MakeCreateSessionTx(ctx context.Context, req *api_gen.CreateSessionRequest) (*api_gen.MakeTxResponse, error)
	MakeRevokeSessionTx(ctx context.Context, req *api_gen.RevokeSessionRequest) (*api_gen.MakeTxResponse, error)
	MakeRevokeAllSessionsTx(ctx context.Context, req *api_gen.RevokeAllSessionsRequest) (*api_gen.MakeTxResponse, error)
	EstimateGas(ctx context.Context, req *api_gen.EstimateGasRequest) (*api_gen.EstimateGasResponse, error)
	EstimateTxFees(ctx context.Context, req *api_gen.EstimateTxFeesRequest) (*api_gen.EstimateTxFeesResponse, error)
	SignTx(ctx context.Context, req *api_gen.SignTxRequest) (*api_gen.SignTxResponse, error)
	AddressToBech32(ctx context.Context, req *api_gen.AddressToBech32Request) (*api_gen.AddressToBech32Response, error)
	AddressFromBech32(ctx context.Context, req *api_gen.AddressFromBech32Request) (*api_gen.AddressFromBech32Response, error)
	AddressFromMnemonic(ctx context.Context, req *api_gen.AddressFromMnemonicRequest) (*api_gen.AddressFromMnemonicResponse, error)
	ValidateMnemonicWord(ctx context.Context, req *api_gen.ValidateMnemonicWordRequest) (*api_gen.ValidateMnemonicWordResponse, error)
	ValidateMnemonicPhrase(ctx context.Context, req *api_gen.ValidateMnemonicPhraseRequest) (*api_gen.ValidateMnemonicPhraseResponse, error)
	PubKeyBytesFromBech32(ctx context.Context, req *api_gen.PubKeyBytesFromBech32Request) (*api_gen.PubKeyBytesFromBech32Response, error)
	Hello(ctx context.Context, req *api_gen.HelloRequest) (*api_gen.HelloResponse, error)

	// Server-streaming
	Call(ctx context.Context, req *api_gen.CallRequest, send func(*api_gen.CallResponse) error) error
	Send(ctx context.Context, req *api_gen.SendRequest, send func(*api_gen.SendResponse) error) error
	Run(ctx context.Context, req *api_gen.RunRequest, send func(*api_gen.RunResponse) error) error
	CreateSession(ctx context.Context, req *api_gen.CreateSessionRequest, send func(*api_gen.CreateSessionResponse) error) error
	RevokeSession(ctx context.Context, req *api_gen.RevokeSessionRequest, send func(*api_gen.RevokeSessionResponse) error) error
	RevokeAllSessions(ctx context.Context, req *api_gen.RevokeAllSessionsRequest, send func(*api_gen.RevokeAllSessionsResponse) error) error
	BroadcastTxCommit(ctx context.Context, req *api_gen.BroadcastTxCommitRequest, send func(*api_gen.BroadcastTxCommitResponse) error) error
	HelloStream(ctx context.Context, req *api_gen.HelloStreamRequest, send func(*api_gen.HelloStreamResponse) error) error
}

var _ GnoNativeApi = (*gnoNativeService)(nil)
