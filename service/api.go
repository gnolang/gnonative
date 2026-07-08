// This file implements the GnoNative API methods defined in api/rpc.proto . For documentation,
// see that file and related request/response fields in the generated api/gnonativetypes.proto .
//
// These are plain Go functions with no dependency on connect/gRPC. The connect handlers in
// api_connect.go are thin wrappers over these functions, and the mobile bridge dispatcher in
// framework/service/ calls them directly.

package service

import (
	"context"
	"errors"
	"slices"
	"strings"
	"time"

	"github.com/gnolang/gno/gno.land/pkg/gnoland"
	"github.com/gnolang/gno/gno.land/pkg/gnoland/ugnot"
	"github.com/gnolang/gno/gno.land/pkg/keyscli"
	"github.com/gnolang/gno/tm2/pkg/amino"
	"github.com/gnolang/gno/tm2/pkg/bech32"
	abci "github.com/gnolang/gno/tm2/pkg/bft/abci/types"
	"github.com/gnolang/gno/tm2/pkg/crypto"
	"github.com/gnolang/gno/tm2/pkg/crypto/bip39"
	crypto_keys "github.com/gnolang/gno/tm2/pkg/crypto/keys"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys/keyerror"
	"github.com/gnolang/gno/tm2/pkg/overflow"
	"github.com/gnolang/gno/tm2/pkg/sdk/auth"
	"github.com/gnolang/gno/tm2/pkg/sdk/bank"
	"github.com/gnolang/gno/tm2/pkg/std"
	"go.uber.org/zap"

	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/gno.land/pkg/sdk/vm"
	rpcclient "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	api_gen "github.com/gnolang/gnonative/v5/api"
)

func (s *gnoNativeService) SetRemote(ctx context.Context, req *api_gen.SetRemoteRequest) (*api_gen.SetRemoteResponse, error) {
	var err error
	s.rpcClient, err = rpcclient.NewHTTPClient(req.Remote)
	if err != nil {
		return nil, api_gen.ErrCode_ErrSetRemote.Wrap(err)
	}
	s.remote = req.Remote
	return &api_gen.SetRemoteResponse{}, nil
}

func (s *gnoNativeService) GetRemote(ctx context.Context, req *api_gen.GetRemoteRequest) (*api_gen.GetRemoteResponse, error) {
	return &api_gen.GetRemoteResponse{Remote: s.remote}, nil
}

func (s *gnoNativeService) SetChainID(ctx context.Context, req *api_gen.SetChainIDRequest) (*api_gen.SetChainIDResponse, error) {
	s.lock.Lock()
	s.chainID = req.ChainID
	s.lock.Unlock()
	return &api_gen.SetChainIDResponse{}, nil
}

func (s *gnoNativeService) GetChainID(ctx context.Context, req *api_gen.GetChainIDRequest) (*api_gen.GetChainIDResponse, error) {
	return &api_gen.GetChainIDResponse{ChainID: s.chainID}, nil
}

func (s *gnoNativeService) GenerateRecoveryPhrase(ctx context.Context, req *api_gen.GenerateRecoveryPhraseRequest) (*api_gen.GenerateRecoveryPhraseResponse, error) {
	const mnemonicEntropySize = 256
	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return nil, err
	}

	phrase, err := bip39.NewMnemonic(entropySeed[:])
	if err != nil {
		return nil, err
	}

	return &api_gen.GenerateRecoveryPhraseResponse{Phrase: phrase}, nil
}

func ConvertKeyInfo(key crypto_keys.Info) (*api_gen.KeyInfo, error) {
	return &api_gen.KeyInfo{
		Type:    uint32(key.GetType()),
		Name:    key.GetName(),
		Address: key.GetAddress().Bytes(),
		PubKey:  key.GetPubKey().Bytes(),
	}, nil
}

func (s *gnoNativeService) ListKeyInfo(ctx context.Context, req *api_gen.ListKeyInfoRequest) (*api_gen.ListKeyInfoResponse, error) {
	s.logger.Debug("ListKeyInfo called")

	keys, err := s.keybase.List()
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

	return &api_gen.ListKeyInfoResponse{Keys: formatedKeys}, nil
}

func (s *gnoNativeService) HasKeyByName(ctx context.Context, req *api_gen.HasKeyByNameRequest) (*api_gen.HasKeyByNameResponse, error) {
	s.logger.Debug("HasKeyByName called")

	has, err := s.keybase.HasByName(req.Name)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return &api_gen.HasKeyByNameResponse{Has: has}, nil
}

func (s *gnoNativeService) HasKeyByAddress(ctx context.Context, req *api_gen.HasKeyByAddressRequest) (*api_gen.HasKeyByAddressResponse, error) {
	s.logger.Debug("HasKeyByAddress called")

	addr, err := crypto.AddressFromBytes(req.Address)
	if err != nil {
		return nil, getGrpcError(err)
	}
	has, err := s.keybase.HasByAddress(addr)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return &api_gen.HasKeyByAddressResponse{Has: has}, nil
}

func (s *gnoNativeService) HasKeyByNameOrAddress(ctx context.Context, req *api_gen.HasKeyByNameOrAddressRequest) (*api_gen.HasKeyByNameOrAddressResponse, error) {
	s.logger.Debug("HasKeyByNameOrAddress called")

	has, err := s.keybase.HasByNameOrAddress(req.NameOrBech32)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return &api_gen.HasKeyByNameOrAddressResponse{Has: has}, nil
}

func (s *gnoNativeService) GetKeyInfoByName(ctx context.Context, req *api_gen.GetKeyInfoByNameRequest) (*api_gen.GetKeyInfoByNameResponse, error) {
	s.logger.Debug("GetKeyInfoByName called")

	key, err := s.keybase.GetByName(req.Name)
	if err != nil {
		return nil, getGrpcError(err)
	}

	info, err := ConvertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	return &api_gen.GetKeyInfoByNameResponse{Key: info}, nil
}

