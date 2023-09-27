package gnomobiletypes

type SetRemoteRequest struct {
	Remote string `json:"remote" yaml:"remote"`
}

type SetRemoteReply struct {
}

type SetChainIDRequest struct {
	ChainID string `json:"chain_id" yaml:"chain_id"`
}

type SetChainIDReply struct {
}

type SetNameOrBech32Request struct {
	NameOrBech32 string `json:"name_or_bech32" yaml:"name_or_bech32"`
}

type SetNameOrBech32Reply struct {
}

type SetPasswordRequest struct {
	Password string `json:"password" yaml:"password"`
}

type SetPasswordReply struct {
}

type GenerateRecoveryPhraseRequest struct {
}

type GenerateRecoveryPhraseReply struct {
	Phrase string `json:"phrase" yaml:"phrase"`
}

type QueryRequest struct {
	// Example: "vm/qrender"
	Path string `json:"path" yaml:"path"`
	// Example: "gno.land/r/demo/boards\ntestboard"
	Data string `json:"data" yaml:"data"`
}

type QueryReply struct {
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

type CallReply struct {
	Result []byte `json:"result" yaml:"result"`
}
