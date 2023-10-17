package service

import (
	"context"

	"connectrpc.com/connect"
	"github.com/gnolang/gno/tm2/pkg/crypto/bip39"
	crypto_keys "github.com/gnolang/gno/tm2/pkg/crypto/keys"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys/keyerror"
	"go.uber.org/zap"

	rpcclient "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	"github.com/gnolang/gnomobile/gnoclient"
	"github.com/gnolang/gnomobile/service/rpc"
)

// Set the connection addresse for the remote node. If you don't call this, the default is
// "127.0.0.1:26657"
func (s *gnomobileService) SetRemote(ctx context.Context, req *connect.Request[rpc.SetRemoteRequest]) (*connect.Response[rpc.SetRemoteResponse], error) {
	s.client.RPCClient = rpcclient.NewHTTP(req.Msg.Remote, "/websocket")
	return connect.NewResponse(&rpc.SetRemoteResponse{}), nil
}

// Set the chain ID for the remote node. If you don't call this, the default is "dev"
func (s *gnomobileService) SetChainID(ctx context.Context, req *connect.Request[rpc.SetChainIDRequest]) (*connect.Response[rpc.SetChainIDResponse], error) {
	s.getSigner().ChainID = req.Msg.ChainId
	return connect.NewResponse(&rpc.SetChainIDResponse{}), nil
}

// Set the password for the account in the keybase, used for later operations
func (s *gnomobileService) SetPassword(ctx context.Context, req *connect.Request[rpc.SetPasswordRequest]) (*connect.Response[rpc.SetPasswordResponse], error) {
	s.getSigner().Password = req.Msg.Password
	return connect.NewResponse(&rpc.SetPasswordResponse{}), nil
}

// Generate a recovery phrase of BIP39 mnemonic words using entropy from the crypto library
// random number generator. This can be used as the mnemonic in CreateAccount.
func (s *gnomobileService) GenerateRecoveryPhrase(ctx context.Context, req *connect.Request[rpc.GenerateRecoveryPhraseRequest]) (*connect.Response[rpc.GenerateRecoveryPhraseResponse], error) {
	const mnemonicEntropySize = 256
	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return nil, err
	}

	phrase, err := bip39.NewMnemonic(entropySeed[:])
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&rpc.GenerateRecoveryPhraseResponse{Phrase: phrase}), nil
}

func convertKeyInfo(key crypto_keys.Info) (*rpc.KeyInfo, error) {
	return &rpc.KeyInfo{
		Type:    uint32(key.GetType()),
		Name:    key.GetName(),
		Address: key.GetAddress().Bytes(),
		PubKey:  key.GetPubKey().Bytes(),
	}, nil
}

// Get the keys informations in the keybase
func (s *gnomobileService) ListKeyInfo(ctx context.Context, req *connect.Request[rpc.ListKeyInfoRequest]) (*connect.Response[rpc.ListKeyInfoResponse], error) {
	s.logger.Debug("ListKeyInfo called")

	keys, err := s.getSigner().Keybase.List()
	if err != nil {
		return nil, err
	}

	formatedKeys := make([]*rpc.KeyInfo, 0)

	for _, key := range keys {
		info, err := convertKeyInfo(key)
		if err != nil {
			return nil, err
		}

		formatedKeys = append(formatedKeys, info)
	}

	return connect.NewResponse(&rpc.ListKeyInfoResponse{Keys: formatedKeys}), nil
}

// Create a new account in the keybase
func (s *gnomobileService) CreateAccount(ctx context.Context, req *connect.Request[rpc.CreateAccountRequest]) (*connect.Response[rpc.CreateAccountResponse], error) {
	s.logger.Debug("CreateAccount called", zap.String("NameOrBech32", req.Msg.NameOrBech32))

	key, err := s.getSigner().Keybase.CreateAccount(req.Msg.NameOrBech32, req.Msg.Mnemonic, req.Msg.Bip39Passwd, req.Msg.Password, req.Msg.Account, req.Msg.Index)
	if err != nil {
		return nil, err
	}

	info, err := convertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&rpc.CreateAccountResponse{Key: info}), nil
}