func (s *gnoNativeService) GetKeyInfoByAddress(ctx context.Context, req *api_gen.GetKeyInfoByAddressRequest) (*api_gen.GetKeyInfoByAddressResponse, error) {
	s.logger.Debug("GetKeyInfoByAddress called")

	addr, err := crypto.AddressFromBytes(req.Address)
	if err != nil {
		return nil, getGrpcError(err)
	}
	key, err := s.keybase.GetByAddress(addr)
	if err != nil {
		return nil, getGrpcError(err)
	}

	info, err := ConvertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	return &api_gen.GetKeyInfoByAddressResponse{Key: info}, nil
}

func (s *gnoNativeService) GetKeyInfoByNameOrAddress(ctx context.Context, req *api_gen.GetKeyInfoByNameOrAddressRequest) (*api_gen.GetKeyInfoByNameOrAddressResponse, error) {
	s.logger.Debug("GetKeyInfoByNameOrAddress called")

	key, err := s.keybase.GetByNameOrAddress(req.NameOrBech32)
	if err != nil {
		return nil, getGrpcError(err)
	}

	info, err := ConvertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	return &api_gen.GetKeyInfoByNameOrAddressResponse{Key: info}, nil
}

func (s *gnoNativeService) CreateAccount(ctx context.Context, req *api_gen.CreateAccountRequest) (*api_gen.CreateAccountResponse, error) {
	s.logger.Debug("CreateAccount called", zap.String("NameOrBech32", req.NameOrBech32))

	key, err := s.keybase.CreateAccount(req.NameOrBech32, req.Mnemonic, req.Bip39Passwd, req.Password, req.Account, req.Index)
	if err != nil {
		return nil, getGrpcError(err)
	}

	info, err := ConvertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	return &api_gen.CreateAccountResponse{Key: info}, nil
}

func (s *gnoNativeService) CreateLedger(ctx context.Context, req *api_gen.CreateLedgerRequest) (*api_gen.CreateLedgerResponse, error) {
	s.logger.Debug("CreateLedger called", zap.String("Name", req.Name))

	key, err := s.keybase.CreateLedger(req.Name, crypto_keys.SigningAlgo(req.Algorithm), req.HRP, req.Account, req.Index)
	if err != nil {
		return nil, getGrpcError(err)
	}

	info, err := ConvertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	return &api_gen.CreateLedgerResponse{Key: info}, nil
}

func (s *gnoNativeService) ActivateAccount(ctx context.Context, req *api_gen.ActivateAccountRequest) (*api_gen.ActivateAccountResponse, error) {
	s.logger.Debug("ActivateAccount called", zap.String("NameOrBech32", req.NameOrBech32))

	// The key may already be in s.userAccounts, but the info may have changed on disk. So always get from disk.
	key, err := s.keybase.GetByNameOrAddress(req.NameOrBech32)
	if err != nil {
		return nil, getGrpcError(err)
	}

	info, err := ConvertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	b32, err := bech32.Encode(crypto.Bech32AddrPrefix(), key.GetAddress().Bytes())
	if err != nil {
		return nil, getGrpcError(err)
	}
	s.lock.Lock()
	account, ok := s.userAccounts[b32]
	if !ok {
		account = &userAccount{}
		account.signer = &gnoclient.SignerFromKeybase{
			Keybase: s.keybase,
			ChainID: s.chainID,
		}
		s.userAccounts[b32] = account
	}
	account.keyInfo = key
	s.lock.Unlock()

	account.signer.Account = req.NameOrBech32
	if req.Master != nil {
		masterAddr, err := crypto.AddressFromBytes(req.Master)
		if err != nil {
			return nil, getGrpcError(err)
		}
		account.signer.Master = masterAddr
	} else {
		// Clear any stale master from a previous activation as a session account.
		account.signer.Master = crypto.Address{}
	}
	return &api_gen.ActivateAccountResponse{
		Key:         info,
		HasPassword: account.signer.Password != "",
	}, nil
}

func (s *gnoNativeService) SetPassword(ctx context.Context, req *api_gen.SetPasswordRequest) (*api_gen.SetPasswordResponse, error) {
	signer, err := s.getSigner(req.Address)
	if err != nil {
		return nil, err
	}
	signer.Password = req.Password

	// Check the password.
	if err := signer.Validate(); err != nil {
		if keyerror.IsErrWrongPassword(err) {
			// Wrong password, so unset the password
			signer.Password = ""
		}
		return nil, getGrpcError(err)
	}

	return &api_gen.SetPasswordResponse{}, nil
}

func (s *gnoNativeService) RenameKey(ctx context.Context, req *api_gen.RenameKeyRequest) (*api_gen.RenameKeyResponse, error) {
	// We may need to change the key name in s.userAccounts, but we don't know the address. So get from disk.
	key, err := s.keybase.GetByName(req.OldName)
	if err != nil {
		return nil, getGrpcError(err)
	}

	err = s.keybase.Rename(req.OldName, req.NewName)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			return nil, api_gen.ErrCode_ErrKeyNameExists.Wrap(err)
		}
		return nil, getGrpcError(err)
	}

	b32, err := bech32.Encode(crypto.Bech32AddrPrefix(), key.GetAddress().Bytes())
	if err != nil {
		return nil, getGrpcError(err)
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	if account, ok := s.userAccounts[b32]; ok {
		// Get the key info and update the account with the new name
		newKeyInfo, err := s.keybase.GetByAddress(key.GetAddress())
		if err != nil {
			return nil, getGrpcError(err)
		}

		account.keyInfo = newKeyInfo
	}

	return &api_gen.RenameKeyResponse{}, nil
}

