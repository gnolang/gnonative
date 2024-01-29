package gnonativetypes

type SetRemoteRequest struct {
	Remote string `json:"remote" yaml:"remote"`
}

type SetRemoteResponse struct {
}

type GetRemoteRequest struct {
}

type GetRemoteResponse struct {
	Remote string `json:"remote" yaml:"remote"`
}

type SetChainIDRequest struct {
	ChainID string `json:"chain_id" yaml:"chain_id"`
}

type SetChainIDResponse struct {
}

type GetChainIDRequest struct {
}

type GetChainIDResponse struct {
	ChainID string `json:"chain_id" yaml:"chain_id"`
}

type SetPasswordRequest struct {
	Password string `json:"password" yaml:"password"`
}

type SetPasswordResponse struct {
}

type GenerateRecoveryPhraseRequest struct {
}

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

type ListKeyInfoRequest struct {
}

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

type SelectAccountRequest struct {
	NameOrBech32 string `json:"name_or_bech32" yaml:"name_or_bech32"`
}

type SelectAccountResponse struct {
	Key *KeyInfo `json:"key_info" yaml:"key_info"`
	// True if the password has been set. If false, then call SetPassword.
	HasPassword bool `json:"has_password" yaml:"has_password"`
}

type GetActiveAccountRequest struct {
}

type GetActiveAccountResponse struct {
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

type DeleteAccountResponse struct {
}

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

type CallRequest struct {
	// Example: "gno.land/r/demo/boards"
	PackagePath string `json:"package_path" yaml:"package_path"`
	// Example: "CreateReply"
	Fnc string `json:"fnc" yaml:"fnc"`
	// list of arguments specific to the function
	Args      []string `json:"args" yaml:"args"`
	GasFee    string   `json:"gas_fee" yaml:"gas_fee"`
	GasWanted int64    `json:"gas_wanted" yaml:"gas_wanted"`
	Send      string   `json:"send" yaml:"send"`
	Memo      string   `json:"memo" yaml:"memo"`
}

type CallResponse struct {
	Result []byte `json:"result" yaml:"result"`
}

type RunRequest struct {
	// The code for the script package. Must have main().
	// Example: "package main\nfunc main() {\n  println(\"Hello\")\n}"
	Package   string `json:"package" yaml:"package"`
	GasFee    string `json:"gas_fee" yaml:"gas_fee"`
	GasWanted int64  `json:"gas_wanted" yaml:"gas_wanted"`
	// Memo is optional
	Memo string `json:"memo" yaml:"memo"`
}

type RunResponse struct {
	// The "console" output from the run
	Result string `json:"result" yaml:"result"`
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
