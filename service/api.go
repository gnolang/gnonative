// This file implements the gRPC API methods defined in api/rpc.proto . For documentation,
// see that file and related request/response fields in the generated api/gnonativetypes.proto .

package service

import (
	"context"
	"errors"
	"time"

	"connectrpc.com/connect"
	"github.com/gnolang/gno/tm2/pkg/amino"
	"github.com/gnolang/gno/tm2/pkg/crypto"
	"github.com/gnolang/gno/tm2/pkg/crypto/bip39"
	crypto_keys "github.com/gnolang/gno/tm2/pkg/crypto/keys"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys/keyerror"
	"github.com/gnolang/gno/tm2/pkg/sdk/bank"
	"github.com/gnolang/gno/tm2/pkg/std"
	"go.uber.org/zap"

	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/gno.land/pkg/sdk/vm"
	rpcclient "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	api_gen "github.com/gnolang/gnonative/api/gen/go"
)

func (s *gnoNativeService) SetRemote(ctx context.Context, req *connect.Request[api_gen.SetRemoteRequest]) (*connect.Response[api_gen.SetRemoteResponse], error) {
	var err error
	s.rpcClient, err = rpcclient.NewHTTPClient(req.Msg.Remote)
	if err != nil {
		return nil, api_gen.ErrCode_ErrSetRemote.Wrap(err)
	}
	s.remote = req.Msg.Remote
	return connect.NewResponse(&api_gen.SetRemoteResponse{}), nil
}

func (s *gnoNativeService) GetRemote(ctx context.Context, req *connect.Request[api_gen.GetRemoteRequest]) (*connect.Response[api_gen.GetRemoteResponse], error) {
	if s.useGnokeyMobile {
		// Always get the remote from the Gnokey Mobile service
		res, err := s.gnokeyMobileClient.GetRemote(context.Background(), req)
		if err != nil {
			return nil, err
		}

		return connect.NewResponse(res.Msg), nil
	}

	return connect.NewResponse(&api_gen.GetRemoteResponse{Remote: s.ClientGetRemote()}), nil
}

func (s *gnoNativeService) ClientGetRemote() string {
	return s.remote
}

func (s *gnoNativeService) SetChainID(ctx context.Context, req *connect.Request[api_gen.SetChainIDRequest]) (*connect.Response[api_gen.SetChainIDResponse], error) {
	s.lock.Lock()
	s.chainID = req.Msg.ChainId
	s.lock.Unlock()
	return connect.NewResponse(&api_gen.SetChainIDResponse{}), nil
}

func (s *gnoNativeService) GetChainID(ctx context.Context, req *connect.Request[api_gen.GetChainIDRequest]) (*connect.Response[api_gen.GetChainIDResponse], error) {
	return connect.NewResponse(&api_gen.GetChainIDResponse{ChainId: s.chainID}), nil
}

func (s *gnoNativeService) GenerateRecoveryPhrase(ctx context.Context, req *connect.Request[api_gen.GenerateRecoveryPhraseRequest]) (*connect.Response[api_gen.GenerateRecoveryPhraseResponse], error) {
	const mnemonicEntropySize = 256
	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return nil, err
	}

	phrase, err := bip39.NewMnemonic(entropySeed[:])
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&api_gen.GenerateRecoveryPhraseResponse{Phrase: phrase}), nil
}

func ConvertKeyInfo(key crypto_keys.Info) (*api_gen.KeyInfo, error) {
	return &api_gen.KeyInfo{
		Type:    uint32(key.GetType()),
		Name:    key.GetName(),
		Address: key.GetAddress().Bytes(),
		PubKey:  key.GetPubKey().Bytes(),
	}, nil
}

func (s *gnoNativeService) ListKeyInfo(ctx context.Context, req *connect.Request[api_gen.ListKeyInfoRequest]) (*connect.Response[api_gen.ListKeyInfoResponse], error) {
	s.logger.Debug("ListKeyInfo called")

	if s.useGnokeyMobile {
		// Always get the list of keys from the Gnokey Mobile service
		res, err := s.gnokeyMobileClient.ListKeyInfo(context.Background(), req)
		if err != nil {
			return nil, err
		}

		return connect.NewResponse(res.Msg), nil
	}

	keys, err := s.ClientListKeyInfo()
	if err != nil {
		return nil, err
	}

	formatedKeys := make([]*api_gen.KeyInfo, 0)

	for _, key := range keys {
		info, err := ConvertKeyInfo(key)
		if err != nil {
			return nil, err
		}

		formatedKeys = append(formatedKeys, info)
	}

	return connect.NewResponse(&api_gen.ListKeyInfoResponse{Keys: formatedKeys}), nil
}

