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
	DeleteAccountRequest{},
	DeleteAccountResponse{},
	QueryRequest{},
	QueryResponse{},
	CallRequest{},
	CallResponse{},
))