func (s *gnoNativeService) RotatePassword(ctx context.Context, req *api_gen.RotatePasswordRequest) (*api_gen.RotatePasswordResponse, error) {
	// Get all the signers, before trying to update the password.
	signers := make([]*gnoclient.SignerFromKeybase, len(req.Addresses))
	for i := range len(req.Addresses) {
		var err error
		if signers[i], err = s.getSigner(req.Addresses[i]); err != nil {
			return nil, err
		}
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	getNewPassword := func() (string, error) { return req.NewPassword, nil }
	for i := range len(req.Addresses) {
		if err := s.keybase.Rotate(signers[i].Account, signers[i].Password, getNewPassword); err != nil {
			// Roll back the passwords. Don't check the error from Rotate.
			for j := range i {
				getOldPassword := func() (string, error) { return signers[j].Password, nil }
				s.keybase.Rotate(signers[j].Account, req.NewPassword, getOldPassword)
			}
			return nil, getGrpcError(err)
		}
	}

	// Success. Update the Password in all the signers.
	for i := range len(req.Addresses) {
		signers[i].Password = req.NewPassword
	}

	return &api_gen.RotatePasswordResponse{}, nil
}

func (s *gnoNativeService) GetActivatedAccount(ctx context.Context, req *api_gen.GetActivatedAccountRequest) (*api_gen.GetActivatedAccountResponse, error) {
	s.logger.Debug("GetActivatedAccount called")

	if req.Address == nil {
		return nil, api_gen.ErrCode_ErrInvalidAddress
	}

	b32, err := bech32.Encode(crypto.Bech32AddrPrefix(), req.Address)
	if err != nil {
		return nil, getGrpcError(err)
	}
	s.lock.Lock()
	account, ok := s.userAccounts[b32]
	s.lock.Unlock()
	if !ok {
		return nil, api_gen.ErrCode_ErrNoActiveAccount
	}

	info, err := ConvertKeyInfo(account.keyInfo)
	if err != nil {
		return nil, err
	}

	var master []byte
	if !account.signer.Master.IsZero() {
		master = account.signer.Master.Bytes()
	}

	return &api_gen.GetActivatedAccountResponse{
		Key:         info,
		Master:      master,
		HasPassword: account.signer.Password != "",
	}, nil
}

func (s *gnoNativeService) QueryAccount(ctx context.Context, req *api_gen.QueryAccountRequest) (*api_gen.QueryAccountResponse, error) {
	s.logger.Debug("QueryAccount", zap.ByteString("address", req.Address))

	c, err := s.getClient(nil)
	if err != nil {
		return nil, getGrpcError(err)
	}
	// gnoclient wants the crypto.Address.
	addr, err := crypto.AddressFromBytes(req.Address)
	if err != nil {
		return nil, getGrpcError(err)
	}
	account, _, err := c.QueryAccount(addr)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return &api_gen.QueryAccountResponse{AccountInfo: convertBaseAccount(account)}, nil
}

func (s *gnoNativeService) QuerySessionAccount(ctx context.Context, req *api_gen.QuerySessionAccountRequest) (*api_gen.QuerySessionAccountResponse, error) {
	s.logger.Debug("QuerySessionAccount", zap.ByteString("master-address", req.MasterAddress), zap.ByteString("session-address", req.SessionAddress))

	c, err := s.getClient(nil)
	if err != nil {
		return nil, getGrpcError(err)
	}
	// gnoclient wants the crypto.Address.
	masterAddr, err := crypto.AddressFromBytes(req.MasterAddress)
	if err != nil {
		return nil, getGrpcError(err)
	}
	sessionAddr, err := crypto.AddressFromBytes(req.SessionAddress)
	if err != nil {
		return nil, getGrpcError(err)
	}
	account, _, err := c.QuerySessionAccount(masterAddr, sessionAddr)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return &api_gen.QuerySessionAccountResponse{AccountInfo: convertSessionAccount(account)}, nil
}

func convertBaseAccount(account *std.BaseAccount) *api_gen.BaseAccount {
	var pubKeyBytes []byte
	if account.PubKey != nil {
		pubKeyBytes = account.PubKey.Bytes()
	}

	return &api_gen.BaseAccount{
		Address:       account.Address.Bytes(),
		Coins:         convertStdCoins(account.Coins),
		PubKey:        pubKeyBytes,
		AccountNumber: account.AccountNumber,
		Sequence:      account.Sequence,
	}
}

func convertSessionAccount(account *gnoland.GnoSessionAccount) *api_gen.SessionAccount {
	return &api_gen.SessionAccount{
		BaseAccount:   convertBaseAccount(&account.BaseAccount),
		MasterAddress: account.MasterAddress.Bytes(),
		ExpiresAt:     account.ExpiresAt,
		SpendLimit:    convertStdCoins(account.SpendLimit),
		SpendPeriod:   account.SpendPeriod,
		SpendUsed:     convertStdCoins(account.SpendUsed),
		SpendReset:    account.SpendReset,
		AllowPaths:    account.AllowPaths,
	}
}

func (s *gnoNativeService) DeleteAccount(ctx context.Context, req *api_gen.DeleteAccountRequest) (*api_gen.DeleteAccountResponse, error) {
	// Get the key from the Keybase so that we know its address
	key, err := s.keybase.GetByNameOrAddress(req.NameOrBech32)
	if err != nil {
		return nil, getGrpcError(err)
	}
	if err := s.keybase.Delete(req.NameOrBech32, req.Password, req.SkipPassword); err != nil {
		return nil, getGrpcError(err)
	}

	b32, err := bech32.Encode(crypto.Bech32AddrPrefix(), key.GetAddress().Bytes())
	if err != nil {
		return nil, getGrpcError(err)
	}
	s.lock.Lock()
	delete(s.userAccounts, b32)
	s.lock.Unlock()
	return &api_gen.DeleteAccountResponse{}, nil
}

func (s *gnoNativeService) Query(ctx context.Context, req *api_gen.QueryRequest) (*api_gen.QueryResponse, error) {
	s.logger.Debug("Query", zap.String("path", req.Path), zap.ByteString("data", req.Data))

	cfg := gnoclient.QueryCfg{
		Path: req.Path,
		Data: req.Data,
	}

	c, err := s.getClient(nil)
	if err != nil {
		return nil, getGrpcError(err)
	}
	bres, err := c.Query(cfg)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return &api_gen.QueryResponse{Result: bres.Response.Data}, nil
}

func (s *gnoNativeService) Render(ctx context.Context, req *api_gen.RenderRequest) (*api_gen.RenderResponse, error) {
	s.logger.Debug("Render", zap.String("packagePath", req.PackagePath), zap.String("args", req.Args))

	c, err := s.getClient(nil)
	if err != nil {
		return nil, getGrpcError(err)
	}
	result, _, err := c.Render(req.PackagePath, req.Args)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return &api_gen.RenderResponse{Result: result}, nil
}

func (s *gnoNativeService) QEval(ctx context.Context, req *api_gen.QEvalRequest) (*api_gen.QEvalResponse, error) {
	s.logger.Debug("QEval", zap.String("packagePath", req.PackagePath), zap.String("expression", req.Expression))

	c, err := s.getClient(nil)
	if err != nil {
		return nil, getGrpcError(err)
	}
	result, _, err := c.QEval(req.PackagePath, req.Expression)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return &api_gen.QEvalResponse{Result: result}, nil
}

func (s *gnoNativeService) Call(ctx context.Context, req *api_gen.CallRequest, send func(*api_gen.CallResponse) error) error {
	for _, msg := range req.Msgs {
		s.logger.Debug("Call", zap.String("package", msg.PackagePath), zap.String("function", msg.Fnc), zap.Any("args", msg.Args))
	}

	cfg := &gnoclient.BaseTxCfg{
		GasFee:    req.GasFee,
		GasWanted: req.GasWanted,
		Memo:      req.Memo,
	}

	signer, err := s.getSigner(req.SignerAddress)
	if err != nil {
		return err
	}

	var callerAddress []byte
	if !signer.Master.IsZero() {
		callerAddress = signer.Master.Bytes()
	} else {
		callerAddress = req.SignerAddress
	}
	msgs, err := convertCallMsgs(callerAddress, req.Msgs)
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

	if err := send(&api_gen.CallResponse{
		Result: bres.DeliverTx.Data,
		Hash:   bres.Hash,
		Height: bres.Height,
	}); err != nil {
		return err
	}

	return nil
}

// convertCallMsgs converts to use vm.MsgCall. Ignore req.SignerAddress and use callerAddress.
func convertCallMsgs(callerAddress []byte, callMsgs []*api_gen.MsgCall) ([]vm.MsgCall, error) {
	addr, err := crypto.AddressFromBytes(callerAddress)
	if err != nil {
		return nil, getGrpcError(err)
	}

	msgs := make([]vm.MsgCall, 0)

	for _, msg := range callMsgs {
		sendCoins, err := convertCoins(msg.Send)
		if err != nil {
			return nil, err
		}
		maxDepositCoins, err := convertCoins(msg.MaxDeposit)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, vm.MsgCall{
			Caller:     addr,
			PkgPath:    msg.PackagePath,
			Func:       msg.Fnc,
			Args:       msg.Args,
			Send:       sendCoins,
			MaxDeposit: maxDepositCoins,
		})
	}

	return msgs, nil
}