func (s *gnoNativeService) ClientListKeyInfo() ([]crypto_keys.Info, error) {
	return s.keybase.List()
}

func (s *gnoNativeService) HasKeyByName(ctx context.Context, req *connect.Request[api_gen.HasKeyByNameRequest]) (*connect.Response[api_gen.HasKeyByNameResponse], error) {
	s.logger.Debug("HasKeyByName called")

	has, err := s.keybase.HasByName(req.Msg.Name)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return connect.NewResponse(&api_gen.HasKeyByNameResponse{Has: has}), nil
}

func (s *gnoNativeService) HasKeyByAddress(ctx context.Context, req *connect.Request[api_gen.HasKeyByAddressRequest]) (*connect.Response[api_gen.HasKeyByAddressResponse], error) {
	s.logger.Debug("HasKeyByAddress called")

	has, err := s.keybase.HasByAddress(crypto.AddressFromBytes(req.Msg.Address))
	if err != nil {
		return nil, getGrpcError(err)
	}

	return connect.NewResponse(&api_gen.HasKeyByAddressResponse{Has: has}), nil
}

func (s *gnoNativeService) HasKeyByNameOrAddress(ctx context.Context, req *connect.Request[api_gen.HasKeyByNameOrAddressRequest]) (*connect.Response[api_gen.HasKeyByNameOrAddressResponse], error) {
	s.logger.Debug("HasKeyByNameOrAddress called")

	has, err := s.keybase.HasByNameOrAddress(req.Msg.NameOrBech32)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return connect.NewResponse(&api_gen.HasKeyByNameOrAddressResponse{Has: has}), nil
}

func (s *gnoNativeService) GetKeyInfoByName(ctx context.Context, req *connect.Request[api_gen.GetKeyInfoByNameRequest]) (*connect.Response[api_gen.GetKeyInfoByNameResponse], error) {
	s.logger.Debug("GetKeyInfoByName called")

	key, err := s.keybase.GetByName(req.Msg.Name)
	if err != nil {
		return nil, getGrpcError(err)
	}

	info, err := ConvertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&api_gen.GetKeyInfoByNameResponse{Key: info}), nil
}

func (s *gnoNativeService) GetKeyInfoByAddress(ctx context.Context, req *connect.Request[api_gen.GetKeyInfoByAddressRequest]) (*connect.Response[api_gen.GetKeyInfoByAddressResponse], error) {
	s.logger.Debug("GetKeyInfoByAddress called")

	key, err := s.keybase.GetByAddress(crypto.AddressFromBytes(req.Msg.Address))
	if err != nil {
		return nil, getGrpcError(err)
	}

	info, err := ConvertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&api_gen.GetKeyInfoByAddressResponse{Key: info}), nil
}

func (s *gnoNativeService) GetKeyInfoByNameOrAddress(ctx context.Context, req *connect.Request[api_gen.GetKeyInfoByNameOrAddressRequest]) (*connect.Response[api_gen.GetKeyInfoByNameOrAddressResponse], error) {
	s.logger.Debug("GetKeyInfoByNameOrAddress called")

	key, err := s.keybase.GetByNameOrAddress(req.Msg.NameOrBech32)
	if err != nil {
		return nil, getGrpcError(err)
	}

	info, err := ConvertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&api_gen.GetKeyInfoByNameOrAddressResponse{Key: info}), nil
}

func (s *gnoNativeService) CreateAccount(ctx context.Context, req *connect.Request[api_gen.CreateAccountRequest]) (*connect.Response[api_gen.CreateAccountResponse], error) {
	s.logger.Debug("CreateAccount called", zap.String("NameOrBech32", req.Msg.NameOrBech32))

	key, err := s.keybase.CreateAccount(req.Msg.NameOrBech32, req.Msg.Mnemonic, req.Msg.Bip39Passwd, req.Msg.Password, req.Msg.Account, req.Msg.Index)
	if err != nil {
		return nil, getGrpcError(err)
	}

	info, err := ConvertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&api_gen.CreateAccountResponse{Key: info}), nil
}

