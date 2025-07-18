// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: rpc.proto

package _go

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The ErrCode enum defines errors for gRPC API functions. These are converted
// from the Go error types returned by gnoclient.
type ErrCode int32

const (
	// Undefined is the default value. It should never be set manually
	ErrCode_Undefined ErrCode = 0
	// TODO indicates that you plan to create an error later
	ErrCode_TODO ErrCode = 1
	// ErrNotImplemented indicates that a method is not implemented yet
	ErrCode_ErrNotImplemented ErrCode = 2
	// ErrInternal indicates an unknown error (without Code), i.e. in gRPC
	ErrCode_ErrInternal             ErrCode = 3
	ErrCode_ErrInvalidInput         ErrCode = 100
	ErrCode_ErrBridgeInterrupted    ErrCode = 101
	ErrCode_ErrMissingInput         ErrCode = 102
	ErrCode_ErrSerialization        ErrCode = 103
	ErrCode_ErrDeserialization      ErrCode = 104
	ErrCode_ErrInitService          ErrCode = 105
	ErrCode_ErrSetRemote            ErrCode = 106
	ErrCode_ErrCryptoKeyTypeUnknown ErrCode = 150
	// ErrCryptoKeyNotFound indicates that the doesn't exist in the keybase
	ErrCode_ErrCryptoKeyNotFound ErrCode = 151
	// ErrNoActiveAccount indicates that no account with the given address has been activated with ActivateAccount
	ErrCode_ErrNoActiveAccount ErrCode = 152
	ErrCode_ErrRunGRPCServer   ErrCode = 153
	// ErrDecryptionFailed indicates a decryption failure including a wrong password
	ErrCode_ErrDecryptionFailed ErrCode = 154
	ErrCode_ErrTxDecode         ErrCode = 200
	ErrCode_ErrInvalidSequence  ErrCode = 201
	ErrCode_ErrUnauthorized     ErrCode = 202
	// ErrInsufficientFunds indicates that there are insufficient funds to pay for fees
	ErrCode_ErrInsufficientFunds ErrCode = 203
	// ErrUnknownRequest indicates that the path of a realm function call is unrecognized
	ErrCode_ErrUnknownRequest ErrCode = 204
	// ErrInvalidAddress indicates that an account address is blank or the bech32 can't be decoded
	ErrCode_ErrInvalidAddress ErrCode = 205
	// ErrUnknownAddress indicates that the address is unknown on the blockchain
	ErrCode_ErrUnknownAddress ErrCode = 206
	// ErrInvalidPubKey indicates that the public key was not found or has an invalid algorithm or format
	ErrCode_ErrInvalidPubKey ErrCode = 207
	// ErrInsufficientCoins indicates that the transaction has insufficient account funds to send
	ErrCode_ErrInsufficientCoins ErrCode = 208
	// ErrInvalidCoins indicates that the transaction Coins are not sorted, or don't have a
	// positive amount, or the coin Denom contains upper case characters
	ErrCode_ErrInvalidCoins ErrCode = 209
	// ErrInvalidGasWanted indicates that the transaction gas wanted is too large or otherwise invalid
	ErrCode_ErrInvalidGasWanted ErrCode = 210
	// ErrOutOfGas indicates that the transaction doesn't have enough gas
	ErrCode_ErrOutOfGas ErrCode = 211
	// ErrMemoTooLarge indicates that the transaction memo is too large
	ErrCode_ErrMemoTooLarge ErrCode = 212
	// ErrInsufficientFee indicates that the gas fee is insufficient
	ErrCode_ErrInsufficientFee ErrCode = 213
	// ErrTooManySignatures indicates that the transaction has too many signatures
	ErrCode_ErrTooManySignatures ErrCode = 214
	// ErrNoSignatures indicates that the transaction has no signatures
	ErrCode_ErrNoSignatures ErrCode = 215
	// ErrGasOverflow indicates that an action results in a gas consumption unsigned integer overflow
	ErrCode_ErrGasOverflow ErrCode = 216
	// ErrInvalidPkgPath indicates that the package path is not recognized.
	ErrCode_ErrInvalidPkgPath ErrCode = 217
	ErrCode_ErrInvalidStmt    ErrCode = 218
	ErrCode_ErrInvalidExpr    ErrCode = 219
)

