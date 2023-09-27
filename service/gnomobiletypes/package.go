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
	SetRemoteRequest{}, "SetRemote_Request",
	SetRemoteReply{}, "SetRemote_Reply",
	SetChainIDRequest{}, "SetChainID_Request",
	SetChainIDReply{}, "SetChainID_Reply",
	SetNameOrBech32Request{}, "SetNameOrBech32_Request",
	SetNameOrBech32Reply{}, "SetNameOrBech32_Reply",
	SetPasswordRequest{}, "SetPassword_Request",
	SetPasswordReply{}, "SetPassword_Reply",
	GenerateRecoveryPhraseRequest{}, "GenerateRecoveryPhrase_Request",
	GenerateRecoveryPhraseReply{}, "GenerateRecoveryPhrase_Reply",
	QueryRequest{}, "Query_Request",
	QueryReply{}, "Query_Reply",
	CallRequest{}, "Call_Request",
	CallReply{}, "Call_Reply",
))