func (s *gnoNativeService) SelectAccount(ctx context.Context, req *connect.Request[api_gen.SelectAccountRequest]) (*connect.Response[api_gen.SelectAccountResponse], error) {
	s.logger.Debug("DEPRECATED: SelectAccount called", zap.String("NameOrBech32", req.Msg.NameOrBech32))

	// The key may already be in s.userAccounts, but the info may have changed on disk. So always get from disk.
	key, err := s.keybase.GetByNameOrAddress(req.Msg.NameOrBech32)
	if err != nil {
		return nil, getGrpcError(err)
	}

	info, err := ConvertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	bech32 := crypto.AddressToBech32(key.GetAddress())
	s.lock.Lock()
	account, ok := s.userAccounts[bech32]
	if !ok {
		account = &userAccount{}
		account.signer = &gnoclient.SignerFromKeybase{
			Keybase: s.keybase,
			ChainID: s.chainID,
		}
		s.userAccounts[bech32] = account
	}
	account.keyInfo = key
	s.activeAccount = account
	s.lock.Unlock()

	account.signer.Account = req.Msg.NameOrBech32
	return connect.NewResponse(&api_gen.SelectAccountResponse{
		Key:         info,
		HasPassword: account.signer.Password != "",
	}), nil
}

func (s *gnoNativeService) ActivateAccount(ctx context.Context, req *connect.Request[api_gen.ActivateAccountRequest]) (*connect.Response[api_gen.ActivateAccountResponse], error) {
	s.logger.Debug("ActivateAccount called", zap.String("NameOrBech32", req.Msg.NameOrBech32))

	// The key may already be in s.userAccounts, but the info may have changed on disk. So always get from disk.
	key, err := s.keybase.GetByNameOrAddress(req.Msg.NameOrBech32)
	if err != nil {
		return nil, getGrpcError(err)
	}

	info, err := ConvertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	bech32 := crypto.AddressToBech32(key.GetAddress())
	s.lock.Lock()
	account, ok := s.userAccounts[bech32]
	if !ok {
		account = &userAccount{}
		account.signer = &gnoclient.SignerFromKeybase{
			Keybase: s.keybase,
			ChainID: s.chainID,
		}
		s.userAccounts[bech32] = account
	}
	account.keyInfo = key
	s.lock.Unlock()

	account.signer.Account = req.Msg.NameOrBech32
	return connect.NewResponse(&api_gen.ActivateAccountResponse{
		Key:         info,
		HasPassword: account.signer.Password != "",
	}), nil
}

func (s *gnoNativeService) SetPassword(ctx context.Context, req *connect.Request[api_gen.SetPasswordRequest]) (*connect.Response[api_gen.SetPasswordResponse], error) {
	signer, err := s.getSigner(req.Msg.Address)
	if err != nil {
		return nil, err
	}
	signer.Password = req.Msg.Password

	// Check the password.
	if err := signer.Validate(); err != nil {
		if keyerror.IsErrWrongPassword(err) {
			// Wrong password, so unset the password
			signer.Password = ""
		}
		return nil, getGrpcError(err)
	}

	return connect.NewResponse(&api_gen.SetPasswordResponse{}), nil
}

func (s *gnoNativeService) UpdatePassword(ctx context.Context, req *connect.Request[api_gen.UpdatePasswordRequest]) (*connect.Response[api_gen.UpdatePasswordResponse], error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	// Get all the signers, before trying to update the password.
	var signers = make([]*gnoclient.SignerFromKeybase, len(req.Msg.Addresses))
	for i := range len(req.Msg.Addresses) {
		var err error
		if signers[i], err = s.getSigner(req.Msg.Addresses[i]); err != nil {
			return nil, err
		}
	}

	getNewPassword := func() (string, error) { return req.Msg.NewPassword, nil }
	for i := range len(req.Msg.Addresses) {
		if err := s.keybase.Update(signers[i].Account, signers[i].Password, getNewPassword); err != nil {
			// Roll back the passwords. Don't check the error from Update.
			for j := range i {
				getOldPassword := func() (string, error) { return signers[j].Password, nil }
				s.keybase.Update(signers[j].Account, req.Msg.NewPassword, getOldPassword)
			}
			return nil, getGrpcError(err)
		}
	}

	// Success. Update the Password in all the signers.
	for i := range len(req.Msg.Addresses) {
		signers[i].Password = req.Msg.NewPassword
	}

	return connect.NewResponse(&api_gen.UpdatePasswordResponse{}), nil
}

func (s *gnoNativeService) GetActiveAccount(ctx context.Context, req *connect.Request[api_gen.GetActiveAccountRequest]) (*connect.Response[api_gen.GetActiveAccountResponse], error) {
	s.logger.Debug("DEPRECATED: GetActiveAccount called")

	s.lock.RLock()
	account := s.activeAccount
	s.lock.RUnlock()

	if account == nil {
		return nil, api_gen.ErrCode_ErrNoActiveAccount
	}

	info, err := ConvertKeyInfo(account.keyInfo)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&api_gen.GetActiveAccountResponse{
		Key:         info,
		HasPassword: account.signer.Password != "",
	}), nil
}