// Enum value maps for ErrCode.
var (
	ErrCode_name = map[int32]string{
		0:   "Undefined",
		1:   "TODO",
		2:   "ErrNotImplemented",
		3:   "ErrInternal",
		100: "ErrInvalidInput",
		101: "ErrBridgeInterrupted",
		102: "ErrMissingInput",
		103: "ErrSerialization",
		104: "ErrDeserialization",
		105: "ErrInitService",
		106: "ErrSetRemote",
		150: "ErrCryptoKeyTypeUnknown",
		151: "ErrCryptoKeyNotFound",
		152: "ErrNoActiveAccount",
		153: "ErrRunGRPCServer",
		154: "ErrDecryptionFailed",
		200: "ErrTxDecode",
		201: "ErrInvalidSequence",
		202: "ErrUnauthorized",
		203: "ErrInsufficientFunds",
		204: "ErrUnknownRequest",
		205: "ErrInvalidAddress",
		206: "ErrUnknownAddress",
		207: "ErrInvalidPubKey",
		208: "ErrInsufficientCoins",
		209: "ErrInvalidCoins",
		210: "ErrInvalidGasWanted",
		211: "ErrOutOfGas",
		212: "ErrMemoTooLarge",
		213: "ErrInsufficientFee",
		214: "ErrTooManySignatures",
		215: "ErrNoSignatures",
		216: "ErrGasOverflow",
		217: "ErrInvalidPkgPath",
		218: "ErrInvalidStmt",
		219: "ErrInvalidExpr",
	}
	ErrCode_value = map[string]int32{
		"Undefined":               0,
		"TODO":                    1,
		"ErrNotImplemented":       2,
		"ErrInternal":             3,
		"ErrInvalidInput":         100,
		"ErrBridgeInterrupted":    101,
		"ErrMissingInput":         102,
		"ErrSerialization":        103,
		"ErrDeserialization":      104,
		"ErrInitService":          105,
		"ErrSetRemote":            106,
		"ErrCryptoKeyTypeUnknown": 150,
		"ErrCryptoKeyNotFound":    151,
		"ErrNoActiveAccount":      152,
		"ErrRunGRPCServer":        153,
		"ErrDecryptionFailed":     154,
		"ErrTxDecode":             200,
		"ErrInvalidSequence":      201,
		"ErrUnauthorized":         202,
		"ErrInsufficientFunds":    203,
		"ErrUnknownRequest":       204,
		"ErrInvalidAddress":       205,
		"ErrUnknownAddress":       206,
		"ErrInvalidPubKey":        207,
		"ErrInsufficientCoins":    208,
		"ErrInvalidCoins":         209,
		"ErrInvalidGasWanted":     210,
		"ErrOutOfGas":             211,
		"ErrMemoTooLarge":         212,
		"ErrInsufficientFee":      213,
		"ErrTooManySignatures":    214,
		"ErrNoSignatures":         215,
		"ErrGasOverflow":          216,
		"ErrInvalidPkgPath":       217,
		"ErrInvalidStmt":          218,
		"ErrInvalidExpr":          219,
	}
)

func (x ErrCode) Enum() *ErrCode {
	p := new(ErrCode)
	*p = x
	return p
}

func (x ErrCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_proto_enumTypes[0].Descriptor()
}

func (ErrCode) Type() protoreflect.EnumType {
	return &file_rpc_proto_enumTypes[0]
}

