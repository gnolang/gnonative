package service

import (
	"context"

	crypto_keys "github.com/gnolang/gno/tm2/pkg/crypto/keys"
	"go.uber.org/zap"

	"github.com/gnolang/gnomobile/service/gnomobiletypes"
)

// Set the connection addresse for the remote node. If you don't call this, the default is
// "127.0.0.1:26657"
func (s *gnomobileService) SetRemote(ctx context.Context, req *gnomobiletypes.SetRemote_Request) (*gnomobiletypes.SetRemote_Reply, error) {
	s.client.SetRemote(req.Remote)
	return &gnomobiletypes.SetRemote_Reply{}, nil
}

// Set the chain ID for the remote node. If you don't call this, the default is "dev"
func (s *gnomobileService) SetChainID(ctx context.Context, req *gnomobiletypes.SetChainID_Request) (*gnomobiletypes.SetChainID_Reply, error) {
	s.client.SetChainID(req.ChainID)
	return &gnomobiletypes.SetChainID_Reply{}, nil
}

// Set the nameOrBech32 for the account in the keybase, used for later operations
func (s *gnomobileService) SetNameOrBench32(ctx context.Context, req *gnomobiletypes.SetNameOrBech32_Request) (*gnomobiletypes.SetNameOrBech32_Reply, error) {
	s.client.SetNameOrBech32(req.NameOrBech32)
	return &gnomobiletypes.SetNameOrBech32_Reply{}, nil
}

// Set the password for the account in the keybase, used for later operations
func (s *gnomobileService) SetPassword(ctx context.Context, req *gnomobiletypes.SetPassword_Request) (*gnomobiletypes.SetPassword_Reply, error) {
	s.client.SetPassword(req.Password)
	return &gnomobiletypes.SetPassword_Reply{}, nil
}

func convertKeyInfo(key crypto_keys.Info) (*gnomobiletypes.KeyInfo, error) {
	var keyType gnomobiletypes.KeyType

	switch key.GetType() {
	case crypto_keys.TypeLocal:
		keyType = gnomobiletypes.KeyType_TypeLocal
	case crypto_keys.TypeLedger:
		keyType = gnomobiletypes.KeyType_TypeLedger
	case crypto_keys.TypeOffline:
		keyType = gnomobiletypes.KeyType_TypeOffline
	case crypto_keys.TypeMulti:
		keyType = gnomobiletypes.KeyType_TypeMulti
	default:
		return nil, gnomobiletypes.ErrCode_ErrCryptoKeyTypeUnknown
	}

	return &gnomobiletypes.KeyInfo{
		Type:    keyType,
		Name:    key.GetName(),
		Address: key.GetAddress().Bytes(),
		PubKey:  key.GetPubKey().Bytes(),
	}, nil
}

// Get the keys informations in the keybase
func (s *gnomobileService) ListKeyInfo(ctx context.Context, req *gnomobiletypes.ListKeyInfo_Request) (*gnomobiletypes.ListKeyInfo_Reply, error) {
	keys, err := s.client.GetKeys()
	if err != nil {
		return nil, err
	}

	formatedKeys := make([]*gnomobiletypes.KeyInfo, len(keys))

	for _, key := range keys {
		info, err := convertKeyInfo(key)
		if err != nil {
			return nil, err
		}

		formatedKeys = append(formatedKeys, info)
	}

	return &gnomobiletypes.ListKeyInfo_Reply{Keys: formatedKeys}, nil
}

// Create a new account in the keybase
func (s *gnomobileService) CreateAccount(ctx context.Context, req *gnomobiletypes.CreateAccount_Request) (*gnomobiletypes.CreateAccount_Reply, error) {
	key, err := s.client.CreateAccount(req.NameOrBech32, req.Mnemonic, req.Bip39Passwd, req.Password, req.Account, req.Index)
	if err != nil {
		return nil, err
	}

	info, err := convertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	return &gnomobiletypes.CreateAccount_Reply{Key: info}, nil
}

// SelectAccount selects the account to use for later operations
func (s *gnomobileService) SelectAccount(ctx context.Context, req *gnomobiletypes.SelectAccount_Request) (*gnomobiletypes.SelectAccount_Reply, error) {
	key, err := s.client.GetKeyByNameOrBech32(req.NameOrBech32)
	if err != nil {
		return nil, gnomobiletypes.ErrCode_ErrCryptoKeyNotFound
	}

	s.lock.Lock()
	s.activeAccount = key
	s.lock.Unlock()

	info, err := convertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	return &gnomobiletypes.SelectAccount_Reply{Key: info}, nil
}

// Make an ABCI query to the remote node.
func (s *gnomobileService) Query(ctx context.Context, req *gnomobiletypes.Query_Request) (*gnomobiletypes.Query_Reply, error) {
	return &gnomobiletypes.Query_Reply{}, nil
}

// Call a specific realm function.
func (s *gnomobileService) Call(ctx context.Context, req *gnomobiletypes.Call_Request) (*gnomobiletypes.Call_Reply, error) {
	s.logger.Debug("Call", zap.String("package", req.PackagePath), zap.String("function", req.Fnc), zap.Any("args", req.Args))

	if err := s.client.Call(req.PackagePath, req.Fnc, req.Args, req.GasFee, req.GasWanted, "", s.activeAccount.GetName(), req.Password); err != nil {
		return nil, err
	}

	return &gnomobiletypes.Call_Reply{}, nil
}
