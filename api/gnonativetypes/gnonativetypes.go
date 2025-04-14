package gnonativetypes

type SetRemoteRequest struct {
	Remote string `json:"remote" yaml:"remote"`
}

type SetRemoteResponse struct{}

type GetRemoteRequest struct{}

type GetRemoteResponse struct {
	Remote string `json:"remote" yaml:"remote"`
}

type SetChainIDRequest struct {
	ChainID string `json:"chain_id" yaml:"chain_id"`
}

type SetChainIDResponse struct{}

type GetChainIDRequest struct{}

type GetChainIDResponse struct {
	ChainID string `json:"chain_id" yaml:"chain_id"`
}

type SetPasswordRequest struct {
	Password string `json:"password" yaml:"password"`
	// The address of the account to set the password
	Address []byte `json:"address" yaml:"address"`
}

type SetPasswordResponse struct{}

type RotatePasswordRequest struct {
	NewPassword string `json:"new_password" yaml:"new_password"`
	// The addresses of the account to rotate the password
	Addresses [][]byte `json:"addresses" yaml:"addresses"`
}

type RotatePasswordResponse struct{}

type GenerateRecoveryPhraseRequest struct{}

type GenerateRecoveryPhraseResponse struct {
	Phrase string `json:"phrase" yaml:"phrase"`
}

type KeyInfo struct {
	// 0: local, 1: ledger, 2: offline, 3: multi
	Type    uint32 `json:"type" yaml:"type"`
	Name    string `json:"name" yaml:"name"`
	PubKey  []byte `json:"pub_key" yaml:"pub_key"`
	Address []byte `json:"address" yaml:"address"`
}

// Coin holds some amount of one currency.
// A negative amount is invalid.
type Coin struct {
	// Example: "ugnot"
	Denom  string `json:"denom"`
	Amount int64  `json:"amount"`
}

type BaseAccount struct {
	Address       []byte `json:"address" yaml:"address"`
	Coins         []Coin `json:"coins" yaml:"coins"`
	PubKey        []byte `json:"pub_key" yaml:"pub_key"`
	AccountNumber uint64 `json:"account_number" yaml:"account_number"`
	Sequence      uint64 `json:"sequence" yaml:"sequence"`
}

type ListKeyInfoRequest struct{}

type ListKeyInfoResponse struct {
	Keys []*KeyInfo `json:"key_info" yaml:"key_info"`
}

type GetKeyInfoByNameRequest struct {
	Name string `json:"name" yaml:"name"`
}

type HasKeyByNameRequest struct {
	Name string `json:"name" yaml:"name"`
}

type HasKeyByNameResponse struct {
	Has bool `json:"has" yaml:"has"`
}

type HasKeyByAddressRequest struct {
	Address []byte `json:"address" yaml:"address"`
}

type HasKeyByAddressResponse struct {
	Has bool `json:"has" yaml:"has"`
}

type HasKeyByNameOrAddressRequest struct {
	NameOrBech32 string `json:"name_or_bech32" yaml:"name_or_bech32"`
}

type HasKeyByNameOrAddressResponse struct {
	Has bool `json:"has" yaml:"has"`
}

type GetKeyInfoByNameResponse struct {
	Key *KeyInfo `json:"key_info" yaml:"key_info"`
}

type GetKeyInfoByAddressRequest struct {
	Address []byte `json:"address" yaml:"address"`
}

type GetKeyInfoByAddressResponse struct {
	Key *KeyInfo `json:"key_info" yaml:"key_info"`
}

type GetKeyInfoByNameOrAddressRequest struct {
	NameOrBech32 string `json:"name_or_bech32" yaml:"name_or_bech32"`
}

type GetKeyInfoByNameOrAddressResponse struct {
	Key *KeyInfo `json:"key_info" yaml:"key_info"`
}

type CreateAccountRequest struct {
	NameOrBech32 string `json:"name_or_bech32" yaml:"name_or_bech32"`
	Mnemonic     string `json:"mnemonic" yaml:"mnemonic"`
	Bip39Passwd  string `json:"bip39_passwd" yaml:"bip39_passwd"`
	Password     string `json:"password" yaml:"password"`
	Account      uint32 `json:"account" yaml:"account"`
	Index        uint32 `json:"index" yaml:"index"`
}

type CreateAccountResponse struct {
	Key *KeyInfo `json:"key_info" yaml:"key_info"`
}

type ActivateAccountRequest struct {
	NameOrBech32 string `json:"name_or_bech32" yaml:"name_or_bech32"`
}

