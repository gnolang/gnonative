package service

import (
	"context"

	"github.com/gnolang/gno/tm2/pkg/crypto/bip39"
	crypto_keys "github.com/gnolang/gno/tm2/pkg/crypto/keys"
	"go.uber.org/zap"

	rpcclient "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	"github.com/gnolang/gnomobile/gnoclient"
	"github.com/gnolang/gnomobile/service/rpc"
)

// Set the connection addresse for the remote node. If you don't call this, the default is
// "127.0.0.1:26657"
func (s *gnomobileService) SetRemote(ctx context.Context, req *rpc.SetRemote_Request) (*rpc.SetRemote_Reply, error) {
	s.client.RPCClient = rpcclient.NewHTTP(req.Remote, "/websocket")
	return &rpc.SetRemote_Reply{}, nil
}

// Set the chain ID for the remote node. If you don't call this, the default is "dev"
func (s *gnomobileService) SetChainID(ctx context.Context, req *rpc.SetChainID_Request) (*rpc.SetChainID_Reply, error) {
	s.getSigner().ChainID = req.ChainID
	return &rpc.SetChainID_Reply{}, nil
}

// Set the nameOrBech32 for the account in the keybase, used for later operations
func (s *gnomobileService) SetNameOrBech32(ctx context.Context, req *rpc.SetNameOrBech32_Request) (*rpc.SetNameOrBech32_Reply, error) {
	s.getSigner().Account = req.NameOrBech32
	return &rpc.SetNameOrBech32_Reply{}, nil
}

// Set the password for the account in the keybase, used for later operations
func (s *gnomobileService) SetPassword(ctx context.Context, req *rpc.SetPassword_Request) (*rpc.SetPassword_Reply, error) {
	s.getSigner().Password = req.Password
	return &rpc.SetPassword_Reply{}, nil
}

// Generate a recovery phrase of BIP39 mnemonic words using entropy from the crypto library
// random number generator. This can be used as the mnemonic in CreateAccount.
func (s *gnomobileService) GenerateRecoveryPhrase(ctx context.Context, req *rpc.GenerateRecoveryPhrase_Request) (*rpc.GenerateRecoveryPhrase_Reply, error) {
	const mnemonicEntropySize = 256
	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return nil, err
	}

	phrase, err := bip39.NewMnemonic(entropySeed[:])
	if err != nil {
		return nil, err
	}

	return &rpc.GenerateRecoveryPhrase_Reply{Phrase: phrase}, nil
}

func convertKeyInfo(key crypto_keys.Info) (*rpc.KeyInfo, error) {
	var keyType rpc.KeyType

	switch key.GetType() {
	case crypto_keys.TypeLocal:
		keyType = rpc.KeyType_TypeLocal
	case crypto_keys.TypeLedger:
		keyType = rpc.KeyType_TypeLedger
	case crypto_keys.TypeOffline:
		keyType = rpc.KeyType_TypeOffline
	case crypto_keys.TypeMulti:
		keyType = rpc.KeyType_TypeMulti
	default:
		return nil, rpc.ErrCode_ErrCryptoKeyTypeUnknown
	}

	return &rpc.KeyInfo{
		Type:    keyType,
		Name:    key.GetName(),
		Address: key.GetAddress().Bytes(),
		PubKey:  key.GetPubKey().Bytes(),
	}, nil
}

// Get the keys informations in the keybase
func (s *gnomobileService) ListKeyInfo(ctx context.Context, req *rpc.ListKeyInfo_Request) (*rpc.ListKeyInfo_Reply, error) {
	s.logger.Debug("ListKeyInfo called")

	keys, err := s.getSigner().Keybase.List()
	if err != nil {
		return nil, err
	}

	formatedKeys := make([]*rpc.KeyInfo, len(keys))

	for _, key := range keys {
		info, err := convertKeyInfo(key)
		if err != nil {
			return nil, err
		}

		formatedKeys = append(formatedKeys, info)
	}

	return &rpc.ListKeyInfo_Reply{Keys: formatedKeys}, nil
}

// Create a new account in the keybase
func (s *gnomobileService) CreateAccount(ctx context.Context, req *rpc.CreateAccount_Request) (*rpc.CreateAccount_Reply, error) {
	s.logger.Debug("CreateAccount called", zap.String("NameOrBech32", req.NameOrBech32))

	key, err := s.getSigner().Keybase.CreateAccount(req.NameOrBech32, req.Mnemonic, req.Bip39Passwd, req.Password, req.Account, req.Index)
	if err != nil {
		return nil, err
	}

	info, err := convertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	return &rpc.CreateAccount_Reply{Key: info}, nil
}

// SelectAccount selects the account to use for later operations
func (s *gnomobileService) SelectAccount(ctx context.Context, req *rpc.SelectAccount_Request) (*rpc.SelectAccount_Reply, error) {
	s.logger.Debug("SelectAccount called", zap.String("NameOrBech32", req.NameOrBech32))

	key, err := s.getSigner().Keybase.GetByNameOrAddress(req.NameOrBech32)
	if err != nil {
		return nil, rpc.ErrCode_ErrCryptoKeyNotFound
	}

	s.lock.Lock()
	s.activeAccount = key
	s.lock.Unlock()

	info, err := convertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	s.getSigner().Account = req.NameOrBech32
	return &rpc.SelectAccount_Reply{Key: info}, nil
}

// Make an ABCI query to the remote node.
func (s *gnomobileService) Query(ctx context.Context, req *rpc.Query_Request) (*rpc.Query_Reply, error) {
	return &rpc.Query_Reply{}, nil
}

// Call a specific realm function.
func (s *gnomobileService) Call(ctx context.Context, req *rpc.Call_Request) (*rpc.Call_Reply, error) {
	s.logger.Debug("Call", zap.String("package", req.PackagePath), zap.String("function", req.Fnc), zap.Any("args", req.Args))

	if s.activeAccount == nil {
		return nil, rpc.ErrCode_ErrNoActiveAccount
	}

	cfg := gnoclient.CallCfg{
		PkgPath:   req.PackagePath,
		FuncName:  req.Fnc,
		Args:      req.Args,
		GasFee:    req.GasFee,
		GasWanted: req.GasWanted,
		Send:      req.Send,
		Memo:      req.Memo,
	}

	bres, err := s.client.Call(cfg)
	if err != nil {
		return nil, err
	}

	return &rpc.Call_Reply{Result: bres.DeliverTx.Data}, nil
}