func (s *gnoNativeService) Send(ctx context.Context, req *api_gen.SendRequest, send func(*api_gen.SendResponse) error) error {
	for _, msg := range req.Msgs {
		for _, coin := range msg.Amount {
			if b32, err := bech32.Encode(crypto.Bech32AddrPrefix(), msg.ToAddress); err == nil {
				s.logger.Debug("Send", zap.String("toAddress", b32), zap.String("denom", coin.Denom),
					zap.Int64("amount", coin.Amount))
			}
		}
	}

	cfg := &gnoclient.BaseTxCfg{
		GasFee:    req.GasFee,
		GasWanted: req.GasWanted,
		Memo:      req.Memo,
	}

	signer, err := s.getSigner(req.SignerAddress)
	if err != nil {
		return err
	}

	var callerAddress []byte
	if !signer.Master.IsZero() {
		callerAddress = signer.Master.Bytes()
	} else {
		callerAddress = req.SignerAddress
	}
	msgs, err := convertSendMsgs(callerAddress, req.Msgs)
	if err != nil {
		return err
	}

	c, err := s.getClient(signer)
	if err != nil {
		return getGrpcError(err)
	}
	bres, err := c.Send(*cfg, msgs...)
	if err != nil {
		return getGrpcError(err)
	}

	if err := send(&api_gen.SendResponse{
		Hash:   bres.Hash,
		Height: bres.Height,
	}); err != nil {
		s.logger.Error("Send send returned error", zap.Error(err))
		return err
	}

	return nil
}

// convertSendMsgs converts to use bank.MsgSend. Ignore req.SignerAddress and use callerAddress.
func convertSendMsgs(callerAddress []byte, sendMsgs []*api_gen.MsgSend) ([]bank.MsgSend, error) {
	fromAddr, err := crypto.AddressFromBytes(callerAddress)
	if err != nil {
		return nil, getGrpcError(err)
	}

	msgs := make([]bank.MsgSend, 0)

	for _, msg := range sendMsgs {
		toAddr, err := crypto.AddressFromBytes(msg.ToAddress)
		if err != nil {
			return nil, getGrpcError(err)
		}
		amountCoins, err := convertCoins(msg.Amount)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, bank.MsgSend{
			FromAddress: fromAddr,
			ToAddress:   toAddr,
			Amount:      amountCoins,
		})
	}

	return msgs, nil
}