type ActivateAccountResponse struct {
	Key *KeyInfo `json:"key_info" yaml:"key_info"`
	// True if the password has been set. If false, then call SetPassword.
	HasPassword bool `json:"has_password" yaml:"has_password"`
}

type GetActivatedAccountRequest struct {
	Address []byte `json:"address" yaml:"address"`
}

type GetActivatedAccountResponse struct {
	Key *KeyInfo `json:"key_info" yaml:"key_info"`
	// True if the password has been set. If false, then call SetPassword.
	HasPassword bool `json:"has_password" yaml:"has_password"`
}

type QueryAccountRequest struct {
	Address []byte `json:"address" yaml:"address"`
}

type QueryAccountResponse struct {
	AccountInfo *BaseAccount `json:"account_info" yaml:"account_info"`
}

type DeleteAccountRequest struct {
	NameOrBech32 string `json:"name_or_bech32" yaml:"name_or_bech32"`
	Password     string `json:"password" yaml:"password"`
	SkipPassword bool   `json:"skip_password" yaml:"skip_password"`
}

type DeleteAccountResponse struct{}

type QueryRequest struct {
	// Example: "vm/qrender"
	Path string `json:"path" yaml:"path"`
	// Example: "gno.land/r/demo/boards\ntestboard"
	Data []byte `json:"data" yaml:"data"`
}

type QueryResponse struct {
	Result []byte `json:"result" yaml:"result"`
}

type RenderRequest struct {
	// Example: "gno.land/r/demo/boards"
	PackagePath string `json:"package_path" yaml:"package_path"`
	// Example: "testboard/1"
	Args string `json:"args" yaml:"args"`
}

type RenderResponse struct {
	// The Render function result (typically markdown)
	Result string `json:"result" yaml:"result"`
}

type QEvalRequest struct {
	// Example: "gno.land/r/demo/boards"
	PackagePath string `json:"package_path" yaml:"package_path"`
	// Example: "GetBoardIDFromName(\"testboard\")"
	Expression string `json:"expression" yaml:"expression"`
}

type QEvalResponse struct {
	// A typed expression like "(1 gno.land/r/demo/boards.BoardID)\n(true bool)"
	Result string `json:"result" yaml:"result"`
}

type MsgCall struct {
	// Example: "gno.land/r/demo/boards"
	PackagePath string `json:"package_path" yaml:"package_path"`
	// Example: "CreateReply"
	Fnc string `json:"fnc" yaml:"fnc"`
	// list of arguments specific to the function
	// Example: ["1", "1", "2", "my reply"]
	Args []string `json:"args" yaml:"args"`
	Send []Coin   `json:"send" yaml:"send"`
}

type CallRequest struct {
	GasFee    string `json:"gas_fee" yaml:"gas_fee"`
	GasWanted int64  `json:"gas_wanted" yaml:"gas_wanted"`
	Memo      string `json:"memo" yaml:"memo"`
	// The address of the account to sign the transaction
	CallerAddress []byte `json:"caller_address" yaml:"caller_address"`
	// list of calls to make in one transaction
	Msgs []MsgCall
}

type CallResponse struct {
	Result []byte `json:"result" yaml:"result"`
	// The transaction hash
	Hash []byte `json:"hash" yaml:"hash"`
	// The transaction height
	Height int64 `json:"height" yaml:"height"`
}

type MsgSend struct {
	// Example: The response of calling AddressFromBech32 with
	// "g1juz2yxmdsa6audkp6ep9vfv80c8p5u76e03vvh"
	ToAddress []byte `json:"to_address" yaml:"to_address"`
	// Example: [ {Denom: "ugnot", Amount: 1000} ]
	Amount []Coin `json:"amount" yaml:"amount"`
}

type SendRequest struct {
	GasFee    string `json:"gas_fee" yaml:"gas_fee"`
	GasWanted int64  `json:"gas_wanted" yaml:"gas_wanted"`
	// Memo is optional
	Memo string `json:"memo" yaml:"memo"`
	// The address of the account to sign the transaction
	CallerAddress []byte `json:"caller_address" yaml:"caller_address"`
	// list of send operations to make in one transaction
	Msgs []MsgSend
}

type SendResponse struct {
	// The transaction hash
	Hash []byte `json:"hash" yaml:"hash"`
	// The transaction height
	Height int64 `json:"height" yaml:"height"`
}

