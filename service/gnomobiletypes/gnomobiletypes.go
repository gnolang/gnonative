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

type DeleteAccountRequest struct {
	NameOrBech32 string `json:"name_or_bech32" yaml:"name_or_bech32"`
	Password     string `json:"password" yaml:"password"`
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
