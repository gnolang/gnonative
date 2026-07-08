// Package api holds the plain Go wire types and error codes for gnonative.
//
// The JSON encoding of these types reproduces the protobuf JSON (protojson)
// dialect that the mobile bridge and the @gnolang/gnonative TS layer expect:
//   - field names use the proto json_name (lowerCamelCase, or an explicit
//     override where noted);
//   - every field is omitempty (protojson omits zero values);
//   - int64/uint64 are encoded as strings (",string");
//   - []byte is standard base64 (encoding/json default, same as protojson).
//
// The authoritative field-by-field spec is expo/src/native/apitypes.ts, and the
// dialect is guarded by api/types_test.go. Keep the three in sync.
package api

type SetRemoteRequest struct {
	Remote string `json:"remote,omitempty"`
}

type SetRemoteResponse struct{}

type GetRemoteRequest struct{}

type GetRemoteResponse struct {
	Remote string `json:"remote,omitempty"`
}

type SetChainIDRequest struct {
	ChainID string `json:"chainId,omitempty"`
}

type SetChainIDResponse struct{}

type GetChainIDRequest struct{}

type GetChainIDResponse struct {
	ChainID string `json:"chainId,omitempty"`
}

type SetPasswordRequest struct {
	Password string `json:"password,omitempty"`
	// The address of the account to set the password
	Address []byte `json:"address,omitempty"`
}

type SetPasswordResponse struct{}

type RenameKeyRequest struct {
	OldName string `json:"oldName,omitempty"`
	NewName string `json:"newName,omitempty"`
}

type RenameKeyResponse struct{}

type RotatePasswordRequest struct {
	NewPassword string `json:"newPassword,omitempty"`
	// The addresses of the account to rotate the password
	Addresses [][]byte `json:"addresses,omitempty"`
}

type RotatePasswordResponse struct{}

type GenerateRecoveryPhraseRequest struct{}

type GenerateRecoveryPhraseResponse struct {
	Phrase string `json:"phrase,omitempty"`
}

type KeyInfo struct {
	// 0: local, 1: ledger, 2: offline, 3: multi
	Type    uint32 `json:"type,omitempty"`
	Name    string `json:"name,omitempty"`
	PubKey  []byte `json:"pubKey,omitempty"`
	Address []byte `json:"address,omitempty"`
}

// Coin holds some amount of one currency.
// A negative amount is invalid.
type Coin struct {
	// Example: "ugnot"
	Denom  string `json:"denom,omitempty"`
	Amount int64  `json:"amount,string,omitempty"`
}

type BaseAccount struct {
	Address       []byte  `json:"address,omitempty"`
	Coins         []*Coin `json:"coins,omitempty"`
	PubKey        []byte  `json:"pubKey,omitempty"`
	AccountNumber uint64  `json:"accountNumber,string,omitempty"`
	Sequence      uint64  `json:"sequence,string,omitempty"`
}

type SessionAccount struct {
	BaseAccount   *BaseAccount `json:"baseAccount,omitempty"`
	MasterAddress []byte       `json:"masterAddress,omitempty"`
	// Unix timestamp; 0 = no expiry
	ExpiresAt int64 `json:"expiresAt,string,omitempty"`
	// Nil/empty = no spending allowed (fail-closed, NOT unrestricted)
	SpendLimit []*Coin `json:"spendLimit,omitempty"`
	// Seconds; 0 = lifetime cap (no reset)
	SpendPeriod int64 `json:"spendPeriod,string,omitempty"`
	// Nil/empty = 0 spent
	SpendUsed []*Coin `json:"spendUsed,omitempty"`
	// Unix timestamp; start of current period
	SpendReset int64    `json:"spendReset,string,omitempty"`
	AllowPaths []string `json:"allowPaths,omitempty"`
}

type ListKeyInfoRequest struct{}

type ListKeyInfoResponse struct {
	Keys []*KeyInfo `json:"key_info,omitempty"` // proto json_name = "key_info"
}

type GetKeyInfoByNameRequest struct {
	Name string `json:"name,omitempty"`
}

type HasKeyByNameRequest struct {
	Name string `json:"name,omitempty"`
}

type HasKeyByNameResponse struct {
	Has bool `json:"has,omitempty"`
}

type HasKeyByAddressRequest struct {
	Address []byte `json:"address,omitempty"`
}

type HasKeyByAddressResponse struct {
	Has bool `json:"has,omitempty"`
}

type HasKeyByNameOrAddressRequest struct {
	NameOrBech32 string `json:"nameOrBech32,omitempty"`
}

type HasKeyByNameOrAddressResponse struct {
	Has bool `json:"has,omitempty"`
}

