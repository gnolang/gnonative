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
	Mnemonic string
}

func CreateDefaulAccount(rootDir string) error {
	cfg := getAccountAndTxCfg(rootDir)

	kb, err := keys.NewKeyBaseFromDir(cfg.TxCfg.rootCfg.Home)
	if err != nil {
		return err
	}
	keyList, err := kb.List()
	if err != nil {
		return err
	}
	if len(keyList) > 0 {
		// Assume the account with cfg.KeyName is already created.
		return nil
	}

	_, err = kb.CreateAccount(cfg.KeyName, cfg.Mnemonic, "", cfg.Password, uint32(0), uint32(0))
	if err != nil {
		return err
	}
	return nil
}

func getAccountAndTxCfg(rootDir string) *accountAndTxCfg {
	dataDir := rootDir + "/data"
	remote := "testnet.gno.berty.io:26657"
	chainID := "dev"
	keyName := "jefft0"
	password := "password"
	mnemonic := "enable until hover project know foam join table educate room better scrub clever powder virus army pitch ranch fix try cupboard scatter dune fee"

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
		Mnemonic: mnemonic,
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

func ExportJsonConfig(rootDir string) string {
	config, err := json.Marshal(getAccountAndTxCfg(rootDir))
	if err != nil {
		return fmt.Sprintf("Error: unable load config: %s", err.Error())
	}
	return string(config)
}

func CreateReply(message string, rootDir string) string {
	cfg := getAccountAndTxCfg(rootDir)

	err := callCreateReply(cfg, "2", "1", "1", message)

	if err != nil {
		return fmt.Sprintf("Error: unable to exec call command: %s", err.Error())
	}

	return fmt.Sprintf("Posted: %s", message)
}