func (s *gnoNativeService) GetActivatedAccount(ctx context.Context, req *connect.Request[api_gen.GetActivatedAccountRequest]) (*connect.Response[api_gen.GetActivatedAccountResponse], error) {
	s.logger.Debug("GetActivatedAccount called")

	if req.Msg.Address == nil {
		return nil, api_gen.ErrCode_ErrInvalidAddress
	}

	s.lock.Lock()
	account, ok := s.userAccounts[crypto.AddressToBech32(crypto.AddressFromBytes(req.Msg.Address))]
	s.lock.Unlock()
	if !ok {
		return nil, api_gen.ErrCode_ErrNoActiveAccount
	}

	info, err := ConvertKeyInfo(account.keyInfo)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&api_gen.GetActivatedAccountResponse{
		Key:         info,
		HasPassword: account.signer.Password != "",
	}), nil
}

func (s *gnoNativeService) QueryAccount(ctx context.Context, req *connect.Request[api_gen.QueryAccountRequest]) (*connect.Response[api_gen.QueryAccountResponse], error) {
	s.logger.Debug("QueryAccount", zap.ByteString("address", req.Msg.Address))

	c, err := s.getClient(nil)
	if err != nil {
		return nil, getGrpcError(err)
	}
	// gnoclient wants the bech32 address.
	account, _, err := c.QueryAccount(crypto.AddressFromBytes(req.Msg.Address))
	if err != nil {
		return nil, getGrpcError(err)
	}

	formattedCoins := make([]*api_gen.Coin, 0)
	for _, coin := range account.Coins {
		formattedCoins = append(formattedCoins, &api_gen.Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount,
		})
	}

	var pubKeyBytes []byte
	if account.PubKey != nil {
		pubKeyBytes = account.PubKey.Bytes()
	}
	res := connect.NewResponse(&api_gen.QueryAccountResponse{AccountInfo: &api_gen.BaseAccount{
		Address:       account.Address.Bytes(),
		Coins:         formattedCoins,
		PubKey:        pubKeyBytes,
		AccountNumber: account.AccountNumber,
		Sequence:      account.Sequence,
	}})
	return res, nil
}

func (s *gnoNativeService) DeleteAccount(ctx context.Context, req *connect.Request[api_gen.DeleteAccountRequest]) (*connect.Response[api_gen.DeleteAccountResponse], error) {
	// Get the key from the Keybase so that we know its address
	key, err := s.keybase.GetByNameOrAddress(req.Msg.NameOrBech32)
	if err != nil {
		return nil, getGrpcError(err)
	}
	if err := s.keybase.Delete(req.Msg.NameOrBech32, req.Msg.Password, req.Msg.SkipPassword); err != nil {
		return nil, getGrpcError(err)
	}

	bech32 := crypto.AddressToBech32(key.GetAddress())
	s.lock.Lock()
	delete(s.userAccounts, bech32)
	if s.activeAccount != nil && crypto.AddressToBech32(s.activeAccount.keyInfo.GetAddress()) == bech32 {
		// The deleted account was the active account.
		s.activeAccount = nil
	}
	s.lock.Unlock()
	return connect.NewResponse(&api_gen.DeleteAccountResponse{}), nil
}

func (s *gnoNativeService) Query(ctx context.Context, req *connect.Request[api_gen.QueryRequest]) (*connect.Response[api_gen.QueryResponse], error) {
	s.logger.Debug("Query", zap.String("path", req.Msg.Path), zap.ByteString("data", req.Msg.Data))

	cfg := gnoclient.QueryCfg{
		Path: req.Msg.Path,
		Data: req.Msg.Data,
	}

	c, err := s.getClient(nil)
	if err != nil {
		return nil, getGrpcError(err)
	}
	bres, err := c.Query(cfg)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return connect.NewResponse(&api_gen.QueryResponse{Result: bres.Response.Data}), nil
}

func (s *gnoNativeService) Render(ctx context.Context, req *connect.Request[api_gen.RenderRequest]) (*connect.Response[api_gen.RenderResponse], error) {
	s.logger.Debug("Render", zap.String("packagePath", req.Msg.PackagePath), zap.String("args", req.Msg.Args))

	c, err := s.getClient(nil)
	if err != nil {
		return nil, getGrpcError(err)
	}
	result, _, err := c.Render(req.Msg.PackagePath, req.Msg.Args)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return connect.NewResponse(&api_gen.RenderResponse{Result: result}), nil
}

