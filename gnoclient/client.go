// This is a temporary shadow of the Client struct to be provided by https://github.com/gnolang/gno/pull/1047 .

package gnoclient

import (
	"fmt"

	"github.com/gnolang/gno/tm2/pkg/amino"
	rpc_client "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	ctypes "github.com/gnolang/gno/tm2/pkg/bft/rpc/core/types"
	"github.com/gnolang/gno/tm2/pkg/commands"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys"
	"github.com/gnolang/gno/tm2/pkg/errors"
	"github.com/gnolang/gno/tm2/pkg/sdk/vm"
	"github.com/gnolang/gno/tm2/pkg/std"
)

// Client represents the Gno.land RPC API client.
type Client struct {
	remote       string
	chainID      string
	keybase      keys.Keybase
	nameOrBech32 string
	password     string
}

type Opts struct {
	Remote       string
	ChainID      string
	NameOrBech32 string
	Password     string
}

func (o *Opts) ApplyDefaults() {
	if o.Remote == "" {
		o.Remote = "127.0.0.1:26657"
	}

	if o.ChainID == "" {
		o.ChainID = "dev"
	}
}

func NewClient(opts Opts) *Client {
	opts.ApplyDefaults()

	return &Client{
		remote:       opts.Remote,
		chainID:      opts.ChainID,
		nameOrBech32: opts.NameOrBech32,
		password:     opts.Password,
	}
}

func (c *Client) SetRemote(remote string) {
	c.remote = remote
}

func (c *Client) SetChainID(chainID string) {
	c.chainID = chainID
}

func (c *Client) InitKeyBaseFromDir(rootDir string) error {
	var err error
	if c.keybase, err = keys.NewKeyBaseFromDir(rootDir); err != nil {
		return err
	}
	return nil
}

func (c *Client) SetNameOrBech32(nameOrBech32 string) {
	c.nameOrBech32 = nameOrBech32
}

func (c *Client) SetPassword(password string) {
	c.password = password
}

func (c *Client) GetKeys() ([]keys.Info, error) {
	keyList, err := c.keybase.List()
	if err != nil {
		return nil, err
	}
	return keyList, nil
}

func (c *Client) GetKeyByNameOrBech32(nameOrBech32 string) (keys.Info, error) {
	return c.keybase.GetByNameOrAddress(nameOrBech32)
}

func (c *Client) CreateAccount(nameOrBech32 string, mnemonic string, bip39Passwd string, password string, account uint32, index uint32) (keys.Info, error) {
	info, err := c.keybase.CreateAccount(nameOrBech32, mnemonic, bip39Passwd, password, account, index)
	if err != nil {
		return nil, err
	}

	return info, err
}

// TODO: port existing code, i.e. faucet?
// TODO: create right now a tm2 generic go client and a gnovm generic go client?
// TODO: Command: Send
// TODO: Command: AddPkg
// TODO: Command: Query
// TODO: Command: Eval
// TODO: Command: Exec
// TODO: Command: Package
// TODO: Command: QFile
// TODO: examples and unit tests
// TODO: Mock
// TODO: alternative configuration (pass existing websocket?)
// TODO: minimal go.mod to make it light to import

func (c *Client) Call(pkgPath string, fnc string, args commands.StringArr, gasFee string, gasWanted int64, send string, nameOrBech32 string, password string) error {
	info, err := c.keybase.GetByNameOrAddress(nameOrBech32)
	if err != nil {
		return err
	}
	caller := info.GetAddress()

	// Parse send amount.
	sendCoins, err := std.ParseCoins(send)
	if err != nil {
		return errors.Wrap(err, "parsing send coins")
	}

	// parse gas wanted & fee.
	gasFeeCoins, err := std.ParseCoin(gasFee)
	if err != nil {
		return errors.Wrap(err, "parsing gas fee coin")
	}

	// construct msg & tx and marshal.
	msg := vm.MsgCall{
		Caller:  caller,
		Send:    sendCoins,
		PkgPath: pkgPath,
		Func:    fnc,
		Args:    args,
	}
	tx := std.Tx{
		Msgs:       []std.Msg{msg},
		Fee:        std.NewFee(gasWanted, gasFeeCoins),
		Signatures: nil,
		Memo:       "",
	}

	qopts := &queryCfg{
		remote: c.remote,
		path:   fmt.Sprintf("auth/accounts/%s", caller),
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
		kb:            c.keybase,
		sequence:      sequence,
		accountNumber: accountNumber,
		chainID:       c.chainID,
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
		remote: c.remote,
		tx:     signedTx,
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