func (x ErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrCode.Descriptor instead.
func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

type ErrDetails struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Codes         []ErrCode              `protobuf:"varint,1,rep,packed,name=codes,proto3,enum=land.gno.gnonative.v1.ErrCode" json:"codes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ErrDetails) Reset() {
	*x = ErrDetails{}
	mi := &file_rpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ErrDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrDetails) ProtoMessage() {}

func (x *ErrDetails) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrDetails.ProtoReflect.Descriptor instead.
func (*ErrDetails) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *ErrDetails) GetCodes() []ErrCode {
	if x != nil {
		return x.Codes
	}
	return nil
}

var File_rpc_proto protoreflect.FileDescriptor

const file_rpc_proto_rawDesc = "" +
	"\n" +
	"\trpc.proto\x12\x15land.gno.gnonative.v1\x1a\x14gnonativetypes.proto\"B\n" +
	"\n" +
	"ErrDetails\x124\n" +
	"\x05codes\x18\x01 \x03(\x0e2\x1e.land.gno.gnonative.v1.ErrCodeR\x05codes*\xb4\x06\n" +
	"\aErrCode\x12\r\n" +
	"\tUndefined\x10\x00\x12\b\n" +
	"\x04TODO\x10\x01\x12\x15\n" +
	"\x11ErrNotImplemented\x10\x02\x12\x0f\n" +
	"\vErrInternal\x10\x03\x12\x13\n" +
	"\x0fErrInvalidInput\x10d\x12\x18\n" +
	"\x14ErrBridgeInterrupted\x10e\x12\x13\n" +
	"\x0fErrMissingInput\x10f\x12\x14\n" +
	"\x10ErrSerialization\x10g\x12\x16\n" +
	"\x12ErrDeserialization\x10h\x12\x12\n" +
	"\x0eErrInitService\x10i\x12\x10\n" +
	"\fErrSetRemote\x10j\x12\x1c\n" +
	"\x17ErrCryptoKeyTypeUnknown\x10\x96\x01\x12\x19\n" +
	"\x14ErrCryptoKeyNotFound\x10\x97\x01\x12\x17\n" +
	"\x12ErrNoActiveAccount\x10\x98\x01\x12\x15\n" +
	"\x10ErrRunGRPCServer\x10\x99\x01\x12\x18\n" +
	"\x13ErrDecryptionFailed\x10\x9a\x01\x12\x10\n" +
	"\vErrTxDecode\x10\xc8\x01\x12\x17\n" +
	"\x12ErrInvalidSequence\x10\xc9\x01\x12\x14\n" +
	"\x0fErrUnauthorized\x10\xca\x01\x12\x19\n" +
	"\x14ErrInsufficientFunds\x10\xcb\x01\x12\x16\n" +
	"\x11ErrUnknownRequest\x10\xcc\x01\x12\x16\n" +
	"\x11ErrInvalidAddress\x10\xcd\x01\x12\x16\n" +
	"\x11ErrUnknownAddress\x10\xce\x01\x12\x15\n" +
	"\x10ErrInvalidPubKey\x10\xcf\x01\x12\x19\n" +
	"\x14ErrInsufficientCoins\x10\xd0\x01\x12\x14\n" +
	"\x0fErrInvalidCoins\x10\xd1\x01\x12\x18\n" +
	"\x13ErrInvalidGasWanted\x10\xd2\x01\x12\x10\n" +
	"\vErrOutOfGas\x10\xd3\x01\x12\x14\n" +
	"\x0fErrMemoTooLarge\x10\xd4\x01\x12\x17\n" +
	"\x12ErrInsufficientFee\x10\xd5\x01\x12\x19\n" +
	"\x14ErrTooManySignatures\x10\xd6\x01\x12\x14\n" +
	"\x0fErrNoSignatures\x10\xd7\x01\x12\x13\n" +
	"\x0eErrGasOverflow\x10\xd8\x01\x12\x16\n" +
	"\x11ErrInvalidPkgPath\x10\xd9\x01\x12\x13\n" +
	"\x0eErrInvalidStmt\x10\xda\x01\x12\x13\n" +
	"\x0eErrInvalidExpr\x10\xdb\x012\xbb \n" +
	"\x10GnoNativeService\x12^\n" +
	"\tSetRemote\x12'.land.gno.gnonative.v1.SetRemoteRequest\x1a(.land.gno.gnonative.v1.SetRemoteResponse\x12^\n" +
	"\tGetRemote\x12'.land.gno.gnonative.v1.GetRemoteRequest\x1a(.land.gno.gnonative.v1.GetRemoteResponse\x12a\n" +
	"\n" +
	"SetChainID\x12(.land.gno.gnonative.v1.SetChainIDRequest\x1a).land.gno.gnonative.v1.SetChainIDResponse\x12a\n" +
	"\n" +
	"GetChainID\x12(.land.gno.gnonative.v1.GetChainIDRequest\x1a).land.gno.gnonative.v1.GetChainIDResponse\x12\x85\x01\n" +
	"\x16GenerateRecoveryPhrase\x124.land.gno.gnonative.v1.GenerateRecoveryPhraseRequest\x1a5.land.gno.gnonative.v1.GenerateRecoveryPhraseResponse\x12d\n" +
	"\vListKeyInfo\x12).land.gno.gnonative.v1.ListKeyInfoRequest\x1a*.land.gno.gnonative.v1.ListKeyInfoResponse\x12g\n" +
	"\fHasKeyByName\x12*.land.gno.gnonative.v1.HasKeyByNameRequest\x1a+.land.gno.gnonative.v1.HasKeyByNameResponse\x12p\n" +
	"\x0fHasKeyByAddress\x12-.land.gno.gnonative.v1.HasKeyByAddressRequest\x1a..land.gno.gnonative.v1.HasKeyByAddressResponse\x12\x82\x01\n" +
	"\x15HasKeyByNameOrAddress\x123.land.gno.gnonative.v1.HasKeyByNameOrAddressRequest\x1a4.land.gno.gnonative.v1.HasKeyByNameOrAddressResponse\x12s\n" +
	"\x10GetKeyInfoByName\x12..land.gno.gnonative.v1.GetKeyInfoByNameRequest\x1a/.land.gno.gnonative.v1.GetKeyInfoByNameResponse\x12|\n" +
	"\x13GetKeyInfoByAddress\x121.land.gno.gnonative.v1.GetKeyInfoByAddressRequest\x1a2.land.gno.gnonative.v1.GetKeyInfoByAddressResponse\x12\x8e\x01\n" +
	"\x19GetKeyInfoByNameOrAddress\x127.land.gno.gnonative.v1.GetKeyInfoByNameOrAddressRequest\x1a8.land.gno.gnonative.v1.GetKeyInfoByNameOrAddressResponse\x12j\n" +
	"\rCreateAccount\x12+.land.gno.gnonative.v1.CreateAccountRequest\x1a,.land.gno.gnonative.v1.CreateAccountResponse\x12g\n" +
	"\fCreateLedger\x12*.land.gno.gnonative.v1.CreateLedgerRequest\x1a+.land.gno.gnonative.v1.CreateLedgerResponse\x12p\n" +
	"\x0fActivateAccount\x12-.land.gno.gnonative.v1.ActivateAccountRequest\x1a..land.gno.gnonative.v1.ActivateAccountResponse\x12d\n" +
	"\vSetPassword\x12).land.gno.gnonative.v1.SetPasswordRequest\x1a*.land.gno.gnonative.v1.SetPasswordResponse\x12m\n" +
	"\x0eRotatePassword\x12,.land.gno.gnonative.v1.RotatePasswordRequest\x1a-.land.gno.gnonative.v1.RotatePasswordResponse\x12|\n" +
	"\x13GetActivatedAccount\x121.land.gno.gnonative.v1.GetActivatedAccountRequest\x1a2.land.gno.gnonative.v1.GetActivatedAccountResponse\x12g\n" +
	"\fQueryAccount\x12*.land.gno.gnonative.v1.QueryAccountRequest\x1a+.land.gno.gnonative.v1.QueryAccountResponse\x12j\n" +
	"\rDeleteAccount\x12+.land.gno.gnonative.v1.DeleteAccountRequest\x1a,.land.gno.gnonative.v1.DeleteAccountResponse\x12R\n" +
	"\x05Query\x12#.land.gno.gnonative.v1.QueryRequest\x1a$.land.gno.gnonative.v1.QueryResponse\x12U\n" +
	"\x06Render\x12$.land.gno.gnonative.v1.RenderRequest\x1a%.land.gno.gnonative.v1.RenderResponse\x12R\n" +
	"\x05QEval\x12#.land.gno.gnonative.v1.QEvalRequest\x1a$.land.gno.gnonative.v1.QEvalResponse\x12Q\n" +
	"\x04Call\x12\".land.gno.gnonative.v1.CallRequest\x1a#.land.gno.gnonative.v1.CallResponse0\x01\x12Q\n" +
	"\x04Send\x12\".land.gno.gnonative.v1.SendRequest\x1a#.land.gno.gnonative.v1.SendResponse0\x01\x12N\n" +
	"\x03Run\x12!.land.gno.gnonative.v1.RunRequest\x1a\".land.gno.gnonative.v1.RunResponse0\x01\x12W\n" +
	"\n" +
	"MakeCallTx\x12\".land.gno.gnonative.v1.CallRequest\x1a%.land.gno.gnonative.v1.MakeTxResponse\x12W\n" +
	"\n" +
	"MakeSendTx\x12\".land.gno.gnonative.v1.SendRequest\x1a%.land.gno.gnonative.v1.MakeTxResponse\x12U\n" +
	"\tMakeRunTx\x12!.land.gno.gnonative.v1.RunRequest\x1a%.land.gno.gnonative.v1.MakeTxResponse\x12d\n" +
	"\vEstimateGas\x12).land.gno.gnonative.v1.EstimateGasRequest\x1a*.land.gno.gnonative.v1.EstimateGasResponse\x12U\n" +
	"\x06SignTx\x12$.land.gno.gnonative.v1.SignTxRequest\x1a%.land.gno.gnonative.v1.SignTxResponse\x12x\n" +
	"\x11BroadcastTxCommit\x12/.land.gno.gnonative.v1.BroadcastTxCommitRequest\x1a0.land.gno.gnonative.v1.BroadcastTxCommitResponse0\x01\x12p\n" +
	"\x0fAddressToBech32\x12-.land.gno.gnonative.v1.AddressToBech32Request\x1a..land.gno.gnonative.v1.AddressToBech32Response\x12v\n" +
	"\x11AddressFromBech32\x12/.land.gno.gnonative.v1.AddressFromBech32Request\x1a0.land.gno.gnonative.v1.AddressFromBech32Response\x12|\n" +
	"\x13AddressFromMnemonic\x121.land.gno.gnonative.v1.AddressFromMnemonicRequest\x1a2.land.gno.gnonative.v1.AddressFromMnemonicResponse\x12\x7f\n" +
	"\x14ValidateMnemonicWord\x122.land.gno.gnonative.v1.ValidateMnemonicWordRequest\x1a3.land.gno.gnonative.v1.ValidateMnemonicWordResponse\x12\x85\x01\n" +
	"\x16ValidateMnemonicPhrase\x124.land.gno.gnonative.v1.ValidateMnemonicPhraseRequest\x1a5.land.gno.gnonative.v1.ValidateMnemonicPhraseResponse\x12R\n" +
	"\x05Hello\x12#.land.gno.gnonative.v1.HelloRequest\x1a$.land.gno.gnonative.v1.HelloResponse\x12f\n" +
	"\vHelloStream\x12).land.gno.gnonative.v1.HelloStreamRequest\x1a*.land.gno.gnonative.v1.HelloStreamResponse0\x01B2Z*github.com/gnolang/gnonative/v4/api/gen/go\xa2\x02\x03RTGb\x06proto3"

var (
	file_rpc_proto_rawDescOnce sync.Once
	file_rpc_proto_rawDescData []byte
)

func file_rpc_proto_rawDescGZIP() []byte {
	file_rpc_proto_rawDescOnce.Do(func() {
		file_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_rpc_proto_rawDesc), len(file_rpc_proto_rawDesc)))
	})
	return file_rpc_proto_rawDescData
}

var file_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rpc_proto_goTypes = []any{
	(ErrCode)(0),                              // 0: land.gno.gnonative.v1.ErrCode
	(*ErrDetails)(nil),                        // 1: land.gno.gnonative.v1.ErrDetails
	(*SetRemoteRequest)(nil),                  // 2: land.gno.gnonative.v1.SetRemoteRequest
	(*GetRemoteRequest)(nil),                  // 3: land.gno.gnonative.v1.GetRemoteRequest
	(*SetChainIDRequest)(nil),                 // 4: land.gno.gnonative.v1.SetChainIDRequest
	(*GetChainIDRequest)(nil),                 // 5: land.gno.gnonative.v1.GetChainIDRequest
	(*GenerateRecoveryPhraseRequest)(nil),     // 6: land.gno.gnonative.v1.GenerateRecoveryPhraseRequest
	(*ListKeyInfoRequest)(nil),                // 7: land.gno.gnonative.v1.ListKeyInfoRequest
	(*HasKeyByNameRequest)(nil),               // 8: land.gno.gnonative.v1.HasKeyByNameRequest
	(*HasKeyByAddressRequest)(nil),            // 9: land.gno.gnonative.v1.HasKeyByAddressRequest
	(*HasKeyByNameOrAddressRequest)(nil),      // 10: land.gno.gnonative.v1.HasKeyByNameOrAddressRequest
	(*GetKeyInfoByNameRequest)(nil),           // 11: land.gno.gnonative.v1.GetKeyInfoByNameRequest
	(*GetKeyInfoByAddressRequest)(nil),        // 12: land.gno.gnonative.v1.GetKeyInfoByAddressRequest
	(*GetKeyInfoByNameOrAddressRequest)(nil),  // 13: land.gno.gnonative.v1.GetKeyInfoByNameOrAddressRequest
	(*CreateAccountRequest)(nil),              // 14: land.gno.gnonative.v1.CreateAccountRequest
	(*CreateLedgerRequest)(nil),               // 15: land.gno.gnonative.v1.CreateLedgerRequest
	(*ActivateAccountRequest)(nil),            // 16: land.gno.gnonative.v1.ActivateAccountRequest
	(*SetPasswordRequest)(nil),                // 17: land.gno.gnonative.v1.SetPasswordRequest
	(*RotatePasswordRequest)(nil),             // 18: land.gno.gnonative.v1.RotatePasswordRequest
	(*GetActivatedAccountRequest)(nil),        // 19: land.gno.gnonative.v1.GetActivatedAccountRequest
	(*QueryAccountRequest)(nil),               // 20: land.gno.gnonative.v1.QueryAccountRequest
	(*DeleteAccountRequest)(nil),              // 21: land.gno.gnonative.v1.DeleteAccountRequest
	(*QueryRequest)(nil),                      // 22: land.gno.gnonative.v1.QueryRequest
	(*RenderRequest)(nil),                     // 23: land.gno.gnonative.v1.RenderRequest
	(*QEvalRequest)(nil),                      // 24: land.gno.gnonative.v1.QEvalRequest
	(*CallRequest)(nil),                       // 25: land.gno.gnonative.v1.CallRequest
	(*SendRequest)(nil),                       // 26: land.gno.gnonative.v1.SendRequest
	(*RunRequest)(nil),                        // 27: land.gno.gnonative.v1.RunRequest
	(*EstimateGasRequest)(nil),                // 28: land.gno.gnonative.v1.EstimateGasRequest
	(*SignTxRequest)(nil),                     // 29: land.gno.gnonative.v1.SignTxRequest
	(*BroadcastTxCommitRequest)(nil),          // 30: land.gno.gnonative.v1.BroadcastTxCommitRequest
	(*AddressToBech32Request)(nil),            // 31: land.gno.gnonative.v1.AddressToBech32Request
	(*AddressFromBech32Request)(nil),          // 32: land.gno.gnonative.v1.AddressFromBech32Request
	(*AddressFromMnemonicRequest)(nil),        // 33: land.gno.gnonative.v1.AddressFromMnemonicRequest
	(*ValidateMnemonicWordRequest)(nil),       // 34: land.gno.gnonative.v1.ValidateMnemonicWordRequest
	(*ValidateMnemonicPhraseRequest)(nil),     // 35: land.gno.gnonative.v1.ValidateMnemonicPhraseRequest
	(*HelloRequest)(nil),                      // 36: land.gno.gnonative.v1.HelloRequest
	(*HelloStreamRequest)(nil),                // 37: land.gno.gnonative.v1.HelloStreamRequest
	(*SetRemoteResponse)(nil),                 // 38: land.gno.gnonative.v1.SetRemoteResponse
	(*GetRemoteResponse)(nil),                 // 39: land.gno.gnonative.v1.GetRemoteResponse
	(*SetChainIDResponse)(nil),                // 40: land.gno.gnonative.v1.SetChainIDResponse
	(*GetChainIDResponse)(nil),                // 41: land.gno.gnonative.v1.GetChainIDResponse
	(*GenerateRecoveryPhraseResponse)(nil),    // 42: land.gno.gnonative.v1.GenerateRecoveryPhraseResponse
	(*ListKeyInfoResponse)(nil),               // 43: land.gno.gnonative.v1.ListKeyInfoResponse
	(*HasKeyByNameResponse)(nil),              // 44: land.gno.gnonative.v1.HasKeyByNameResponse
	(*HasKeyByAddressResponse)(nil),           // 45: land.gno.gnonative.v1.HasKeyByAddressResponse
	(*HasKeyByNameOrAddressResponse)(nil),     // 46: land.gno.gnonative.v1.HasKeyByNameOrAddressResponse
	(*GetKeyInfoByNameResponse)(nil),          // 47: land.gno.gnonative.v1.GetKeyInfoByNameResponse
	(*GetKeyInfoByAddressResponse)(nil),       // 48: land.gno.gnonative.v1.GetKeyInfoByAddressResponse
	(*GetKeyInfoByNameOrAddressResponse)(nil), // 49: land.gno.gnonative.v1.GetKeyInfoByNameOrAddressResponse
	(*CreateAccountResponse)(nil),             // 50: land.gno.gnonative.v1.CreateAccountResponse
	(*CreateLedgerResponse)(nil),              // 51: land.gno.gnonative.v1.CreateLedgerResponse
	(*ActivateAccountResponse)(nil),           // 52: land.gno.gnonative.v1.ActivateAccountResponse
	(*SetPasswordResponse)(nil),               // 53: land.gno.gnonative.v1.SetPasswordResponse
	(*RotatePasswordResponse)(nil),            // 54: land.gno.gnonative.v1.RotatePasswordResponse
	(*GetActivatedAccountResponse)(nil),       // 55: land.gno.gnonative.v1.GetActivatedAccountResponse
	(*QueryAccountResponse)(nil),              // 56: land.gno.gnonative.v1.QueryAccountResponse
	(*DeleteAccountResponse)(nil),             // 57: land.gno.gnonative.v1.DeleteAccountResponse
	(*QueryResponse)(nil),                     // 58: land.gno.gnonative.v1.QueryResponse
	(*RenderResponse)(nil),                    // 59: land.gno.gnonative.v1.RenderResponse
	(*QEvalResponse)(nil),                     // 60: land.gno.gnonative.v1.QEvalResponse
	(*CallResponse)(nil),                      // 61: land.gno.gnonative.v1.CallResponse
	(*SendResponse)(nil),                      // 62: land.gno.gnonative.v1.SendResponse
	(*RunResponse)(nil),                       // 63: land.gno.gnonative.v1.RunResponse
	(*MakeTxResponse)(nil),                    // 64: land.gno.gnonative.v1.MakeTxResponse
	(*EstimateGasResponse)(nil),               // 65: land.gno.gnonative.v1.EstimateGasResponse
	(*SignTxResponse)(nil),                    // 66: land.gno.gnonative.v1.SignTxResponse
	(*BroadcastTxCommitResponse)(nil),         // 67: land.gno.gnonative.v1.BroadcastTxCommitResponse
	(*AddressToBech32Response)(nil),           // 68: land.gno.gnonative.v1.AddressToBech32Response
	(*AddressFromBech32Response)(nil),         // 69: land.gno.gnonative.v1.AddressFromBech32Response
	(*AddressFromMnemonicResponse)(nil),       // 70: land.gno.gnonative.v1.AddressFromMnemonicResponse
	(*ValidateMnemonicWordResponse)(nil),      // 71: land.gno.gnonative.v1.ValidateMnemonicWordResponse
	(*ValidateMnemonicPhraseResponse)(nil),    // 72: land.gno.gnonative.v1.ValidateMnemonicPhraseResponse
	(*HelloResponse)(nil),                     // 73: land.gno.gnonative.v1.HelloResponse
	(*HelloStreamResponse)(nil),               // 74: land.gno.gnonative.v1.HelloStreamResponse
}
var file_rpc_proto_depIdxs = []int32{
	0,  // 0: land.gno.gnonative.v1.ErrDetails.codes:type_name -> land.gno.gnonative.v1.ErrCode
	2,  // 1: land.gno.gnonative.v1.GnoNativeService.SetRemote:input_type -> land.gno.gnonative.v1.SetRemoteRequest
	3,  // 2: land.gno.gnonative.v1.GnoNativeService.GetRemote:input_type -> land.gno.gnonative.v1.GetRemoteRequest
	4,  // 3: land.gno.gnonative.v1.GnoNativeService.SetChainID:input_type -> land.gno.gnonative.v1.SetChainIDRequest
	5,  // 4: land.gno.gnonative.v1.GnoNativeService.GetChainID:input_type -> land.gno.gnonative.v1.GetChainIDRequest
	6,  // 5: land.gno.gnonative.v1.GnoNativeService.GenerateRecoveryPhrase:input_type -> land.gno.gnonative.v1.GenerateRecoveryPhraseRequest
	7,  // 6: land.gno.gnonative.v1.GnoNativeService.ListKeyInfo:input_type -> land.gno.gnonative.v1.ListKeyInfoRequest
	8,  // 7: land.gno.gnonative.v1.GnoNativeService.HasKeyByName:input_type -> land.gno.gnonative.v1.HasKeyByNameRequest
	9,  // 8: land.gno.gnonative.v1.GnoNativeService.HasKeyByAddress:input_type -> land.gno.gnonative.v1.HasKeyByAddressRequest
	10, // 9: land.gno.gnonative.v1.GnoNativeService.HasKeyByNameOrAddress:input_type -> land.gno.gnonative.v1.HasKeyByNameOrAddressRequest
	11, // 10: land.gno.gnonative.v1.GnoNativeService.GetKeyInfoByName:input_type -> land.gno.gnonative.v1.GetKeyInfoByNameRequest
	12, // 11: land.gno.gnonative.v1.GnoNativeService.GetKeyInfoByAddress:input_type -> land.gno.gnonative.v1.GetKeyInfoByAddressRequest
	13, // 12: land.gno.gnonative.v1.GnoNativeService.GetKeyInfoByNameOrAddress:input_type -> land.gno.gnonative.v1.GetKeyInfoByNameOrAddressRequest
	14, // 13: land.gno.gnonative.v1.GnoNativeService.CreateAccount:input_type -> land.gno.gnonative.v1.CreateAccountRequest
	15, // 14: land.gno.gnonative.v1.GnoNativeService.CreateLedger:input_type -> land.gno.gnonative.v1.CreateLedgerRequest
	16, // 15: land.gno.gnonative.v1.GnoNativeService.ActivateAccount:input_type -> land.gno.gnonative.v1.ActivateAccountRequest
	17, // 16: land.gno.gnonative.v1.GnoNativeService.SetPassword:input_type -> land.gno.gnonative.v1.SetPasswordRequest
	18, // 17: land.gno.gnonative.v1.GnoNativeService.RotatePassword:input_type -> land.gno.gnonative.v1.RotatePasswordRequest
	19, // 18: land.gno.gnonative.v1.GnoNativeService.GetActivatedAccount:input_type -> land.gno.gnonative.v1.GetActivatedAccountRequest
	20, // 19: land.gno.gnonative.v1.GnoNativeService.QueryAccount:input_type -> land.gno.gnonative.v1.QueryAccountRequest
	21, // 20: land.gno.gnonative.v1.GnoNativeService.DeleteAccount:input_type -> land.gno.gnonative.v1.DeleteAccountRequest
	22, // 21: land.gno.gnonative.v1.GnoNativeService.Query:input_type -> land.gno.gnonative.v1.QueryRequest
	23, // 22: land.gno.gnonative.v1.GnoNativeService.Render:input_type -> land.gno.gnonative.v1.RenderRequest
	24, // 23: land.gno.gnonative.v1.GnoNativeService.QEval:input_type -> land.gno.gnonative.v1.QEvalRequest
	25, // 24: land.gno.gnonative.v1.GnoNativeService.Call:input_type -> land.gno.gnonative.v1.CallRequest
	26, // 25: land.gno.gnonative.v1.GnoNativeService.Send:input_type -> land.gno.gnonative.v1.SendRequest
	27, // 26: land.gno.gnonative.v1.GnoNativeService.Run:input_type -> land.gno.gnonative.v1.RunRequest
	25, // 27: land.gno.gnonative.v1.GnoNativeService.MakeCallTx:input_type -> land.gno.gnonative.v1.CallRequest
	26, // 28: land.gno.gnonative.v1.GnoNativeService.MakeSendTx:input_type -> land.gno.gnonative.v1.SendRequest
	27, // 29: land.gno.gnonative.v1.GnoNativeService.MakeRunTx:input_type -> land.gno.gnonative.v1.RunRequest
	28, // 30: land.gno.gnonative.v1.GnoNativeService.EstimateGas:input_type -> land.gno.gnonative.v1.EstimateGasRequest
	29, // 31: land.gno.gnonative.v1.GnoNativeService.SignTx:input_type -> land.gno.gnonative.v1.SignTxRequest
	30, // 32: land.gno.gnonative.v1.GnoNativeService.BroadcastTxCommit:input_type -> land.gno.gnonative.v1.BroadcastTxCommitRequest
	31, // 33: land.gno.gnonative.v1.GnoNativeService.AddressToBech32:input_type -> land.gno.gnonative.v1.AddressToBech32Request
	32, // 34: land.gno.gnonative.v1.GnoNativeService.AddressFromBech32:input_type -> land.gno.gnonative.v1.AddressFromBech32Request
	33, // 35: land.gno.gnonative.v1.GnoNativeService.AddressFromMnemonic:input_type -> land.gno.gnonative.v1.AddressFromMnemonicRequest
	34, // 36: land.gno.gnonative.v1.GnoNativeService.ValidateMnemonicWord:input_type -> land.gno.gnonative.v1.ValidateMnemonicWordRequest
	35, // 37: land.gno.gnonative.v1.GnoNativeService.ValidateMnemonicPhrase:input_type -> land.gno.gnonative.v1.ValidateMnemonicPhraseRequest
	36, // 38: land.gno.gnonative.v1.GnoNativeService.Hello:input_type -> land.gno.gnonative.v1.HelloRequest
	37, // 39: land.gno.gnonative.v1.GnoNativeService.HelloStream:input_type -> land.gno.gnonative.v1.HelloStreamRequest
	38, // 40: land.gno.gnonative.v1.GnoNativeService.SetRemote:output_type -> land.gno.gnonative.v1.SetRemoteResponse
	39, // 41: land.gno.gnonative.v1.GnoNativeService.GetRemote:output_type -> land.gno.gnonative.v1.GetRemoteResponse
	40, // 42: land.gno.gnonative.v1.GnoNativeService.SetChainID:output_type -> land.gno.gnonative.v1.SetChainIDResponse
	41, // 43: land.gno.gnonative.v1.GnoNativeService.GetChainID:output_type -> land.gno.gnonative.v1.GetChainIDResponse
	42, // 44: land.gno.gnonative.v1.GnoNativeService.GenerateRecoveryPhrase:output_type -> land.gno.gnonative.v1.GenerateRecoveryPhraseResponse
	43, // 45: land.gno.gnonative.v1.GnoNativeService.ListKeyInfo:output_type -> land.gno.gnonative.v1.ListKeyInfoResponse
	44, // 46: land.gno.gnonative.v1.GnoNativeService.HasKeyByName:output_type -> land.gno.gnonative.v1.HasKeyByNameResponse
	45, // 47: land.gno.gnonative.v1.GnoNativeService.HasKeyByAddress:output_type -> land.gno.gnonative.v1.HasKeyByAddressResponse
	46, // 48: land.gno.gnonative.v1.GnoNativeService.HasKeyByNameOrAddress:output_type -> land.gno.gnonative.v1.HasKeyByNameOrAddressResponse
	47, // 49: land.gno.gnonative.v1.GnoNativeService.GetKeyInfoByName:output_type -> land.gno.gnonative.v1.GetKeyInfoByNameResponse
	48, // 50: land.gno.gnonative.v1.GnoNativeService.GetKeyInfoByAddress:output_type -> land.gno.gnonative.v1.GetKeyInfoByAddressResponse
	49, // 51: land.gno.gnonative.v1.GnoNativeService.GetKeyInfoByNameOrAddress:output_type -> land.gno.gnonative.v1.GetKeyInfoByNameOrAddressResponse
	50, // 52: land.gno.gnonative.v1.GnoNativeService.CreateAccount:output_type -> land.gno.gnonative.v1.CreateAccountResponse
	51, // 53: land.gno.gnonative.v1.GnoNativeService.CreateLedger:output_type -> land.gno.gnonative.v1.CreateLedgerResponse
	52, // 54: land.gno.gnonative.v1.GnoNativeService.ActivateAccount:output_type -> land.gno.gnonative.v1.ActivateAccountResponse
	53, // 55: land.gno.gnonative.v1.GnoNativeService.SetPassword:output_type -> land.gno.gnonative.v1.SetPasswordResponse
	54, // 56: land.gno.gnonative.v1.GnoNativeService.RotatePassword:output_type -> land.gno.gnonative.v1.RotatePasswordResponse
	55, // 57: land.gno.gnonative.v1.GnoNativeService.GetActivatedAccount:output_type -> land.gno.gnonative.v1.GetActivatedAccountResponse
	56, // 58: land.gno.gnonative.v1.GnoNativeService.QueryAccount:output_type -> land.gno.gnonative.v1.QueryAccountResponse
	57, // 59: land.gno.gnonative.v1.GnoNativeService.DeleteAccount:output_type -> land.gno.gnonative.v1.DeleteAccountResponse
	58, // 60: land.gno.gnonative.v1.GnoNativeService.Query:output_type -> land.gno.gnonative.v1.QueryResponse
	59, // 61: land.gno.gnonative.v1.GnoNativeService.Render:output_type -> land.gno.gnonative.v1.RenderResponse
	60, // 62: land.gno.gnonative.v1.GnoNativeService.QEval:output_type -> land.gno.gnonative.v1.QEvalResponse
	61, // 63: land.gno.gnonative.v1.GnoNativeService.Call:output_type -> land.gno.gnonative.v1.CallResponse
	62, // 64: land.gno.gnonative.v1.GnoNativeService.Send:output_type -> land.gno.gnonative.v1.SendResponse
	63, // 65: land.gno.gnonative.v1.GnoNativeService.Run:output_type -> land.gno.gnonative.v1.RunResponse
	64, // 66: land.gno.gnonative.v1.GnoNativeService.MakeCallTx:output_type -> land.gno.gnonative.v1.MakeTxResponse
	64, // 67: land.gno.gnonative.v1.GnoNativeService.MakeSendTx:output_type -> land.gno.gnonative.v1.MakeTxResponse
	64, // 68: land.gno.gnonative.v1.GnoNativeService.MakeRunTx:output_type -> land.gno.gnonative.v1.MakeTxResponse
	65, // 69: land.gno.gnonative.v1.GnoNativeService.EstimateGas:output_type -> land.gno.gnonative.v1.EstimateGasResponse
	66, // 70: land.gno.gnonative.v1.GnoNativeService.SignTx:output_type -> land.gno.gnonative.v1.SignTxResponse
	67, // 71: land.gno.gnonative.v1.GnoNativeService.BroadcastTxCommit:output_type -> land.gno.gnonative.v1.BroadcastTxCommitResponse
	68, // 72: land.gno.gnonative.v1.GnoNativeService.AddressToBech32:output_type -> land.gno.gnonative.v1.AddressToBech32Response
	69, // 73: land.gno.gnonative.v1.GnoNativeService.AddressFromBech32:output_type -> land.gno.gnonative.v1.AddressFromBech32Response
	70, // 74: land.gno.gnonative.v1.GnoNativeService.AddressFromMnemonic:output_type -> land.gno.gnonative.v1.AddressFromMnemonicResponse
	71, // 75: land.gno.gnonative.v1.GnoNativeService.ValidateMnemonicWord:output_type -> land.gno.gnonative.v1.ValidateMnemonicWordResponse
	72, // 76: land.gno.gnonative.v1.GnoNativeService.ValidateMnemonicPhrase:output_type -> land.gno.gnonative.v1.ValidateMnemonicPhraseResponse
	73, // 77: land.gno.gnonative.v1.GnoNativeService.Hello:output_type -> land.gno.gnonative.v1.HelloResponse
	74, // 78: land.gno.gnonative.v1.GnoNativeService.HelloStream:output_type -> land.gno.gnonative.v1.HelloStreamResponse
	40, // [40:79] is the sub-list for method output_type
	1,  // [1:40] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_proto_init() }
func file_rpc_proto_init() {
	if File_rpc_proto != nil {
		return
	}
	file_gnonativetypes_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_rpc_proto_rawDesc), len(file_rpc_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_proto_depIdxs,
		EnumInfos:         file_rpc_proto_enumTypes,
		MessageInfos:      file_rpc_proto_msgTypes,
	}.Build()
	File_rpc_proto = out.File
	file_rpc_proto_goTypes = nil
	file_rpc_proto_depIdxs = nil
}