func (s *gnoNativeService) QEval(ctx context.Context, req *connect.Request[api_gen.QEvalRequest]) (*connect.Response[api_gen.QEvalResponse], error) {
	s.logger.Debug("QEval", zap.String("packagePath", req.Msg.PackagePath), zap.String("expression", req.Msg.Expression))

	c, err := s.getClient(nil)
	if err != nil {
		return nil, getGrpcError(err)
	}
	result, _, err := c.QEval(req.Msg.PackagePath, req.Msg.Expression)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return connect.NewResponse(&api_gen.QEvalResponse{Result: result}), nil
}

func (s *gnoNativeService) Call(ctx context.Context, req *connect.Request[api_gen.CallRequest], stream *connect.ServerStream[api_gen.CallResponse]) error {
	for _, msg := range req.Msg.Msgs {
		s.logger.Debug("Call", zap.String("package", msg.PackagePath), zap.String("function", msg.Fnc), zap.Any("args", msg.Args))
	}

	cfg, msgs, err := s.convertCallRequest(req.Msg)
	if err != nil {
		return err
	}

	if s.useGnokeyMobile {
		tx, err := gnoclient.NewCallTx(*cfg, msgs...)
		if err != nil {
			return err
		}
		txJSON, err := amino.MarshalJSON(tx)
		if err != nil {
			return err
		}

		// Use Gnokey Mobile to sign.
		// Note that req.Msg.CallerAddress must be set to the desired signer. The app can get the
		// address using ListKeyInfo.
		signedTxJSON, err := s.gnokeyMobileClient.SignTx(
			context.Background(),
			connect.NewRequest(&api_gen.SignTxRequest{
				TxJson:  string(txJSON),
				Address: req.Msg.CallerAddress,
			}),
		)
		if err != nil {
			return err
		}
		signedTx := &std.Tx{}
		if err := amino.UnmarshalJSON([]byte(signedTxJSON.Msg.SignedTxJson), signedTx); err != nil {
			return err
		}

		// Now broadcast
		c, err := s.getClient(nil)
		if err != nil {
			return getGrpcError(err)
		}
		bres, err := c.BroadcastTxCommit(signedTx)
		if err != nil {
			return getGrpcError(err)
		}

		if err := stream.Send(&api_gen.CallResponse{
			Result: bres.DeliverTx.Data,
		}); err != nil {
			s.logger.Error("Call stream.Send returned error", zap.Error(err))
			return err
		}

		return nil
	}

	signer, err := s.getSigner(req.Msg.CallerAddress)
	if err != nil {
		return err
	}

	c, err := s.getClient(signer)
	if err != nil {
		return getGrpcError(err)
	}
	bres, err := c.Call(*cfg, msgs...)
	if err != nil {
		return getGrpcError(err)
	}

	if err := stream.Send(&api_gen.CallResponse{
		Result: bres.DeliverTx.Data,
	}); err != nil {
		s.logger.Error("Call stream.Send returned error", zap.Error(err))
		return err
	}

	return nil
}

func (s *gnoNativeService) convertCallRequest(req *api_gen.CallRequest) (*gnoclient.BaseTxCfg, []vm.MsgCall, error) {
	var callerAddress crypto.Address
	if req.CallerAddress != nil {
		callerAddress = crypto.AddressFromBytes(req.CallerAddress)
	} else {
		// Get the caller address from the active account
		s.lock.RLock()
		account := s.activeAccount
		s.lock.RUnlock()
		if account == nil {
			return nil, nil, api_gen.ErrCode_ErrNoActiveAccount
		}

		callerAddress = account.keyInfo.GetAddress()
	}
	cfg := &gnoclient.BaseTxCfg{
		GasFee:    req.GasFee,
		GasWanted: req.GasWanted,
		Memo:      req.Memo,
	}

	msgs := make([]vm.MsgCall, 0)

	for _, msg := range req.Msgs {
		send, err := std.ParseCoins(msg.Send)
		if err != nil {
			return nil, nil, getGrpcError(err)
		}

		msgs = append(msgs, vm.MsgCall{
			Caller:  callerAddress,
			PkgPath: msg.PackagePath,
			Func:    msg.Fnc,
			Args:    msg.Args,
			Send:    send,
		})
	}

	return cfg, msgs, nil
}