func (s *gnoNativeService) Run(ctx context.Context, req *api_gen.RunRequest, send func(*api_gen.RunResponse) error) error {
	cfg := &gnoclient.BaseTxCfg{
		GasFee:    req.GasFee,
		GasWanted: req.GasWanted,
		Memo:      req.Memo,
	}

	signer, err := s.getSigner(req.SignerAddress)
	if err != nil {
		return err
	}

	var callerAddress []byte
	if !signer.Master.IsZero() {
		callerAddress = signer.Master.Bytes()
	} else {
		callerAddress = req.SignerAddress
	}
	msgs, err := convertRunMsgs(callerAddress, req.Msgs)
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

	if err := send(&api_gen.RunResponse{
		Result: string(bres.DeliverTx.Data),
		Hash:   bres.Hash,
		Height: bres.Height,
	}); err != nil {
		s.logger.Error("Run send returned error", zap.Error(err))
		return err
	}

	return nil
}

// convertRunMsgs converts to use vm.MsgRun. Ignore req.SignerAddress and use callerAddress.
func convertRunMsgs(callerAddress []byte, runMsgs []*api_gen.MsgRun) ([]vm.MsgRun, error) {
	addr, err := crypto.AddressFromBytes(callerAddress)
	if err != nil {
		return nil, getGrpcError(err)
	}

	msgs := make([]vm.MsgRun, 0)

	for _, msg := range runMsgs {
		memPkg := &std.MemPackage{
			Name: "main",
			// Path will be automatically set by handler.
			Files: []*std.MemFile{
				{
					Name: "main.gno",
					Body: msg.Package,
				},
			},
		}
		sendCoins, err := convertCoins(msg.Send)
		if err != nil {
			return nil, err
		}
		maxDepositCoins, err := convertCoins(msg.MaxDeposit)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, vm.MsgRun{
			Caller:     addr,
			Package:    memPkg,
			Send:       sendCoins,
			MaxDeposit: maxDepositCoins,
		})
	}

	return msgs, nil
}

// convertCoins converts an array of api_gen.Coin to an array of std.Coin
// If the values are invalid, return a gRPC error
func convertCoins(apiGenCoins []*api_gen.Coin) ([]std.Coin, error) {
	coins := make([]std.Coin, 0)
	for _, coin := range apiGenCoins {
		stdCoin, err := std.NewCoinSafe(coin.Denom, coin.Amount)
		if err != nil {
			return nil, getGrpcError(err)
		}
		coins = append(coins, stdCoin)
	}
	return coins, nil
}

// convertStdCoins converts an array of std.Coin to an array of api_gen.Coin.
// This is the reverse of convertCoins.
func convertStdCoins(stdCoins []std.Coin) []*api_gen.Coin {
	coins := make([]*api_gen.Coin, 0)
	for _, coin := range stdCoins {
		coins = append(coins, &api_gen.Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount,
		})
	}

	return coins
}

func (s *gnoNativeService) MakeCallTx(ctx context.Context, req *api_gen.MakeCallTxRequest) (*api_gen.MakeTxResponse, error) {
	for _, msg := range req.Msgs {
		s.logger.Debug("MakeCallTx", zap.String("package", msg.PackagePath), zap.String("function", msg.Fnc), zap.Any("args", msg.Args))
	}

	cfg := &gnoclient.BaseTxCfg{
		GasFee:    req.GasFee,
		GasWanted: req.GasWanted,
		Memo:      req.Memo,
	}
	msgs, err := convertCallMsgs(req.CallerAddress, req.Msgs)
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
	return &api_gen.MakeTxResponse{TxJSON: string(txJSON)}, nil
}

func (s *gnoNativeService) MakeSendTx(ctx context.Context, req *api_gen.MakeSendTxRequest) (*api_gen.MakeTxResponse, error) {
	cfg := &gnoclient.BaseTxCfg{
		GasFee:    req.GasFee,
		GasWanted: req.GasWanted,
		Memo:      req.Memo,
	}
	msgs, err := convertSendMsgs(req.CallerAddress, req.Msgs)
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
	return &api_gen.MakeTxResponse{TxJSON: string(txJSON)}, nil
}

func (s *gnoNativeService) MakeRunTx(ctx context.Context, req *api_gen.MakeRunTxRequest) (*api_gen.MakeTxResponse, error) {
	cfg := &gnoclient.BaseTxCfg{
		GasFee:    req.GasFee,
		GasWanted: req.GasWanted,
		Memo:      req.Memo,
	}
	msgs, err := convertRunMsgs(req.CallerAddress, req.Msgs)
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
	return &api_gen.MakeTxResponse{TxJSON: string(txJSON)}, nil
}

func (s *gnoNativeService) CreateSession(ctx context.Context, req *api_gen.CreateSessionRequest, send func(*api_gen.CreateSessionResponse) error) error {
	for _, msg := range req.Msgs {
		s.logger.Debug("CreateSession", zap.ByteString("creator", req.CreatorAddress), zap.ByteString("sessionKey", msg.SessionKey))
	}

	cfg, msgs, err := s.convertCreateSessionRequest(req)
	if err != nil {
		return err
	}

	creator, err := s.getSigner(req.CreatorAddress)
	if err != nil {
		return err
	}

	c, err := s.getClient(creator)
	if err != nil {
		return getGrpcError(err)
	}
	bres, err := c.CreateSession(*cfg, msgs...)
	if err != nil {
		return getGrpcError(err)
	}

	if err := send(&api_gen.CreateSessionResponse{
		Result: bres.DeliverTx.Data,
		Hash:   bres.Hash,
		Height: bres.Height,
	}); err != nil {
		return err
	}

	return nil
}

