package gnomobile

import (
	"fmt"

	"github.com/gnolang/gno/tm2/pkg/amino"
	rpc_client "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	ctypes "github.com/gnolang/gno/tm2/pkg/bft/rpc/core/types"
	"github.com/gnolang/gno/tm2/pkg/commands"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys"
	"github.com/gnolang/gno/tm2/pkg/errors"
	"github.com/gnolang/gno/tm2/pkg/std"
)

// From https://github.com/gnolang/gno/blob/master/tm2/pkg/crypto/keys/client/common.go
type BaseOptions struct {
	Home                  string
	Remote                string
	Quiet                 bool
	InsecurePasswordStdin bool
	Config                string
}

// From https://github.com/gnolang/gno/blob/master/tm2/pkg/crypto/keys/client/root.go
type baseCfg struct {
	BaseOptions
}

// From https://github.com/gnolang/gno/blob/master/tm2/pkg/crypto/keys/client/maketx.go
type makeTxCfg struct {
	rootCfg *baseCfg

	gasWanted int64
	gasFee    string
	memo      string

	broadcast bool
	chainID   string
}

// From https://github.com/gnolang/gno/blob/master/tm2/pkg/crypto/keys/client/query.go
type queryCfg struct {
	remote string

	data   string
	height int64
	prove  bool

	// internal
	path string
}

// From https://github.com/gnolang/gno/blob/master/tm2/pkg/crypto/keys/client/sign.go
type signCfg struct {
	kb keys.Keybase

	txPath        string
	chainID       string
	accountNumber uint64
	sequence      uint64
	showSignBytes bool

	// internal flags, when called programmatically
	nameOrBech32 string
	txJSON       []byte
	pass         string
}

// From https://github.com/gnolang/gno/blob/master/tm2/pkg/crypto/keys/client/broadcast.go
type broadcastCfg struct {
	remote string

	dryRun bool

	// internal
	tx *std.Tx
}

// From https://github.com/gnolang/gno/blob/master/tm2/pkg/crypto/keys/client/call.go
type callCfg struct {
	rootCfg *makeTxCfg

	send     string
	pkgPath  string
	funcName string
	args     commands.StringArr
}

// From https://github.com/gnolang/gno/blob/master/tm2/pkg/crypto/keys/client/call.go
func execCall(cfg *callCfg, nameOrBech32 string, password string) error {
	client := NewClient(cfg.rootCfg.rootCfg.Remote, cfg.rootCfg.chainID)
	client.SetAccount(nameOrBech32, password)
	if err := client.SetKeyBaseFromDir(cfg.rootCfg.rootCfg.Home); err != nil {
		return err
	}
	r := client.NewRequest("call")
	r.StringOption("pkgpath", cfg.pkgPath)
	r.StringOption("func", cfg.funcName)
	r.StringOption("gas-fee", cfg.rootCfg.gasFee)
	r.Int64Option("gas-wanted", cfg.rootCfg.gasWanted)
	for _, arg := range cfg.args {
		r.Argument(arg)
	}

	return r.Send()
}

// From https://github.com/gnolang/gno/blob/master/tm2/pkg/crypto/keys/client/query.go
func queryHandler(cfg *queryCfg) (*ctypes.ResultABCIQuery, error) {
	remote := cfg.remote
	if remote == "" || remote == "y" {
		return nil, errors.New("missing remote url")
	}

	data := []byte(cfg.data)
	opts2 := rpc_client.ABCIQueryOptions{
		// Height: height, XXX
		// Prove: false, XXX
	}
	cli := rpc_client.NewHTTP(remote, "/websocket")
	qres, err := cli.ABCIQueryWithOptions(
		cfg.path, data, opts2)
	if err != nil {
		return nil, errors.Wrap(err, "querying")
	}

	return qres, nil
}

// From https://github.com/gnolang/gno/blob/master/tm2/pkg/crypto/keys/client/broadcast.go
func broadcastHandler(cfg *broadcastCfg) (*ctypes.ResultBroadcastTxCommit, error) {
	if cfg.tx == nil {
		return nil, errors.New("invalid tx")
	}

	remote := cfg.remote
	if remote == "" || remote == "y" {
		return nil, errors.New("missing remote url")
	}

	bz, err := amino.Marshal(cfg.tx)
	if err != nil {
		return nil, errors.Wrap(err, "remarshaling tx binary bytes")
	}

	cli := rpc_client.NewHTTP(remote, "/websocket")

	/*
		if cfg.dryRun {
			return simulateTx(cli, bz)
		}
	*/

	bres, err := cli.BroadcastTxCommit(bz)
	if err != nil {
		return nil, errors.Wrap(err, "broadcasting bytes")
	}

	return bres, nil
}

// From https://github.com/gnolang/gno/blob/master/tm2/pkg/crypto/keys/client/sign.go
// (Even though the original SignHandler is public, its argument signCfg is private.)
func SignHandler(cfg *signCfg) (*std.Tx, error) {
	var err error
	var tx std.Tx

	if cfg.txJSON == nil {
		return nil, errors.New("invalid tx content")
	}

	err = amino.UnmarshalJSON(cfg.txJSON, &tx)
	if err != nil {
		return nil, err
	}

	// fill tx signatures.
	signers := tx.GetSigners()
	if tx.Signatures == nil {
		for range signers {
			tx.Signatures = append(tx.Signatures, std.Signature{
				PubKey:    nil, // zero signature
				Signature: nil, // zero signature
			})
		}
	}

	// validate document to sign.
	err = tx.ValidateBasic()
	if err != nil {
		return nil, err
	}

	// derive sign doc bytes.
	chainID := cfg.chainID
	accountNumber := cfg.accountNumber
	sequence := cfg.sequence
	signbz := tx.GetSignBytes(chainID, accountNumber, sequence)
	if cfg.showSignBytes {
		fmt.Printf("sign bytes: %X\n", signbz)
		return nil, nil
	}

	sig, pub, err := cfg.kb.Sign(cfg.nameOrBech32, cfg.pass, signbz)
	if err != nil {
		return nil, err
	}
	addr := pub.Address()
	found := false
	for i := range tx.Signatures {
		// override signature for matching slot.
		if signers[i] == addr {
			found = true
			tx.Signatures[i] = std.Signature{
				PubKey:    pub,
				Signature: sig,
			}
		}
	}
	if !found {
		return nil, errors.New(
			fmt.Sprintf("addr %v (%s) not in signer set", addr, cfg.nameOrBech32),
		)
	}

	return &tx, nil
}