func (s *gnoNativeService) Send(ctx context.Context, req *connect.Request[api_gen.SendRequest], stream *connect.ServerStream[api_gen.SendResponse]) error {
	for _, msg := range req.Msg.Msgs {
		s.logger.Debug("Send", zap.String("toAddress", crypto.AddressToBech32(crypto.AddressFromBytes(msg.ToAddress))), zap.String("send", msg.Send))
	}

	signer, err := s.getSigner(req.Msg.CallerAddress)
	if err != nil {
		return err
	}

	cfg, msgs, err := s.convertSendRequest(req.Msg)
	if err != nil {
		return err
	}

	c, err := s.getClient(signer)
	if err != nil {
		return getGrpcError(err)
	}
	_, err = c.Send(*cfg, msgs...)
	if err != nil {
		return getGrpcError(err)
	}

	if err := stream.Send(&api_gen.SendResponse{}); err != nil {
		s.logger.Error("Send stream.Send returned error", zap.Error(err))
		return err
	}

	return nil
}

func (s *gnoNativeService) convertSendRequest(req *api_gen.SendRequest) (*gnoclient.BaseTxCfg, []bank.MsgSend, error) {
	var callerAddress crypto.Address
	if req.CallerAddress != nil {
		callerAddress = crypto.AddressFromBytes(req.CallerAddress)
	} else {
		// Get the caller address from the active account
		s.lock.RLock()
		account := s.activeAccount
		s.lock.RUnlock()
		if account == nil {
			return nil, nil, api_gen.ErrCode_ErrNoActiveAccount
		}

		callerAddress = account.keyInfo.GetAddress()
	}
	cfg := &gnoclient.BaseTxCfg{
		GasFee:    req.GasFee,
		GasWanted: req.GasWanted,
		Memo:      req.Memo,
	}

	msgs := make([]bank.MsgSend, 0)

	for _, msg := range req.Msgs {
		send, err := std.ParseCoins(msg.Send)
		if err != nil {
			return nil, nil, getGrpcError(err)
		}

		msgs = append(msgs, bank.MsgSend{
			FromAddress: callerAddress,
			ToAddress:   crypto.AddressFromBytes(msg.ToAddress),
			Amount:      send,
		})
	}

	return cfg, msgs, nil
}

func (s *gnoNativeService) Run(ctx context.Context, req *connect.Request[api_gen.RunRequest], stream *connect.ServerStream[api_gen.RunResponse]) error {
	signer, err := s.getSigner(req.Msg.CallerAddress)
	if err != nil {
		return err
	}

	cfg, msgs, err := s.convertRunRequest(req.Msg)
	if err != nil {
		return err
	}

	c, err := s.getClient(signer)
	if err != nil {
		return getGrpcError(err)
	}
	bres, err := c.Run(*cfg, msgs...)
	if err != nil {
		return getGrpcError(err)
	}

	if err := stream.Send(&api_gen.RunResponse{
		Result: string(bres.DeliverTx.Data),
	}); err != nil {
		s.logger.Error("Run stream.Send returned error", zap.Error(err))
		return err
	}

	return nil
}

func (s *gnoNativeService) convertRunRequest(req *api_gen.RunRequest) (*gnoclient.BaseTxCfg, []vm.MsgRun, error) {
	var callerAddress crypto.Address
	if req.CallerAddress != nil {
		callerAddress = crypto.AddressFromBytes(req.CallerAddress)
	} else {
		// Get the caller address from the active account
		s.lock.RLock()
		account := s.activeAccount
		s.lock.RUnlock()
		if account == nil {
			return nil, nil, api_gen.ErrCode_ErrNoActiveAccount
		}

		callerAddress = account.keyInfo.GetAddress()
	}
	cfg := &gnoclient.BaseTxCfg{
		GasFee:    req.GasFee,
		GasWanted: req.GasWanted,
		Memo:      req.Memo,
	}

	msgs := make([]vm.MsgRun, 0)

	for _, msg := range req.Msgs {
		send, err := std.ParseCoins(msg.Send)
		if err != nil {
			return nil, nil, getGrpcError(err)
		}

		memPkg := &std.MemPackage{
			Files: []*std.MemFile{
				{
					Name: "main.gno",
					Body: msg.Package,
				},
			},
		}
		msgs = append(msgs, vm.MsgRun{
			Caller:  callerAddress,
			Package: memPkg,
			Send:    send,
		})
	}

	return cfg, msgs, nil
}

func (s *gnoNativeService) MakeCallTx(ctx context.Context, req *connect.Request[api_gen.CallRequest]) (*connect.Response[api_gen.MakeTxResponse], error) {
	for _, msg := range req.Msg.Msgs {
		s.logger.Debug("MakeCallTx", zap.String("package", msg.PackagePath), zap.String("function", msg.Fnc), zap.Any("args", msg.Args))
	}

	cfg, msgs, err := s.convertCallRequest(req.Msg)
	if err != nil {
		return nil, err
	}
	tx, err := gnoclient.NewCallTx(*cfg, msgs...)
	if err != nil {
		return nil, getGrpcError(err)
	}

	txJSON, err := amino.MarshalJSON(tx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&api_gen.MakeTxResponse{TxJson: string(txJSON)}), nil
}