func (s *gnoNativeService) MakeCreateSessionTx(ctx context.Context, req *api_gen.CreateSessionRequest) (*api_gen.MakeTxResponse, error) {
	for _, msg := range req.Msgs {
		s.logger.Debug("MakeCreateSessionTx", zap.ByteString("creator", req.CreatorAddress), zap.ByteString("sessionKey", msg.SessionKey))
	}

	cfg, msgs, err := s.convertCreateSessionRequest(req)
	if err != nil {
		return nil, err
	}
	tx, err := gnoclient.NewCreateSessionTx(*cfg, msgs...)
	if err != nil {
		return nil, getGrpcError(err)
	}

	txJSON, err := amino.MarshalJSON(tx)
	if err != nil {
		return nil, err
	}
	return &api_gen.MakeTxResponse{TxJSON: string(txJSON)}, nil
}

func (s *gnoNativeService) convertCreateSessionRequest(req *api_gen.CreateSessionRequest) (*gnoclient.BaseTxCfg, []auth.MsgCreateSession, error) {
	creator, err := crypto.AddressFromBytes(req.CreatorAddress)
	if err != nil {
		return nil, nil, getGrpcError(err)
	}

	cfg := &gnoclient.BaseTxCfg{
		GasFee:    req.GasFee,
		GasWanted: req.GasWanted,
		Memo:      req.Memo,
	}

	msgs := make([]auth.MsgCreateSession, 0)

	for _, msg := range req.Msgs {
		sessionKey, err := crypto.PubKeyFromBytes(msg.SessionKey)
		if err != nil {
			return nil, nil, err
		}
		spendLimitCoins, err := convertCoins(msg.SpendLimit)
		if err != nil {
			return nil, nil, err
		}
		msgs = append(msgs, auth.MsgCreateSession{
			Creator:     creator,
			SessionKey:  sessionKey,
			ExpiresAt:   msg.ExpiresAt,
			AllowPaths:  msg.AllowPaths,
			SpendLimit:  spendLimitCoins,
			SpendPeriod: msg.SpendPeriod,
		})
	}

	return cfg, msgs, nil
}

func (s *gnoNativeService) RevokeSession(ctx context.Context, req *api_gen.RevokeSessionRequest, send func(*api_gen.RevokeSessionResponse) error) error {
	for _, msg := range req.Msgs {
		s.logger.Debug("RevokeSession", zap.ByteString("creator", req.CreatorAddress), zap.ByteString("sessionKey", msg.SessionKey))
	}

	cfg, msgs, err := s.convertRevokeSessionRequest(req)
	if err != nil {
		return err
	}

	creator, err := s.getSigner(req.CreatorAddress)
	if err != nil {
		return err
	}

	c, err := s.getClient(creator)
	if err != nil {
		return getGrpcError(err)
	}
	bres, err := c.RevokeSession(*cfg, msgs...)
	if err != nil {
		return getGrpcError(err)
	}

	if err := send(&api_gen.RevokeSessionResponse{
		Result: bres.DeliverTx.Data,
		Hash:   bres.Hash,
		Height: bres.Height,
	}); err != nil {
		return err
	}

	return nil
}

func (s *gnoNativeService) MakeRevokeSessionTx(ctx context.Context, req *api_gen.RevokeSessionRequest) (*api_gen.MakeTxResponse, error) {
	for _, msg := range req.Msgs {
		s.logger.Debug("MakeRevokeSessionTx", zap.ByteString("creator", req.CreatorAddress), zap.ByteString("sessionKey", msg.SessionKey))
	}

	cfg, msgs, err := s.convertRevokeSessionRequest(req)
	if err != nil {
		return nil, err
	}
	tx, err := gnoclient.NewRevokeSessionTx(*cfg, msgs...)
	if err != nil {
		return nil, getGrpcError(err)
	}

	txJSON, err := amino.MarshalJSON(tx)
	if err != nil {
		return nil, err
	}
	return &api_gen.MakeTxResponse{TxJSON: string(txJSON)}, nil
}

func (s *gnoNativeService) convertRevokeSessionRequest(req *api_gen.RevokeSessionRequest) (*gnoclient.BaseTxCfg, []auth.MsgRevokeSession, error) {
	creator, err := crypto.AddressFromBytes(req.CreatorAddress)
	if err != nil {
		return nil, nil, getGrpcError(err)
	}

	cfg := &gnoclient.BaseTxCfg{
		GasFee:    req.GasFee,
		GasWanted: req.GasWanted,
		Memo:      req.Memo,
	}

	msgs := make([]auth.MsgRevokeSession, 0)

	for _, msg := range req.Msgs {
		sessionKey, err := crypto.PubKeyFromBytes(msg.SessionKey)
		if err != nil {
			return nil, nil, err
		}
		msgs = append(msgs, auth.MsgRevokeSession{
			Creator:    creator,
			SessionKey: sessionKey,
		})
	}

	return cfg, msgs, nil
}

func (s *gnoNativeService) RevokeAllSessions(ctx context.Context, req *api_gen.RevokeAllSessionsRequest, send func(*api_gen.RevokeAllSessionsResponse) error) error {
	s.logger.Debug("RevokeAllSessions", zap.ByteString("creator", req.CreatorAddress))

	cfg, msgs, err := s.convertRevokeAllSessionsRequest(req)
	if err != nil {
		return err
	}

	creator, err := s.getSigner(req.CreatorAddress)
	if err != nil {
		return err
	}

	c, err := s.getClient(creator)
	if err != nil {
		return getGrpcError(err)
	}
	bres, err := c.RevokeAllSessions(*cfg, msgs...)
	if err != nil {
		return getGrpcError(err)
	}

	if err := send(&api_gen.RevokeAllSessionsResponse{
		Result: bres.DeliverTx.Data,
		Hash:   bres.Hash,
		Height: bres.Height,
	}); err != nil {
		return err
	}

	return nil
}

