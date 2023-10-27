package gnomobiletypes

import (
	"path"

	"github.com/gnolang/gno/tm2/pkg/amino"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/gnolang/gnomobile/service/gnomobiletypes",
	"land.gno.gnomobile.v1",
	path.Join(amino.GetCallersDirname(), "../rpc"),
).WithP3GoPkgPath("github.com/gnolang/gnomobile/service/rpc").
	WithDependencies().WithTypes(
	SetRemoteRequest{},
	SetRemoteResponse{},
	SetChainIDRequest{},
	SetChainIDResponse{},
	SetPasswordRequest{},
	SetPasswordResponse{},
	GenerateRecoveryPhraseRequest{},
	GenerateRecoveryPhraseResponse{},
	KeyInfo{},
	Coin{},
	BaseAccount{},
	ListKeyInfoRequest{},
	ListKeyInfoResponse{},
	GetKeyInfoByNameRequest{},
	GetKeyInfoByNameResponse{},
	GetKeyInfoByAddressRequest{},
	GetKeyInfoByAddressResponse{},
	GetKeyInfoByNameOrAddressRequest{},
	GetKeyInfoByNameOrAddressResponse{},
	CreateAccountRequest{},
	CreateAccountResponse{},
	SelectAccountRequest{},
	SelectAccountResponse{},
	GetActiveAccountRequest{},
	GetActiveAccountResponse{},
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
	CallRequest{},
	CallResponse{},
	AddressToBech32Request{},
	AddressToBech32Response{},
	AddressFromBech32Request{},
	AddressFromBech32Response{},
	HelloRequest{},
	HelloResponse{},
	HelloStreamRequest{},
	HelloStreamResponse{},
).WithComments(path.Join(amino.GetCallersDirname(), "gnomobiletypes.go")))
