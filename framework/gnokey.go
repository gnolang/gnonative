package gnomobile

import (
	"encoding/json"
	"fmt"
)

type PromiseBlock interface {
	CallResolve(reply string)
	CallReject(error error)
}

type accountAndTxCfg struct {
	Client *Client

	GasWanted int64
	GasFee    string
	Mnemonic  string
}

var client *Client = NewClient()

func ClientExec(command string, rootDir string) string {
	cfg, err := getAccountAndTxCfg(rootDir)
	if err != nil {
		return fmt.Sprintf("Error: unable to get config: %s", err.Error())
	}

	type Call struct {
		PackagePath string
		Fnc         string
		Args        []string
		GasFee      string
		GasWanted   int
	}

	var call Call
	err = json.Unmarshal([]byte(command), &call)
	if err == nil {
		err = cfg.Client.Call(
			call.PackagePath, call.Fnc, call.Args, call.GasFee, int64(call.GasWanted), "")
		if err != nil {
			return fmt.Sprintf("Error: unable to exec call command: %s", err.Error())
		}
		return fmt.Sprintf("Posted: %s", call.Args[len(call.Args)-1])
	}

	return "ERROR: Unrecognized json message"
}

func CreateDefaultAccount(rootDir string) error {
	cfg, err := getAccountAndTxCfg(rootDir)
	if err != nil {
		return err
	}

	keyCount, err := cfg.Client.GetKeyCount()
	if err != nil {
		return err
	}
	if keyCount > 0 {
		// Assume the account with cfg.KeyName is already created.
		return nil
	}

	return cfg.Client.CreateAccount(cfg.Mnemonic, "", 0, 0)
}

func getAccountAndTxCfg(rootDir string) (*accountAndTxCfg, error) {
	dataDir := rootDir + "/data"
	remote := "testnet.gno.berty.io:26657"
	chainID := "dev"
	keyName := "jefft0"
	password := "password"
	mnemonic := "enable until hover project know foam join table educate room better scrub clever powder virus army pitch ranch fix try cupboard scatter dune fee"

	client.SetRemote(remote, chainID)
	client.SetAccount(keyName, password)
	if err := client.SetKeyBaseFromDir(dataDir); err != nil {
		return nil, err
	}

	return &accountAndTxCfg{
		Client:    client,
		GasWanted: 2000000,
		GasFee:    "1000000ugnot",
		Mnemonic:  mnemonic,
	}, nil
}

func ExportJsonConfig(rootDir string) string {
	cfg, err := getAccountAndTxCfg(rootDir)
	if err != nil {
		return fmt.Sprintf("Error: unable make config: %s", err.Error())
	}
	config, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Sprintf("Error: unable load config: %s", err.Error())
	}
	return string(config)
}
