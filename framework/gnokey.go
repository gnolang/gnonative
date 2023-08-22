package gnomobile

import (
	"encoding/json"
	"fmt"

	"github.com/gnolang/gno/tm2/pkg/crypto/keys"
)

type PromiseBlock interface {
	CallResolve(reply string)
	CallReject(error error)
}

type accountAndTxCfg struct {
	TxCfg *makeTxCfg

	KeyName  string
	Password string
}

func Hello(rootDir string) string {
	cfg := getAccountAndTxCfg(rootDir)

	// Debug: We should only have to do this once. It seems that the Keybase dir is deleted when we reinstall the app.
	kb, err := keys.NewKeyBaseFromDir(cfg.TxCfg.rootCfg.Home)
	if err != nil {
		return fmt.Sprintf("Error: unable to open Keybase: %s", err.Error())
	}
	_, err = kb.CreateAccount(cfg.KeyName,
		"enable until hover project know foam join table educate room better scrub clever powder virus army pitch ranch fix try cupboard scatter dune fee",
		"", cfg.Password, uint32(0), uint32(0))
	if err != nil {
		return fmt.Sprintf("Error: unable to create account: %s", err.Error())
	}

	message := "Hello from GnoMobile demo"
	err = callCreateReply(cfg, "2", "1", "1", message)
	if err != nil {
		return fmt.Sprintf("Error: unable to exec call command: %s", err.Error())
	}

	return fmt.Sprintf("Posted: %s", message)
}

func ExportJsonConfig() (string, error) {
	config, err := json.Marshal(getAccountAndTxCfg())
	if err != nil {
		return "", err
	}

	return string(config), nil
}

func getAccountAndTxCfg() *accountAndTxCfg {
	dataDir := "data"
	remote := "testnet.gno.berty.io:26657"
	chainID := "dev"
	keyName := "jefft0"
	password := "password"

	return &accountAndTxCfg{
		TxCfg: &makeTxCfg{
			rootCfg: &baseCfg{
				BaseOptions: BaseOptions{
					Home:   dataDir,
					Remote: remote,
				},
			},
			gasWanted: 2000000,
			gasFee:    "1000000ugnot",

			broadcast: true,
			chainID:   chainID,
		},
		KeyName:  keyName,
		Password: password,
	}
}

func callCreateThread(cfg *accountAndTxCfg, boardId string, title string, body string) error {
	callCfg := &callCfg{
		rootCfg:  cfg.TxCfg,
		pkgPath:  "gno.land/r/demo/boards",
		funcName: "CreateThread",
		args:     []string{boardId, title, body},
	}
	return execCall(callCfg, cfg.KeyName, cfg.Password)
}

func callCreateReply(cfg *accountAndTxCfg, boardId string, threadId string, postId string, body string) error {
	callCfg := &callCfg{
		rootCfg:  cfg.TxCfg,
		pkgPath:  "gno.land/r/demo/boards",
		funcName: "CreateReply",
		args:     []string{boardId, threadId, postId, body},
	}
	return execCall(callCfg, cfg.KeyName, cfg.Password)
}