type GetKeyInfoByNameResponse struct {
	Key *KeyInfo `json:"key_info,omitempty"` // proto json_name = "key_info"
}

type GetKeyInfoByAddressRequest struct {
	Address []byte `json:"address,omitempty"`
}

type GetKeyInfoByAddressResponse struct {
	Key *KeyInfo `json:"key_info,omitempty"` // proto json_name = "key_info"
}

type GetKeyInfoByNameOrAddressRequest struct {
	NameOrBech32 string `json:"nameOrBech32,omitempty"`
}

type GetKeyInfoByNameOrAddressResponse struct {
	Key *KeyInfo `json:"key_info,omitempty"` // proto json_name = "key_info"
}

type CreateAccountRequest struct {
	NameOrBech32 string `json:"nameOrBech32,omitempty"`
	Mnemonic     string `json:"mnemonic,omitempty"`
	Bip39Passwd  string `json:"bip39Passwd,omitempty"`
	Password     string `json:"password,omitempty"`
	Account      uint32 `json:"account,omitempty"`
	Index        uint32 `json:"index,omitempty"`
}

type CreateAccountResponse struct {
	Key *KeyInfo `json:"key_info,omitempty"` // proto json_name = "key_info"
}

type CreateLedgerRequest struct {
	Name string `json:"name,omitempty"`
	// Supported algorithm is "secp256k1"
	Algorithm string `json:"algorithm,omitempty"`
	// The human readable part of the address. Example: "g"
	HRP     string `json:"hrp,omitempty"`
	Account uint32 `json:"account,omitempty"`
	Index   uint32 `json:"index,omitempty"`
}

type CreateLedgerResponse struct {
	Key *KeyInfo `json:"key_info,omitempty"` // proto json_name = "key_info"
}

type ActivateAccountRequest struct {
	NameOrBech32 string `json:"nameOrBech32,omitempty"`
	// (Optional) The address of the master account if this is a session account.
	Master []byte `json:"master,omitempty"`
}

type ActivateAccountResponse struct {
	Key *KeyInfo `json:"key_info,omitempty"` // proto json_name = "key_info"
	// True if the password has been set. If false, then call SetPassword.
	HasPassword bool `json:"hasPassword,omitempty"`
}

type GetActivatedAccountRequest struct {
	Address []byte `json:"address,omitempty"`
}

type GetActivatedAccountResponse struct {
	Key *KeyInfo `json:"key_info,omitempty"` // proto json_name = "key_info"
	// The Master which was given to ActivateAccount.
	Master []byte `json:"master,omitempty"`
	// True if the password has been set. If false, then call SetPassword.
	HasPassword bool `json:"hasPassword,omitempty"`
}

type QueryAccountRequest struct {
	Address []byte `json:"address,omitempty"`
}

type QueryAccountResponse struct {
	AccountInfo *BaseAccount `json:"accountInfo,omitempty"`
}

type QuerySessionAccountRequest struct {
	MasterAddress  []byte `json:"masterAddress,omitempty"`
	SessionAddress []byte `json:"sessionAddress,omitempty"`
}

type QuerySessionAccountResponse struct {
	AccountInfo *SessionAccount `json:"accountInfo,omitempty"`
}

type DeleteAccountRequest struct {
	NameOrBech32 string `json:"nameOrBech32,omitempty"`
	Password     string `json:"password,omitempty"`
	SkipPassword bool   `json:"skipPassword,omitempty"`
}

type DeleteAccountResponse struct{}

type QueryRequest struct {
	// Example: "vm/qrender"
	Path string `json:"path,omitempty"`
	// Example: "gno.land/r/demo/boards\ntestboard"
	Data []byte `json:"data,omitempty"`
}

type QueryResponse struct {
	Result []byte `json:"result,omitempty"`
}

type RenderRequest struct {
	// Example: "gno.land/r/demo/boards"
	PackagePath string `json:"packagePath,omitempty"`
	// Example: "testboard/1"
	Args string `json:"args,omitempty"`
}

type RenderResponse struct {
	// The Render function result (typically markdown)
	Result string `json:"result,omitempty"`
}

type QEvalRequest struct {
	// Example: "gno.land/r/demo/boards"
	PackagePath string `json:"packagePath,omitempty"`
	// Example: "GetBoardIDFromName(\"testboard\")"
	Expression string `json:"expression,omitempty"`
}

type QEvalResponse struct {
	// A typed expression like "(1 gno.land/r/demo/boards.BoardID)\n(true bool)"
	Result string `json:"result,omitempty"`
}

