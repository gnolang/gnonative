package gnomobile

import (
	"fmt"

	"github.com/gnolang/gno/tm2/pkg/amino"
	rpc_client "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	ctypes "github.com/gnolang/gno/tm2/pkg/bft/rpc/core/types"
	"github.com/gnolang/gno/tm2/pkg/commands"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys/client"
	"github.com/gnolang/gno/tm2/pkg/errors"
	"github.com/gnolang/gno/tm2/pkg/sdk/vm"
	"github.com/gnolang/gno/tm2/pkg/std"
)

// From https://github.com/gnolang/gno/blob/master/tm2/pkg/crypto/keys/client/root.go
type baseCfg struct {
	client.BaseOptions
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
	rootCfg *baseCfg

	data   string
	height int64
	prove  bool

	// internal
	path string
}

// From https://github.com/gnolang/gno/blob/master/tm2/pkg/crypto/keys/client/sign.go
type signCfg struct {
	rootCfg *baseCfg

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
	rootCfg *baseCfg

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
	if cfg.pkgPath == "" {
		return errors.New("pkgpath not specified")
	}
	if cfg.funcName == "" {
		return errors.New("func not specified")
	}
	if cfg.rootCfg.gasWanted == 0 {
		return errors.New("gas-wanted not specified")
	}
	if cfg.rootCfg.gasFee == "" {
		return errors.New("gas-fee not specified")
	}

	// read statement.
	fnc := cfg.funcName

	// read account pubkey.
	kb, err := keys.NewKeyBaseFromDir(cfg.rootCfg.rootCfg.Home)
	if err != nil {
		return err
	}
	info, err := kb.GetByNameOrAddress(nameOrBech32)
	if err != nil {
		return err
	}
	caller := info.GetAddress()
	// info.GetPubKey()

	// Parse send amount.
	send, err := std.ParseCoins(cfg.send)
	if err != nil {
		return errors.Wrap(err, "parsing send coins")
	}

	// parse gas wanted & fee.
	gaswanted := cfg.rootCfg.gasWanted
	gasfee, err := std.ParseCoin(cfg.rootCfg.gasFee)
	if err != nil {
		return errors.Wrap(err, "parsing gas fee coin")
	}

	// construct msg & tx and marshal.
	msg := vm.MsgCall{
		Caller:  caller,
		Send:    send,
		PkgPath: cfg.pkgPath,
		Func:    fnc,
		Args:    cfg.args,
	}
	tx := std.Tx{
		Msgs:       []std.Msg{msg},
		Fee:        std.NewFee(gaswanted, gasfee),
		Signatures: nil,
		Memo:       cfg.rootCfg.memo,
	}

	if cfg.rootCfg.broadcast {
		err := signAndBroadcast(cfg.rootCfg, tx, kb, nameOrBech32, password)
		if err != nil {
			return err
		}
	} else {
		errors.New(string(amino.MustMarshalJSON(tx)))
	}
	return nil
}

// From https://github.com/gnolang/gno/blob/master/tm2/pkg/crypto/keys/client/addpkg.go
func signAndBroadcast(
	cfg *makeTxCfg,
	tx std.Tx,
	kb keys.Keybase,
	nameOrBech32 string,
	password string,
) error {
	baseopts := cfg.rootCfg
	txopts := cfg

	// query account
	kb, err := keys.NewKeyBaseFromDir(baseopts.Home)
	if err != nil {
		return err
	}
	info, err := kb.GetByNameOrAddress(nameOrBech32)
	if err != nil {
		return err
	}
	accountAddr := info.GetAddress()

	qopts := &queryCfg{
		rootCfg: baseopts,
		path:    fmt.Sprintf("auth/accounts/%s", accountAddr),
	}
	qres, err := queryHandler(qopts)
	if err != nil {
		return errors.Wrap(err, "query account")
	}
	var qret struct{ BaseAccount std.BaseAccount }
	err = amino.UnmarshalJSON(qres.Response.Data, &qret)
	if err != nil {
		return err
	}

	// sign tx
	accountNumber := qret.BaseAccount.AccountNumber
	sequence := qret.BaseAccount.Sequence
	sopts := &signCfg{
		rootCfg:       baseopts,
		sequence:      sequence,
		accountNumber: accountNumber,
		chainID:       txopts.chainID,
		nameOrBech32:  nameOrBech32,
		txJSON:        amino.MustMarshalJSON(tx),
		pass:          password,
	}

	signedTx, err := SignHandler(sopts)
	if err != nil {
		return errors.Wrap(err, "sign tx")
	}

	// broadcast signed tx
	bopts := &broadcastCfg{
		rootCfg: baseopts,
		tx:      signedTx,
	}
	bres, err := broadcastHandler(bopts)
	if err != nil {
		return errors.Wrap(err, "broadcast tx")
	}
	if bres.CheckTx.IsErr() {
		return errors.Wrap(bres.CheckTx.Error, "check transaction failed: log:%s", bres.CheckTx.Log)
	}
	if bres.DeliverTx.IsErr() {
		return errors.Wrap(bres.DeliverTx.Error, "deliver transaction failed: log:%s", bres.DeliverTx.Log)
	}

	return nil
}

// From https://github.com/gnolang/gno/blob/master/tm2/pkg/crypto/keys/client/query.go
func queryHandler(cfg *queryCfg) (*ctypes.ResultABCIQuery, error) {
	remote := cfg.rootCfg.Remote
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

	remote := cfg.rootCfg.Remote
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

	kb, err := keys.NewKeyBaseFromDir(cfg.rootCfg.Home)
	if err != nil {
		return nil, err
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

	sig, pub, err := kb.Sign(cfg.nameOrBech32, cfg.pass, signbz)
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
