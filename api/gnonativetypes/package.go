package gnonativetypes

import (
	"path"

	"github.com/gnolang/gno/tm2/pkg/amino"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/gnolang/gnonative/api/gnonativetypes",
	"land.gno.gnonative.v1",
	path.Join(amino.GetCallersDirname(), ".."),
).WithP3GoPkgPath("github.com/gnolang/gnonative/api/gen/go").
	WithDependencies().WithTypes(
	SetRemoteRequest{},
	SetRemoteResponse{},
	GetRemoteRequest{},
	GetRemoteResponse{},
	SetChainIDRequest{},
	SetChainIDResponse{},
	GetChainIDRequest{},
	GetChainIDResponse{},
	SetPasswordRequest{},
	SetPasswordResponse{},
	RotatePasswordRequest{},
	RotatePasswordResponse{},
	GenerateRecoveryPhraseRequest{},
	GenerateRecoveryPhraseResponse{},
	KeyInfo{},
	Coin{},
	BaseAccount{},
	ListKeyInfoRequest{},
	ListKeyInfoResponse{},
	HasKeyByNameRequest{},
	HasKeyByNameResponse{},
	HasKeyByAddressRequest{},
	HasKeyByAddressResponse{},
	HasKeyByNameOrAddressRequest{},
	HasKeyByNameOrAddressResponse{},
	GetKeyInfoByNameRequest{},
	GetKeyInfoByNameResponse{},
	GetKeyInfoByAddressRequest{},
	GetKeyInfoByAddressResponse{},
	GetKeyInfoByNameOrAddressRequest{},
	GetKeyInfoByNameOrAddressResponse{},
	CreateAccountRequest{},
	CreateAccountResponse{},
	ActivateAccountRequest{},
	ActivateAccountResponse{},
	GetActivatedAccountRequest{},
	GetActivatedAccountResponse{},
	QueryAccountRequest{},
	QueryAccountResponse{},
	DeleteAccountRequest{},
	DeleteAccountResponse{},
	QueryRequest{},
	QueryResponse{},
	RenderRequest{},
	RenderResponse{},
	QEvalRequest{},
	QEvalResponse{},
	MsgCall{},
	CallRequest{},
	CallResponse{},
	MsgSend{},
	SendRequest{},
	SendResponse{},
	MsgRun{},
	RunRequest{},
	RunResponse{},
	MakeTxResponse{},
	SignTxRequest{},
	SignTxResponse{},
	EstimateGasRequest{},
	EstimateGasResponse{},
	BroadcastTxCommitRequest{},
	BroadcastTxCommitResponse{},
	AddressToBech32Request{},
	AddressToBech32Response{},
	AddressFromBech32Request{},
	AddressFromBech32Response{},
	AddressFromMnemonicRequest{},
	AddressFromMnemonicResponse{},
	HelloRequest{},
	HelloResponse{},
	HelloStreamRequest{},
	HelloStreamResponse{},
).WithComments(path.Join(amino.GetCallersDirname(), "gnonativetypes.go")))