type MsgCall struct {
	// Example: "gno.land/r/demo/boards"
	PackagePath string `json:"packagePath,omitempty"`
	// Example: "CreateReply"
	Fnc string `json:"fnc,omitempty"`
	// list of arguments specific to the function
	// Example: ["1", "1", "2", "my reply"]
	Args []string `json:"args,omitempty"`
	// Optional. Example: [ {Denom: "ugnot", Amount: 1000} ]
	Send []*Coin `json:"send,omitempty"`
	// Optional max storage deposit. Example: [ {Denom: "ugnot", Amount: 500000} ]
	MaxDeposit []*Coin `json:"maxDeposit,omitempty"`
}

type CallRequest struct {
	GasFee    string `json:"gasFee,omitempty"`
	GasWanted int64  `json:"gasWanted,string,omitempty"`
	Memo      string `json:"memo,omitempty"`
	// The address of the account to sign the transaction
	SignerAddress []byte `json:"signerAddress,omitempty"`
	// list of calls to make in one transaction
	Msgs []*MsgCall `json:"Msgs,omitempty"` // proto json_name = "Msgs"
}

type CallResponse struct {
	Result []byte `json:"result,omitempty"`
	// The transaction hash
	Hash []byte `json:"hash,omitempty"`
	// The transaction height
	Height int64 `json:"height,string,omitempty"`
}

type MsgSend struct {
	// Example: The response of calling AddressFromBech32 with
	// "g1juz2yxmdsa6audkp6ep9vfv80c8p5u76e03vvh"
	ToAddress []byte `json:"toAddress,omitempty"`
	// Example: [ {Denom: "ugnot", Amount: 1000} ]
	Amount []*Coin `json:"amount,omitempty"`
}

type SendRequest struct {
	GasFee    string `json:"gasFee,omitempty"`
	GasWanted int64  `json:"gasWanted,string,omitempty"`
	// Memo is optional
	Memo string `json:"memo,omitempty"`
	// The address of the account to sign the transaction
	SignerAddress []byte `json:"signerAddress,omitempty"`
	// list of send operations to make in one transaction
	Msgs []*MsgSend `json:"Msgs,omitempty"` // proto json_name = "Msgs"
}

type SendResponse struct {
	// The transaction hash
	Hash []byte `json:"hash,omitempty"`
	// The transaction height
	Height int64 `json:"height,string,omitempty"`
}

type MsgRun struct {
	// The code for the script package. Must have main().
	// Example: "package main\nfunc main() {\n  println(\"Hello\")\n}"
	Package string `json:"package,omitempty"`
	// Optional. Example: [ {Denom: "ugnot", Amount: 1000} ]
	Send []*Coin `json:"send,omitempty"`
	// Optional max storage deposit. Example: [ {Denom: "ugnot", Amount: 500000} ]
	MaxDeposit []*Coin `json:"maxDeposit,omitempty"`
}

type RunRequest struct {
	GasFee    string `json:"gasFee,omitempty"`
	GasWanted int64  `json:"gasWanted,string,omitempty"`
	// Memo is optional
	Memo string `json:"memo,omitempty"`
	// The address of the account to sign the transaction
	SignerAddress []byte `json:"signerAddress,omitempty"`
	// list of run operations to make in one transaction
	Msgs []*MsgRun `json:"Msgs,omitempty"` // proto json_name = "Msgs"
}

type RunResponse struct {
	// The "console" output from the run
	Result string `json:"result,omitempty"`
	// The transaction hash
	Hash []byte `json:"hash,omitempty"`
	// The transaction height
	Height int64 `json:"height,string,omitempty"`
}

type MakeCallTxRequest struct {
	GasFee    string `json:"gasFee,omitempty"`
	GasWanted int64  `json:"gasWanted,string,omitempty"`
	Memo      string `json:"memo,omitempty"`
	// The address of the caller. If signing with a session account, this is the master account address
	CallerAddress []byte `json:"callerAddress,omitempty"`
	// list of calls to make in one transaction
	Msgs []*MsgCall `json:"Msgs,omitempty"` // proto json_name = "Msgs"
}

type MakeSendTxRequest struct {
	GasFee    string `json:"gasFee,omitempty"`
	GasWanted int64  `json:"gasWanted,string,omitempty"`
	// Memo is optional
	Memo string `json:"memo,omitempty"`
	// The address of the caller. If signing with a session account, this is the master account address
	CallerAddress []byte `json:"callerAddress,omitempty"`
	// list of send operations to make in one transaction
	Msgs []*MsgSend `json:"Msgs,omitempty"` // proto json_name = "Msgs"
}