func (s *gnoNativeService) MakeRevokeAllSessionsTx(ctx context.Context, req *api_gen.RevokeAllSessionsRequest) (*api_gen.MakeTxResponse, error) {
	s.logger.Debug("MakeRevokeAllSessionsTx", zap.ByteString("creator", req.CreatorAddress))

	cfg, msgs, err := s.convertRevokeAllSessionsRequest(req)
	if err != nil {
		return nil, err
	}
	tx, err := gnoclient.NewRevokeAllSessionsTx(*cfg, msgs...)
	if err != nil {
		return nil, getGrpcError(err)
	}

	txJSON, err := amino.MarshalJSON(tx)
	if err != nil {
		return nil, err
	}
	return &api_gen.MakeTxResponse{TxJSON: string(txJSON)}, nil
}

func (s *gnoNativeService) convertRevokeAllSessionsRequest(req *api_gen.RevokeAllSessionsRequest) (*gnoclient.BaseTxCfg, []auth.MsgRevokeAllSessions, error) {
	creator, err := crypto.AddressFromBytes(req.CreatorAddress)
	if err != nil {
		return nil, nil, getGrpcError(err)
	}

	cfg := &gnoclient.BaseTxCfg{
		GasFee:    req.GasFee,
		GasWanted: req.GasWanted,
		Memo:      req.Memo,
	}

	msgs := []auth.MsgRevokeAllSessions{{Creator: creator}}
	return cfg, msgs, nil
}

