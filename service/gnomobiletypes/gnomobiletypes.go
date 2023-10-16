package gnomobiletypes

type SetRemoteRequest struct {
	Remote string `json:"remote" yaml:"remote"`
}

type SetRemoteResponse struct {
}

type SetChainIDRequest struct {
	ChainID string `json:"chain_id" yaml:"chain_id"`
}

type SetChainIDResponse struct {
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

type ListKeyInfoRequest struct {
}

type ListKeyInfoResponse struct {
	Keys []*KeyInfo `json:"key_info" yaml:"key_info"`
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
}

type GetActiveAccountRequest struct {
}

type GetActiveAccountResponse struct {
	Key *KeyInfo `json:"key_info" yaml:"key_info"`
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

type HelloRequest struct {
	Name string
}

type HelloResponse struct {
	Greeting string
}