type MsgRun struct {
	// The code for the script package. Must have main().
	// Example: "package main\nfunc main() {\n  println(\"Hello\")\n}"
	Package string `json:"package" yaml:"package"`
	// Optional. Example: "1000ugnot"
	Send string `json:"send" yaml:"send"`
}

type RunRequest struct {
	GasFee    string `json:"gas_fee" yaml:"gas_fee"`
	GasWanted int64  `json:"gas_wanted" yaml:"gas_wanted"`
	// Memo is optional
	Memo string `json:"memo" yaml:"memo"`
	// The address of the account to sign the transaction
	CallerAddress []byte `json:"caller_address" yaml:"caller_address"`
	// list of run operations to make in one transaction
	Msgs []MsgRun
}

type RunResponse struct {
	// The "console" output from the run
	Result string `json:"result" yaml:"result"`
	// The transaction hash
	Hash []byte `json:"hash" yaml:"hash"`
	// The transaction height
	Height int64 `json:"height" yaml:"height"`
}

type MakeTxResponse struct {
	// The JSON encoding of the unsigned transaction
	TxJSON string `json:"tx_json" yaml:"tx_json"`
}

type SignTxRequest struct {
	// The JSON encoding of the unsigned transaction (from MakeCallTx, etc.)
	TxJSON string `json:"tx_json" yaml:"tx_json"`
	// The address of the account to sign the transaction
	Address []byte `json:"address" yaml:"address"`
	// The signer's account number on the blockchain. If 0 then query the blockchain for the value.
	AccountNumber uint64 `json:"account_number" yaml:"account_number"`
	// The sequence number of the signer's transactions on the blockchain. If 0 then query the blockchain for the value.
	SequenceNumber uint64 `json:"sequence_number" yaml:"sequence_number"`
}

type SignTxResponse struct {
	// The JSON encoding of the signed transaction (to use in BroadcastTx)
	SignedTxJSON string `json:"tx_json" yaml:"tx_json"`
}

type EstimateGasRequest struct {
	// The JSON encoding of the unsigned transaction (from MakeCallTx, etc.)
	TxJSON string `json:"tx_json" yaml:"tx_json"`
	// The address of the account to sign the transaction
	Address []byte `json:"address" yaml:"address"`
	// The security margin to apply to the estimated gas amount.
	// This number represents a decimal numeral value with two decimals precision, without the decimal separator. E.g. 1 means 0.01 and 10000 means 100.00.
	// It will be multiplied by the estimated gas amount.
	SecurityMargin uint32 `json:"security_margin" yaml:"security_margin"`
	// The update boolean to update the gas wanted field in the transaction if true.
	UpdateTx bool `json:"update_tx" yaml:"update_tx"`
	// The signer's account number on the blockchain. If 0 then query the blockchain for the value.
	AccountNumber uint64 `json:"account_number" yaml:"account_number"`
	// The sequence number of the signer's transactions on the blockchain. If 0 then query the blockchain for the value.
	SequenceNumber uint64 `json:"sequence_number" yaml:"sequence_number"`
}

type EstimateGasResponse struct {
	// The JSON encoding of the unsigned transaction
	TxJSON string `json:"tx_json" yaml:"tx_json"`
	// The estimated gas wanted for the transaction
	GasWanted int64 `json:"gas_wanted" yaml:"gas_wanted"`
}

type BroadcastTxCommitRequest struct {
	// The JSON encoding of the signed transaction (from SignTx)
	SignedTxJSON string `json:"tx_json" yaml:"tx_json"`
}

type BroadcastTxCommitResponse struct {
	Result []byte `json:"result" yaml:"result"`
	// The transaction hash
	Hash []byte `json:"hash" yaml:"hash"`
	// The transaction height
	Height int64 `json:"height" yaml:"height"`
}

type AddressToBech32Request struct {
	Address []byte `json:"address" yaml:"address"`
}

type AddressToBech32Response struct {
	Bech32Address string `json:"bech32_address" yaml:"bech32_address"`
}

type AddressFromBech32Request struct {
	Bech32Address string `json:"bech32_address" yaml:"bech32_address"`
}

type AddressFromBech32Response struct {
	Address []byte `json:"address" yaml:"address"`
}

type AddressFromMnemonicRequest struct {
	Mnemonic string `json:"mnemonic" yaml:"mnemonic"`
}

type AddressFromMnemonicResponse struct {
	Address []byte `json:"address" yaml:"address"`
}

type HelloRequest struct {
	Name string
}

type HelloResponse struct {
	Greeting string
}

type HelloStreamRequest struct {
	Name string
}

type HelloStreamResponse struct {
	Greeting string
}
