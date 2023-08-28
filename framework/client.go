// This is a temporary shadow of the Client struct to be provided by https://github.com/gnolang/gno/pull/1047 .

package gnomobile

import (
	"fmt"

	"github.com/gnolang/gno/tm2/pkg/amino"
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

func NewClient() *Client {
	// Set defaults.
	return &Client{
		remote:  "127.0.0.1:26657",
		chainID: "dev",
	}
}

func (c *Client) SetRemote(remote string, chainID string) {
	c.remote = remote
	c.chainID = chainID
}

func (c *Client) SetAccount(nameOrBech32 string, password string) {
	c.nameOrBech32 = nameOrBech32
	c.password = password
}

// Create an account using the nameOrBech32 and password from SetAccount.
func (c *Client) CreateAccount(mnemonic string, bip39Passwd string, account int, index int) error {
	_, err := c.keybase.CreateAccount(c.nameOrBech32, mnemonic, bip39Passwd, c.password, uint32(account), uint32(index))
	return err
}

func (c *Client) SetKeyBaseFromDir(rootDir string) error {
	var err error
	if c.keybase, err = keys.NewKeyBaseFromDir(rootDir); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetKeyCount() (int, error) {
	keyList, err := c.keybase.List()
	if err != nil {
		return 0, err
	}
	return len(keyList), nil
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

func (c *Client) Call(pkgPath string, fnc string, args commands.StringArr, gasFee string, gasWanted int64, send string) error {
	info, err := c.keybase.GetByNameOrAddress(c.nameOrBech32)
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
		nameOrBech32:  c.nameOrBech32,
		txJSON:        amino.MustMarshalJSON(tx),
		pass:          c.password,
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