func (s *gnoNativeService) MakeSendTx(ctx context.Context, req *connect.Request[api_gen.SendRequest]) (*connect.Response[api_gen.MakeTxResponse], error) {
	cfg, msgs, err := s.convertSendRequest(req.Msg)
	if err != nil {
		return nil, err
	}
	tx, err := gnoclient.NewSendTx(*cfg, msgs...)
	if err != nil {
		return nil, getGrpcError(err)
	}

	txJSON, err := amino.MarshalJSON(tx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&api_gen.MakeTxResponse{TxJson: string(txJSON)}), nil
}

func (s *gnoNativeService) MakeRunTx(ctx context.Context, req *connect.Request[api_gen.RunRequest]) (*connect.Response[api_gen.MakeTxResponse], error) {
	cfg, msgs, err := s.convertRunRequest(req.Msg)
	if err != nil {
		return nil, err
	}
	tx, err := gnoclient.NewRunTx(*cfg, msgs...)
	if err != nil {
		return nil, getGrpcError(err)
	}

	txJSON, err := amino.MarshalJSON(tx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&api_gen.MakeTxResponse{TxJson: string(txJSON)}), nil
}

func (s *gnoNativeService) SignTx(ctx context.Context, req *connect.Request[api_gen.SignTxRequest]) (*connect.Response[api_gen.SignTxResponse], error) {
	var tx std.Tx
	if err := amino.UnmarshalJSON([]byte(req.Msg.TxJson), &tx); err != nil {
		return nil, err
	}

	signedTx, err := s.ClientSignTx(tx, req.Msg.Address, req.Msg.AccountNumber, req.Msg.SequenceNumber)
	if err != nil {
		return nil, getGrpcError(err)
	}

	signedTxJSON, err := amino.MarshalJSON(signedTx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&api_gen.SignTxResponse{SignedTxJson: string(signedTxJSON)}), nil
}

func (s *gnoNativeService) ClientSignTx(tx std.Tx, addr []byte, accountNumber, sequenceNumber uint64) (*std.Tx, error) {
	signer, err := s.getSigner(addr)
	if err != nil {
		return nil, err
	}
	c := &gnoclient.Client{
		Signer:    signer,
		RPCClient: s.rpcClient,
	}
	return c.SignTx(tx, accountNumber, sequenceNumber)
}

func (s *gnoNativeService) BroadcastTxCommit(ctx context.Context, req *connect.Request[api_gen.BroadcastTxCommitRequest],
	stream *connect.ServerStream[api_gen.BroadcastTxCommitResponse]) error {
	signedTx := &std.Tx{}
	if err := amino.UnmarshalJSON([]byte(req.Msg.SignedTxJson), signedTx); err != nil {
		return err
	}

	c, err := s.getClient(nil)
	if err != nil {
		return getGrpcError(err)
	}
	bres, err := c.BroadcastTxCommit(signedTx)
	if err != nil {
		return getGrpcError(err)
	}

	if err := stream.Send(&api_gen.BroadcastTxCommitResponse{
		Result: bres.DeliverTx.Data,
	}); err != nil {
		s.logger.Error("BroadcastTxCommit stream.Send returned error", zap.Error(err))
		return err
	}

	return nil
}

func (s *gnoNativeService) AddressToBech32(ctx context.Context, req *connect.Request[api_gen.AddressToBech32Request]) (*connect.Response[api_gen.AddressToBech32Response], error) {
	s.logger.Debug("AddressToBech32", zap.ByteString("address", req.Msg.Address))
	bech32Address := crypto.AddressToBech32(crypto.AddressFromBytes(req.Msg.Address))
	return connect.NewResponse(&api_gen.AddressToBech32Response{Bech32Address: bech32Address}), nil
}

func (s *gnoNativeService) AddressFromBech32(ctx context.Context, req *connect.Request[api_gen.AddressFromBech32Request]) (*connect.Response[api_gen.AddressFromBech32Response], error) {
	address, err := crypto.AddressFromBech32(req.Msg.Bech32Address)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&api_gen.AddressFromBech32Response{Address: address.Bytes()}), nil
}