// SelectAccount selects the active account to use for later operations
func (s *gnomobileService) SelectAccount(ctx context.Context, req *connect.Request[rpc.SelectAccountRequest]) (*connect.Response[rpc.SelectAccountResponse], error) {
	s.logger.Debug("SelectAccount called", zap.String("NameOrBech32", req.Msg.NameOrBech32))

	key, err := s.getSigner().Keybase.GetByNameOrAddress(req.Msg.NameOrBech32)
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

	s.getSigner().Account = req.Msg.NameOrBech32
	return connect.NewResponse(&rpc.SelectAccountResponse{Key: info}), nil
}

// GetActiveAccount gets the active account which was set by SelectAccount
func (s *gnomobileService) GetActiveAccount(ctx context.Context, req *connect.Request[rpc.GetActiveAccountRequest]) (*connect.Response[rpc.GetActiveAccountResponse], error) {
	s.logger.Debug("GetActiveAccount called")

	s.lock.Lock()
	key := s.activeAccount
	s.lock.Unlock()

	if key == nil {
		return nil, rpc.ErrCode_ErrNoActiveAccount
	}

	info, err := convertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&rpc.GetActiveAccountResponse{Key: info}), nil
}

// DeleteAccount deletes the account with the given name, using the password to
// ensure access. However, if skip_password is true, then ignore the password.
// If the account doesn't exist, then return ErrCryptoKeyNotFound.
// If the password is wrong, then return ErrDecryptionFailed.
func (s *gnomobileService) DeleteAccount(ctx context.Context, req *connect.Request[rpc.DeleteAccountRequest]) (*connect.Response[rpc.DeleteAccountResponse], error) {
	if err := s.getSigner().Keybase.Delete(req.Msg.NameOrBech32, req.Msg.Password, req.Msg.SkipPassword); err != nil {
		if keyerror.IsErrKeyNotFound(err) {
			return nil, rpc.ErrCode_ErrCryptoKeyNotFound
		} else if keyerror.IsErrWrongPassword(err) {
			return nil, rpc.ErrCode_ErrDecryptionFailed
		} else {
			return nil, err
		}
	}
	return connect.NewResponse(&rpc.DeleteAccountResponse{}), nil
}

// Make an ABCI query to the remote node.
func (s *gnomobileService) Query(ctx context.Context, req *connect.Request[rpc.QueryRequest]) (*connect.Response[rpc.QueryResponse], error) {
	s.logger.Debug("Query", zap.String("path", req.Msg.Path), zap.ByteString("data", req.Msg.Data))

	cfg := gnoclient.QueryCfg{
		Path: req.Msg.Path,
		Data: req.Msg.Data,
	}

	bres, err := s.client.Query(cfg)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&rpc.QueryResponse{Result: bres.Response.Data}), nil
}

// Call a specific realm function.
func (s *gnomobileService) Call(ctx context.Context, req *connect.Request[rpc.CallRequest]) (*connect.Response[rpc.CallResponse], error) {
	s.logger.Debug("Call", zap.String("package", req.Msg.PackagePath), zap.String("function", req.Msg.Fnc), zap.Any("args", req.Msg.Args))

	if s.activeAccount == nil {
		return nil, rpc.ErrCode_ErrNoActiveAccount
	}

	cfg := gnoclient.CallCfg{
		PkgPath:   req.Msg.PackagePath,
		FuncName:  req.Msg.Fnc,
		Args:      req.Msg.Args,
		GasFee:    req.Msg.GasFee,
		GasWanted: req.Msg.GasWanted,
		Send:      req.Msg.Send,
		Memo:      req.Msg.Memo,
	}

	bres, err := s.client.Call(cfg)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&rpc.CallResponse{Result: bres.DeliverTx.Data}), nil
}

// Hello is for debug purposes
func (s *gnomobileService) Hello(ctx context.Context, req *connect.Request[rpc.HelloRequest]) (*connect.Response[rpc.HelloResponse], error) {
	return connect.NewResponse(&rpc.HelloResponse{
		Greeting: "Hello " + req.Msg.Name,
	}), nil
}