func (s *gnoNativeService) SignTx(ctx context.Context, req *api_gen.SignTxRequest) (*api_gen.SignTxResponse, error) {
	var tx std.Tx
	if err := amino.UnmarshalJSON([]byte(req.TxJSON), &tx); err != nil {
		return nil, err
	}

	signedTx, err := s.ClientSignTx(tx, req.Address, req.AccountNumber, req.SequenceNumber)
	if err != nil {
		return nil, getGrpcError(err)
	}

	signedTxJSON, err := amino.MarshalJSON(signedTx)
	if err != nil {
		return nil, err
	}
	return &api_gen.SignTxResponse{SignedTxJSON: string(signedTxJSON)}, nil
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

func (s *gnoNativeService) EstimateGas(ctx context.Context, req *api_gen.EstimateGasRequest) (*api_gen.EstimateGasResponse, error) {
	var tx std.Tx
	if err := amino.UnmarshalJSON([]byte(req.TxJSON), &tx); err != nil {
		return nil, err
	}

	gasWanted, _, err := s.estimateGasWanted(&tx, req.Address, req.SecurityMargin, req.UpdateTx)
	if err != nil {
		return nil, getGrpcError(err)
	}

	txJSON, err := amino.MarshalJSON(tx)
	if err != nil {
		return nil, err
	}

	return &api_gen.EstimateGasResponse{TxJSON: string(txJSON), GasWanted: gasWanted}, nil
}

func (s *gnoNativeService) EstimateTxFees(ctx context.Context, req *api_gen.EstimateTxFeesRequest) (*api_gen.EstimateTxFeesResponse, error) {
	var tx std.Tx
	if err := amino.UnmarshalJSON([]byte(req.TxJSON), &tx); err != nil {
		return nil, err
	}

	gasWanted, deliverTx, err := s.estimateGasWanted(&tx, req.Address, req.GasSecurityMargin, req.UpdateTx)
	if err != nil {
		return nil, getGrpcError(err)
	}

	c, err := s.getClient(nil)
	if err != nil {
		return nil, getGrpcError(err)
	}

	// Imitate gnokey CLI https://github.com/gnolang/gno/blob/de4b5b56c60126373ec0702234c196fdae365a0b/tm2/pkg/crypto/keys/client/broadcast.go#L142
	qres, err := c.Query(gnoclient.QueryCfg{Path: "auth/gasprice", Data: []byte{}})
	if err != nil {
		return nil, getGrpcError(err)
	}
	gp := std.GasPrice{}
	err = amino.UnmarshalJSON(qres.Response.Data, &gp)
	if err != nil {
		return nil, getGrpcError(err)
	}
	if gp.Gas == 0 {
		// Can't get the gas price from the node
		return nil, api_gen.ErrCode_ErrInvalidCoins
	}
	if gp.Price.Denom != ugnot.Denom {
		return nil, api_gen.ErrCode_ErrInvalidCoins
	}
	fee := gasWanted/gp.Gas + 1
	fee = overflow.Mulp(fee, gp.Price.Amount)
	// fee buffer to cover the sudden change of gas price
	feeBuffer := float64(req.GasPriceSecurityMargin) / 100
	fee = int64(float64(fee) * feeBuffer)

	totalFee := fee
	if req.UpdateTx {
		tx.Fee.GasFee = std.NewCoin(gp.Price.Denom, fee)
	}

	txJSON, err := amino.MarshalJSON(tx)
	if err != nil {
		return nil, err
	}

	response := &api_gen.EstimateTxFeesResponse{
		TxJSON:    string(txJSON),
		GasWanted: gasWanted,
		GasFee: &api_gen.Coin{
			Denom:  gp.Price.Denom,
			Amount: fee,
		},
	}
	if delta, storageFee, ok := keyscli.GetStorageInfo(deliverTx.Events); ok {
		response.StorageDelta = delta

		storageCoins := make([]*api_gen.Coin, 0)
		for _, coin := range storageFee {
			storageCoins = append(storageCoins, &api_gen.Coin{
				Denom:  coin.Denom,
				Amount: coin.Amount,
			})
			if coin.Denom == ugnot.Denom {
				totalFee += coin.Amount
			}
		}
		response.StorageFee = storageCoins
	}

	response.TotalFee = &api_gen.Coin{
		Denom:  ugnot.Denom,
		Amount: totalFee,
	}
	return response, nil
}

// estimateGasWanted is a helper for EstimateGas, etc. Use the tx and address to call gnoclient.Simulate, then
// multiply the GasUsed by securityMarginPercent/100 and return the gas wanted. If updateTx is true, then update tx.Fee.GasWanted .
func (s *gnoNativeService) estimateGasWanted(tx *std.Tx, address []byte, securityMarginPercent uint32, updateTx bool) (int64, *abci.ResponseDeliverTx, error) {
	signer, err := s.getSigner(address)
	if err != nil {
		return 0, nil, err
	}
	info, err := signer.Info()
	if err != nil {
		return 0, nil, err
	}

	// Set the tx signature using the public key. No need to sign to get the actual signature bytes.
	tx.Signatures = []std.Signature{
		{PubKey: info.GetPubKey()},
	}

	c, err := s.getClient(nil)
	if err != nil {
		return 0, nil, getGrpcError(err)
	}

	deliverTx, err := c.Simulate(tx)
	if err != nil {
		return 0, nil, getGrpcError(err)
	}

	// Apply the security margin.
	// The security margin is a decimal numeral without the decimal seprator.
	gasWanted := int64(float64(deliverTx.GasUsed) * float64(securityMarginPercent) / 100)

	// Update the transaction
	if updateTx {
		tx.Fee.GasWanted = gasWanted
	}

	return gasWanted, deliverTx, nil
}

func (s *gnoNativeService) BroadcastTxCommit(ctx context.Context, req *api_gen.BroadcastTxCommitRequest, send func(*api_gen.BroadcastTxCommitResponse) error) error {
	signedTx := &std.Tx{}
	if err := amino.UnmarshalJSON([]byte(req.SignedTxJSON), signedTx); err != nil {
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

	if err := send(&api_gen.BroadcastTxCommitResponse{
		Result: bres.DeliverTx.Data,
		Hash:   bres.Hash,
		Height: bres.Height,
	}); err != nil {
		s.logger.Error("BroadcastTxCommit send returned error", zap.Error(err))
		return err
	}

	return nil
}

func (s *gnoNativeService) AddressToBech32(ctx context.Context, req *api_gen.AddressToBech32Request) (*api_gen.AddressToBech32Response, error) {
	s.logger.Debug("AddressToBech32", zap.ByteString("address", req.Address))
	bech32Address, err := bech32.Encode(crypto.Bech32AddrPrefix(), req.Address)
	if err != nil {
		return nil, getGrpcError(err)
	}
	return &api_gen.AddressToBech32Response{Bech32Address: bech32Address}, nil
}

func (s *gnoNativeService) AddressFromBech32(ctx context.Context, req *api_gen.AddressFromBech32Request) (*api_gen.AddressFromBech32Response, error) {
	address, err := crypto.AddressFromBech32(req.Bech32Address)
	if err != nil {
		return nil, err
	}

	return &api_gen.AddressFromBech32Response{Address: address.Bytes()}, nil
}

func (s *gnoNativeService) AddressFromMnemonic(ctx context.Context, req *api_gen.AddressFromMnemonicRequest) (*api_gen.AddressFromMnemonicResponse, error) {
	kb := crypto_keys.NewInMemory()
	info, err := kb.CreateAccount("temporary", req.Mnemonic, "", "", uint32(0), uint32(0))
	if err != nil {
		return nil, err
	}

	return &api_gen.AddressFromMnemonicResponse{Address: info.GetAddress().Bytes()}, nil
}

func (s *gnoNativeService) ValidateMnemonicWord(ctx context.Context, req *api_gen.ValidateMnemonicWordRequest) (*api_gen.ValidateMnemonicWordResponse, error) {
	valid := slices.Contains(bip39.EnglishWordList, req.Word)
	return &api_gen.ValidateMnemonicWordResponse{Valid: valid}, nil
}

func (s *gnoNativeService) ValidateMnemonicPhrase(ctx context.Context, req *api_gen.ValidateMnemonicPhraseRequest) (*api_gen.ValidateMnemonicPhraseResponse, error) {
	_, err := bip39.MnemonicToByteArray(req.Phrase)
	return &api_gen.ValidateMnemonicPhraseResponse{Valid: err == nil}, nil
}

func (s *gnoNativeService) PubKeyBytesFromBech32(ctx context.Context, req *api_gen.PubKeyBytesFromBech32Request) (*api_gen.PubKeyBytesFromBech32Response, error) {
	pubKey, err := crypto.PubKeyFromBech32(req.Bech32PubKey)
	if err != nil {
		return nil, err
	}

	return &api_gen.PubKeyBytesFromBech32Response{PubKeyBytes: pubKey.Bytes()}, nil
}

func (s *gnoNativeService) Hello(ctx context.Context, req *api_gen.HelloRequest) (*api_gen.HelloResponse, error) {
	s.logger.Debug("Hello called")
	defer s.logger.Debug("Hello returned ok")
	return &api_gen.HelloResponse{
		Greeting: "Hello " + req.Name,
	}, nil
}

// HelloStream is for debug purposes
func (s *gnoNativeService) HelloStream(ctx context.Context, req *api_gen.HelloStreamRequest, send func(*api_gen.HelloStreamResponse) error) error {
	s.logger.Debug("HelloStream called")
	for i := 0; i < 4; i++ {
		if err := send(&api_gen.HelloStreamResponse{
			Greeting: "Hello " + req.Name,
		}); err != nil {
			s.logger.Error("HelloStream returned error", zap.Error(err))
			return err
		}
		// Respect ctx cancellation so a closed direct stream doesn't leak this goroutine.
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(2 * time.Second):
		}
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

// Temporary: Remove after merging https://github.com/gnolang/gno/pull/4630
type StorageDepositEvent struct {
	// "StorageDeposit" or "UnlockDeposit"
	Type       string `json:"type"`
	BytesDelta int64  `json:"bytes_delta"`
	FeeDelta   string `json:"fee_delta"`
}

func (e StorageDepositEvent) AssertABCIEvent() {}
