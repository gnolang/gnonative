// This file implements the gRPC API methods defined in service/rpc/rpc.proto . For documentation,
// see that file and related request/response fields in the generated gnomobiletypes.proto .

package service

import (
	"context"

	"connectrpc.com/connect"
	"github.com/gnolang/gno/tm2/pkg/crypto"
	"github.com/gnolang/gno/tm2/pkg/crypto/bip39"
	crypto_keys "github.com/gnolang/gno/tm2/pkg/crypto/keys"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys/keyerror"
	"go.uber.org/zap"

	rpcclient "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	"github.com/gnolang/gnomobile/gnoclient"
	"github.com/gnolang/gnomobile/service/rpc"
)

func (s *gnomobileService) SetRemote(ctx context.Context, req *connect.Request[rpc.SetRemoteRequest]) (*connect.Response[rpc.SetRemoteResponse], error) {
	s.client.RPCClient = rpcclient.NewHTTP(req.Msg.Remote, "/websocket")
	return connect.NewResponse(&rpc.SetRemoteResponse{}), nil
}

func (s *gnomobileService) SetChainID(ctx context.Context, req *connect.Request[rpc.SetChainIDRequest]) (*connect.Response[rpc.SetChainIDResponse], error) {
	s.getSigner().ChainID = req.Msg.ChainId
	return connect.NewResponse(&rpc.SetChainIDResponse{}), nil
}

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

func (s *gnomobileService) SelectAccount(ctx context.Context, req *connect.Request[rpc.SelectAccountRequest]) (*connect.Response[rpc.SelectAccountResponse], error) {
	s.logger.Debug("SelectAccount called", zap.String("NameOrBech32", req.Msg.NameOrBech32))

	key, err := s.getSigner().Keybase.GetByNameOrAddress(req.Msg.NameOrBech32)
	if err != nil {
		return nil, rpc.ErrCode_ErrCryptoKeyNotFound
	}

	info, err := convertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	s.lock.Lock()
	s.activeAccount = key
	s.lock.Unlock()

	s.getSigner().Account = req.Msg.NameOrBech32
	return connect.NewResponse(&rpc.SelectAccountResponse{Key: info}), nil
}

func (s *gnomobileService) SetPassword(ctx context.Context, req *connect.Request[rpc.SetPasswordRequest]) (*connect.Response[rpc.SetPasswordResponse], error) {
	s.getSigner().Password = req.Msg.Password
	return connect.NewResponse(&rpc.SetPasswordResponse{}), nil
}

func (s *gnomobileService) GetActiveAccount(ctx context.Context, req *connect.Request[rpc.GetActiveAccountRequest]) (*connect.Response[rpc.GetActiveAccountResponse], error) {
	s.logger.Debug("GetActiveAccount called")

	s.lock.RLock()
	key := s.activeAccount
	s.lock.RUnlock()

	if key == nil {
		return nil, rpc.ErrCode_ErrNoActiveAccount
	}

	info, err := convertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&rpc.GetActiveAccountResponse{Key: info}), nil
}

func (s *gnomobileService) QueryAccount(ctx context.Context, req *connect.Request[rpc.QueryAccountRequest]) (*connect.Response[rpc.QueryAccountResponse], error) {
	s.logger.Debug("QueryAccount", zap.ByteString("address", req.Msg.Address))

	// gnoclient wants the bech32 address.
	account, _, err := s.client.QueryAccount(crypto.AddressToBech32(crypto.AddressFromBytes(req.Msg.Address)))
	if err != nil {
		return nil, err
	}

	formattedCoins := make([]*rpc.Coin, 0)
	for _, coin := range account.Coins {
		formattedCoins = append(formattedCoins, &rpc.Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount,
		})
	}

	res := connect.NewResponse(&rpc.QueryAccountResponse{AccountInfo: &rpc.BaseAccount{
		Address:       account.Address.Bytes(),
		Coins:         formattedCoins,
		PubKey:        account.PubKey.Bytes(),
		AccountNumber: account.AccountNumber,
		Sequence:      account.Sequence,
	}})
	return res, nil
}

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

	s.lock.Lock()
	if s.activeAccount != nil &&
		(s.activeAccount.GetName() == req.Msg.NameOrBech32 || crypto.AddressToBech32(s.activeAccount.GetAddress()) == req.Msg.NameOrBech32) {
		// The deleted account was the active account.
		s.activeAccount = nil
	}
	s.lock.Unlock()
	return connect.NewResponse(&rpc.DeleteAccountResponse{}), nil
}

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

func (s *gnomobileService) Call(ctx context.Context, req *connect.Request[rpc.CallRequest]) (*connect.Response[rpc.CallResponse], error) {
	s.logger.Debug("Call", zap.String("package", req.Msg.PackagePath), zap.String("function", req.Msg.Fnc), zap.Any("args", req.Msg.Args))

	s.lock.RLock()
	if s.activeAccount == nil {
		s.lock.RUnlock()
		return nil, rpc.ErrCode_ErrNoActiveAccount
	}
	s.lock.RUnlock()

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

func (s *gnomobileService) AddressToBech32(ctx context.Context, req *connect.Request[rpc.AddressToBech32Request]) (*connect.Response[rpc.AddressToBech32Response], error) {
	s.logger.Debug("AddressToBech32", zap.ByteString("address", req.Msg.Address))
	bech32Address := crypto.AddressToBech32(crypto.AddressFromBytes(req.Msg.Address))
	return connect.NewResponse(&rpc.AddressToBech32Response{Bech32Address: bech32Address}), nil
}

func (s *gnomobileService) AddressFromBech32(ctx context.Context, req *connect.Request[rpc.AddressFromBech32Request]) (*connect.Response[rpc.AddressFromBech32Response], error) {
	address, err := crypto.AddressFromBech32(req.Msg.Bech32Address)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&rpc.AddressFromBech32Response{Address: address.Bytes()}), nil
}

func (s *gnomobileService) Hello(ctx context.Context, req *connect.Request[rpc.HelloRequest]) (*connect.Response[rpc.HelloResponse], error) {
	return connect.NewResponse(&rpc.HelloResponse{
		Greeting: "Hello " + req.Msg.Name,
	}), nil
}