type MakeRunTxRequest struct {
	GasFee    string `json:"gasFee,omitempty"`
	GasWanted int64  `json:"gasWanted,string,omitempty"`
	// Memo is optional
	Memo string `json:"memo,omitempty"`
	// The address of the caller. If signing with a session account, this is the master account address
	CallerAddress []byte `json:"callerAddress,omitempty"`
	// list of run operations to make in one transaction
	Msgs []*MsgRun `json:"Msgs,omitempty"` // proto json_name = "Msgs"
}

type MakeTxResponse struct {
	// The JSON encoding of the unsigned transaction
	TxJSON string `json:"txJson,omitempty"`
}

type SignTxRequest struct {
	// The JSON encoding of the unsigned transaction (from MakeCallTx, etc.)
	TxJSON string `json:"txJson,omitempty"`
	// The address of the account to sign the transaction
	Address []byte `json:"address,omitempty"`
	// The signer's account number on the blockchain. If 0 then query the blockchain for the value.
	AccountNumber uint64 `json:"accountNumber,string,omitempty"`
	// The sequence number of the signer's transactions on the blockchain. If 0 then query the blockchain for the value.
	SequenceNumber uint64 `json:"sequenceNumber,string,omitempty"`
}

type SignTxResponse struct {
	// The JSON encoding of the signed transaction (to use in BroadcastTx)
	SignedTxJSON string `json:"tx_json,omitempty"` // proto json_name = "tx_json"
}

type MsgCreateSession struct {
	// Full session public key
	SessionKey []byte `json:"sessionKey,omitempty"`
	// unix timestamp; 0 = no expiry
	ExpiresAt int64 `json:"expiresAt,string,omitempty"`
	// Typed entries: "*" or <route>/<type>[:<path>]; required (gno.land-specific grammar)
	AllowPaths []string `json:"allowPaths,omitempty"`
	// Max spend per period; empty = no spending
	SpendLimit []*Coin `json:"spendLimit,omitempty"`
	// Seconds; 0 = lifetime cap
	SpendPeriod int64 `json:"spendPeriod,string,omitempty"`
}

type CreateSessionRequest struct {
	GasFee    string `json:"gasFee,omitempty"`
	GasWanted int64  `json:"gasWanted,string,omitempty"`
	Memo      string `json:"memo,omitempty"`
	// The address of the creator (master) account to sign the transaction
	CreatorAddress []byte `json:"creatorAddress,omitempty"`
	// list of calls to make in one transaction
	Msgs []*MsgCreateSession `json:"Msgs,omitempty"` // proto json_name = "Msgs"
}

type CreateSessionResponse struct {
	Result []byte `json:"result,omitempty"`
	// The transaction hash
	Hash []byte `json:"hash,omitempty"`
	// The transaction height
	Height int64 `json:"height,string,omitempty"`
}

type MsgRevokeSession struct {
	// Full session public key
	SessionKey []byte `json:"sessionKey,omitempty"`
}

type RevokeSessionRequest struct {
	GasFee    string `json:"gasFee,omitempty"`
	GasWanted int64  `json:"gasWanted,string,omitempty"`
	Memo      string `json:"memo,omitempty"`
	// The address of the creator (master) account to sign the transaction
	CreatorAddress []byte `json:"creatorAddress,omitempty"`
	// list of calls to make in one transaction
	Msgs []*MsgRevokeSession `json:"Msgs,omitempty"` // proto json_name = "Msgs"
}

type RevokeSessionResponse struct {
	Result []byte `json:"result,omitempty"`
	// The transaction hash
	Hash []byte `json:"hash,omitempty"`
	// The transaction height
	Height int64 `json:"height,string,omitempty"`
}

type RevokeAllSessionsRequest struct {
	GasFee    string `json:"gasFee,omitempty"`
	GasWanted int64  `json:"gasWanted,string,omitempty"`
	Memo      string `json:"memo,omitempty"`
	// The address of the creator (master) account to sign the transaction
	CreatorAddress []byte `json:"creatorAddress,omitempty"`
}

type RevokeAllSessionsResponse struct {
	Result []byte `json:"result,omitempty"`
	// The transaction hash
	Hash []byte `json:"hash,omitempty"`
	// The transaction height
	Height int64 `json:"height,string,omitempty"`
}

type EstimateGasRequest struct {
	// The JSON encoding of the unsigned transaction (from MakeCallTx, etc.)
	TxJSON string `json:"txJson,omitempty"`
	// The address of the account that will sign the transaction
	Address []byte `json:"address,omitempty"`
	// The security margin to apply to the estimated gas amount.
	// This number represents a decimal numeral value with two decimals precision, without the decimal separator. E.g. 1 means 0.01 and 10000 means 100.00.
	// It will be multiplied by the estimated gas amount.
	SecurityMargin uint32 `json:"securityMargin,omitempty"`
	// The update boolean to update the gas wanted field in the transaction if true.
	UpdateTx bool `json:"updateTx,omitempty"`
}

