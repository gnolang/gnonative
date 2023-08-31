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

	client := NewClient()
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

func callCreateThread(cfg *accountAndTxCfg, boardId string, title string, body string) error {
	return cfg.Client.Call("gno.land/r/demo/boards", "CreateThread", []string{boardId, title, body}, cfg.GasFee, cfg.GasWanted, "")
}

func callCreateReply(cfg *accountAndTxCfg, boardId string, threadId string, postId string, body string) error {
	return cfg.Client.Call("gno.land/r/demo/boards", "CreateReply", []string{boardId, threadId, postId, body}, cfg.GasFee, cfg.GasWanted, "")
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

func CreateReply(message string, rootDir string) string {
	cfg, err := getAccountAndTxCfg(rootDir)
	if err != nil {
		return fmt.Sprintf("Error: unable to get config: %s", err.Error())
	}

	err = callCreateReply(cfg, "2", "1", "1", message)

	if err != nil {
		return fmt.Sprintf("Error: unable to exec call command: %s", err.Error())
	}

	return fmt.Sprintf("Posted: %s", message)
}