func (s *gnoNativeService) AddressFromMnemonic(ctx context.Context, req *connect.Request[api_gen.AddressFromMnemonicRequest]) (*connect.Response[api_gen.AddressFromMnemonicResponse], error) {
	kb := crypto_keys.NewInMemory()
	info, err := kb.CreateAccount("temporary", req.Msg.Mnemonic, "", "", uint32(0), uint32(0))
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&api_gen.AddressFromMnemonicResponse{Address: info.GetAddress().Bytes()}), nil
}

func (s *gnoNativeService) Hello(ctx context.Context, req *connect.Request[api_gen.HelloRequest]) (*connect.Response[api_gen.HelloResponse], error) {
	s.logger.Debug("Hello called")
	defer s.logger.Debug("Hello returned ok")
	return connect.NewResponse(&api_gen.HelloResponse{
		Greeting: "Hello " + req.Msg.Name,
	}), nil
}

// HelloStream is for debug purposes
func (s *gnoNativeService) HelloStream(ctx context.Context, req *connect.Request[api_gen.HelloStreamRequest], stream *connect.ServerStream[api_gen.HelloStreamResponse]) error {
	s.logger.Debug("HelloStream called")
	for i := 0; i < 4; i++ {
		if err := stream.Send(&api_gen.HelloStreamResponse{
			Greeting: "Hello " + req.Msg.Name,
		}); err != nil {
			s.logger.Error("HelloStream returned error", zap.Error(err))
			return err
		}
		time.Sleep(2 * time.Second)
	}

	s.logger.Debug("HelloStream returned ok")
	return nil
}

// If err is a recognized Go error, return the equivalent Grpc error.
// Otherwise, just return err.
func getGrpcError(err error) error {
	if keyerror.IsErrKeyNotFound(err) {
		return api_gen.ErrCode_ErrCryptoKeyNotFound
	} else if keyerror.IsErrWrongPassword(err) {
		return api_gen.ErrCode_ErrDecryptionFailed
	}

	// The following match errors in https://github.com/gnolang/gno/blob/master/tm2/pkg/std/errors.go .
	if errors.As(err, &std.TxDecodeError{}) {
		return api_gen.ErrCode_ErrTxDecode
	} else if errors.As(err, &std.InvalidSequenceError{}) {
		return api_gen.ErrCode_ErrInvalidSequence
	} else if errors.As(err, &std.UnauthorizedError{}) {
		return api_gen.ErrCode_ErrUnauthorized
	} else if errors.As(err, &std.InsufficientFundsError{}) {
		return api_gen.ErrCode_ErrInsufficientFunds
	} else if errors.As(err, &std.UnknownRequestError{}) {
		return api_gen.ErrCode_ErrUnknownRequest
	} else if errors.As(err, &std.InvalidAddressError{}) {
		return api_gen.ErrCode_ErrInvalidAddress
	} else if errors.As(err, &std.UnknownAddressError{}) {
		return api_gen.ErrCode_ErrUnknownAddress
	} else if errors.As(err, &std.InvalidPubKeyError{}) {
		return api_gen.ErrCode_ErrInvalidPubKey
	} else if errors.As(err, &std.InsufficientCoinsError{}) {
		return api_gen.ErrCode_ErrInsufficientCoins
	} else if errors.As(err, &std.InvalidCoinsError{}) {
		return api_gen.ErrCode_ErrInvalidCoins
	} else if errors.As(err, &std.InvalidGasWantedError{}) {
		return api_gen.ErrCode_ErrInvalidGasWanted
	} else if errors.As(err, &std.OutOfGasError{}) {
		return api_gen.ErrCode_ErrOutOfGas
	} else if errors.As(err, &std.MemoTooLargeError{}) {
		return api_gen.ErrCode_ErrMemoTooLarge
	} else if errors.As(err, &std.InsufficientFeeError{}) {
		return api_gen.ErrCode_ErrInsufficientFee
	} else if errors.As(err, &std.TooManySignaturesError{}) {
		return api_gen.ErrCode_ErrTooManySignatures
	} else if errors.As(err, &std.NoSignaturesError{}) {
		return api_gen.ErrCode_ErrNoSignatures
	} else if errors.As(err, &std.GasOverflowError{}) {
		return api_gen.ErrCode_ErrGasOverflow
	}

	// The following match errors in https://github.com/gnolang/gno/blob/master/gno.land/pkg/sdk/vm/errors.go .

	if errors.As(err, &vm.InvalidPkgPathError{}) {
		return api_gen.ErrCode_ErrInvalidPkgPath
	} else if errors.As(err, &vm.InvalidStmtError{}) {
		return api_gen.ErrCode_ErrInvalidStmt
	} else if errors.As(err, &vm.InvalidExprError{}) {
		return api_gen.ErrCode_ErrInvalidExpr
	} else {
		return err
	}
}