type EstimateGasResponse struct {
	// The JSON encoding of the unsigned transaction
	TxJSON string `json:"txJson,omitempty"`
	// The estimated gas wanted for the transaction
	GasWanted int64 `json:"gasWanted,string,omitempty"`
}

type EstimateTxFeesRequest struct {
	// The JSON encoding of the unsigned transaction (from MakeCallTx, etc.)
	TxJSON string `json:"txJson,omitempty"`
	// The address of the account that will sign the transaction
	Address []byte `json:"address,omitempty"`
	// The security margin to apply to the estimated gas amount.
	// This number represents a decimal numeral value with two decimals precision, without the decimal separator. E.g. 1 means 0.01 and 10000 means 100.00.
	// It will be multiplied by the estimated gas amount.
	GasSecurityMargin uint32 `json:"gasSecurityMargin,omitempty"`
	// The security margin to apply to the gas price.
	// This number represents a decimal numeral value with two decimals precision, without the decimal separator. E.g. 1 means 0.01 and 10000 means 100.00.
	// It will be multiplied by the fetched gas price.
	GasPriceSecurityMargin uint32 `json:"gasPriceSecurityMargin,omitempty"`
	// The update boolean to update the gas wanted field in the transaction if true.
	UpdateTx bool `json:"updateTx,omitempty"`
}

type EstimateTxFeesResponse struct {
	// The JSON encoding of the unsigned transaction
	TxJSON string `json:"txJson,omitempty"`
	// The estimated gas wanted for the transaction
	GasWanted int64 `json:"gasWanted,string,omitempty"`
	// The estimated fee for the transaction
	GasFee *Coin `json:"gasFee,omitempty"`
	// The estimated storage delta for the transaction. Can be negative for "unlock"
	StorageDelta int64 `json:"storageDelta,string,omitempty"`
	// The estimated storage fee for the transaction. Does not include RefundWithheld. Can be negative for "unlock"
	StorageFee []*Coin `json:"storageFee,omitempty"`
	// The total transaction fee. Only includes fees in ugnot. Does not include RefundWithheld.
	TotalFee *Coin `json:"TotalFee,omitempty"` // proto json_name = "TotalFee"
}

type BroadcastTxCommitRequest struct {
	// The JSON encoding of the signed transaction (from SignTx)
	SignedTxJSON string `json:"tx_json,omitempty"` // proto json_name = "tx_json"
}

type BroadcastTxCommitResponse struct {
	Result []byte `json:"result,omitempty"`
	// The transaction hash
	Hash []byte `json:"hash,omitempty"`
	// The transaction height
	Height int64 `json:"height,string,omitempty"`
}

type AddressToBech32Request struct {
	Address []byte `json:"address,omitempty"`
}

type AddressToBech32Response struct {
	Bech32Address string `json:"bech32Address,omitempty"`
}

type AddressFromBech32Request struct {
	Bech32Address string `json:"bech32Address,omitempty"`
}

type AddressFromBech32Response struct {
	Address []byte `json:"address,omitempty"`
}

type AddressFromMnemonicRequest struct {
	Mnemonic string `json:"mnemonic,omitempty"`
}

type AddressFromMnemonicResponse struct {
	Address []byte `json:"address,omitempty"`
}

type ValidateMnemonicWordRequest struct {
	Word string `json:"word,omitempty"`
}

type ValidateMnemonicWordResponse struct {
	Valid bool `json:"valid,omitempty"`
}

type ValidateMnemonicPhraseRequest struct {
	Phrase string `json:"phrase,omitempty"`
}

type ValidateMnemonicPhraseResponse struct {
	Valid bool `json:"valid,omitempty"`
}

type PubKeyBytesFromBech32Request struct {
	Bech32PubKey string `json:"bech32PubKey,omitempty"`
}

type PubKeyBytesFromBech32Response struct {
	PubKeyBytes []byte `json:"pubKeyBytes,omitempty"`
}

type HelloRequest struct {
	Name string `json:"Name,omitempty"` // proto json_name = "Name"
}

type HelloResponse struct {
	Greeting string `json:"Greeting,omitempty"` // proto json_name = "Greeting"
}

type HelloStreamRequest struct {
	Name string `json:"Name,omitempty"` // proto json_name = "Name"
}

type HelloStreamResponse struct {
	Greeting string `json:"Greeting,omitempty"` // proto json_name = "Greeting"
}
