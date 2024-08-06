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
	"github.com/gnolang/gno/tm2/pkg/std"
	"go.uber.org/zap"

	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/gno.land/pkg/sdk/vm"
	rpcclient "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	api_gen "github.com/gnolang/gnonative/api/gen/go"
)

func (s *gnoNativeService) SetRemote(ctx context.Context, req *connect.Request[api_gen.SetRemoteRequest]) (*connect.Response[api_gen.SetRemoteResponse], error) {
	var err error
	s.client.RPCClient, err = rpcclient.NewHTTPClient(req.Msg.Remote)
	if err != nil {
		return nil, api_gen.ErrCode_ErrSetRemote.Wrap(err)
	}
	s.remote = req.Msg.Remote
	return connect.NewResponse(&api_gen.SetRemoteResponse{}), nil
}

func (s *gnoNativeService) GetRemote(ctx context.Context, req *connect.Request[api_gen.GetRemoteRequest]) (*connect.Response[api_gen.GetRemoteResponse], error) {
	return connect.NewResponse(&api_gen.GetRemoteResponse{Remote: s.remote}), nil
}

func (s *gnoNativeService) SetChainID(ctx context.Context, req *connect.Request[api_gen.SetChainIDRequest]) (*connect.Response[api_gen.SetChainIDResponse], error) {
	s.getSigner().ChainID = req.Msg.ChainId
	return connect.NewResponse(&api_gen.SetChainIDResponse{}), nil
}

func (s *gnoNativeService) GetChainID(ctx context.Context, req *connect.Request[api_gen.GetChainIDRequest]) (*connect.Response[api_gen.GetChainIDResponse], error) {
	return connect.NewResponse(&api_gen.GetChainIDResponse{ChainId: s.getSigner().ChainID}), nil
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

	keys, err := s.getSigner().Keybase.List()
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

func (s *gnoNativeService) HasKeyByName(ctx context.Context, req *connect.Request[api_gen.HasKeyByNameRequest]) (*connect.Response[api_gen.HasKeyByNameResponse], error) {
	s.logger.Debug("HasKeyByName called")

	has, err := s.getSigner().Keybase.HasByName(req.Msg.Name)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return connect.NewResponse(&api_gen.HasKeyByNameResponse{Has: has}), nil
}

func (s *gnoNativeService) HasKeyByAddress(ctx context.Context, req *connect.Request[api_gen.HasKeyByAddressRequest]) (*connect.Response[api_gen.HasKeyByAddressResponse], error) {
	s.logger.Debug("HasKeyByAddress called")

	has, err := s.getSigner().Keybase.HasByAddress(crypto.AddressFromBytes(req.Msg.Address))
	if err != nil {
		return nil, getGrpcError(err)
	}

	return connect.NewResponse(&api_gen.HasKeyByAddressResponse{Has: has}), nil
}

func (s *gnoNativeService) HasKeyByNameOrAddress(ctx context.Context, req *connect.Request[api_gen.HasKeyByNameOrAddressRequest]) (*connect.Response[api_gen.HasKeyByNameOrAddressResponse], error) {
	s.logger.Debug("HasKeyByNameOrAddress called")

	has, err := s.getSigner().Keybase.HasByNameOrAddress(req.Msg.NameOrBech32)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return connect.NewResponse(&api_gen.HasKeyByNameOrAddressResponse{Has: has}), nil
}

func (s *gnoNativeService) GetKeyInfoByName(ctx context.Context, req *connect.Request[api_gen.GetKeyInfoByNameRequest]) (*connect.Response[api_gen.GetKeyInfoByNameResponse], error) {
	s.logger.Debug("GetKeyInfoByName called")

	key, err := s.getSigner().Keybase.GetByName(req.Msg.Name)
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

	key, err := s.getSigner().Keybase.GetByAddress(crypto.AddressFromBytes(req.Msg.Address))
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

	key, err := s.getSigner().Keybase.GetByNameOrAddress(req.Msg.NameOrBech32)
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

	key, err := s.getSigner().Keybase.CreateAccount(req.Msg.NameOrBech32, req.Msg.Mnemonic, req.Msg.Bip39Passwd, req.Msg.Password, req.Msg.Account, req.Msg.Index)
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
	s.logger.Debug("SelectAccount called", zap.String("NameOrBech32", req.Msg.NameOrBech32))

	// The key may already be in s.userAccounts, but the info may have changed on disk. So always get from disk.
	key, err := s.getSigner().Keybase.GetByNameOrAddress(req.Msg.NameOrBech32)
	if err != nil {
		return nil, getGrpcError(err)
	}

	info, err := ConvertKeyInfo(key)
	if err != nil {
		return nil, err
	}

	s.lock.Lock()
	account, ok := s.userAccounts[req.Msg.NameOrBech32]
	if !ok {
		account = &userAccount{}
		s.userAccounts[req.Msg.NameOrBech32] = account
	}
	account.keyInfo = key
	s.activeAccount = account
	s.lock.Unlock()

	s.getSigner().Account = req.Msg.NameOrBech32
	s.getSigner().Password = account.password
	return connect.NewResponse(&api_gen.SelectAccountResponse{
		Key:         info,
		HasPassword: account.password != "",
	}), nil
}

func (s *gnoNativeService) SetPassword(ctx context.Context, req *connect.Request[api_gen.SetPasswordRequest]) (*connect.Response[api_gen.SetPasswordResponse], error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.activeAccount == nil {
		return nil, api_gen.ErrCode_ErrNoActiveAccount
	}
	s.activeAccount.password = req.Msg.Password

	s.getSigner().Password = req.Msg.Password

	// Check the password.
	if err := s.getSigner().Validate(); err != nil {
		return nil, getGrpcError(err)
	}

	return connect.NewResponse(&api_gen.SetPasswordResponse{}), nil
}

func (s *gnoNativeService) GetActiveAccount(ctx context.Context, req *connect.Request[api_gen.GetActiveAccountRequest]) (*connect.Response[api_gen.GetActiveAccountResponse], error) {
	s.logger.Debug("GetActiveAccount called")

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
		HasPassword: account.password != "",
	}), nil
}

func (s *gnoNativeService) QueryAccount(ctx context.Context, req *connect.Request[api_gen.QueryAccountRequest]) (*connect.Response[api_gen.QueryAccountResponse], error) {
	s.logger.Debug("QueryAccount", zap.ByteString("address", req.Msg.Address))

	// gnoclient wants the bech32 address.
	account, _, err := s.client.QueryAccount(crypto.AddressFromBytes(req.Msg.Address))
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
	if err := s.getSigner().Keybase.Delete(req.Msg.NameOrBech32, req.Msg.Password, req.Msg.SkipPassword); err != nil {
		return nil, getGrpcError(err)
	}

	s.lock.Lock()
	delete(s.userAccounts, req.Msg.NameOrBech32)
	if s.activeAccount != nil &&
		(s.activeAccount.keyInfo.GetName() == req.Msg.NameOrBech32 || crypto.AddressToBech32(s.activeAccount.keyInfo.GetAddress()) == req.Msg.NameOrBech32) {
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

	bres, err := s.client.Query(cfg)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return connect.NewResponse(&api_gen.QueryResponse{Result: bres.Response.Data}), nil
}

func (s *gnoNativeService) Render(ctx context.Context, req *connect.Request[api_gen.RenderRequest]) (*connect.Response[api_gen.RenderResponse], error) {
	s.logger.Debug("Render", zap.String("packagePath", req.Msg.PackagePath), zap.String("args", req.Msg.Args))

	result, _, err := s.client.Render(req.Msg.PackagePath, req.Msg.Args)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return connect.NewResponse(&api_gen.RenderResponse{Result: result}), nil
}

func (s *gnoNativeService) QEval(ctx context.Context, req *connect.Request[api_gen.QEvalRequest]) (*connect.Response[api_gen.QEvalResponse], error) {
	s.logger.Debug("QEval", zap.String("packagePath", req.Msg.PackagePath), zap.String("expression", req.Msg.Expression))

	result, _, err := s.client.QEval(req.Msg.PackagePath, req.Msg.Expression)
	if err != nil {
		return nil, getGrpcError(err)
	}

	return connect.NewResponse(&api_gen.QEvalResponse{Result: result}), nil
}

func (s *gnoNativeService) Call(ctx context.Context, req *connect.Request[api_gen.CallRequest], stream *connect.ServerStream[api_gen.CallResponse]) error {
	for _, msg := range req.Msg.Msgs {
		s.logger.Debug("Call", zap.String("package", msg.PackagePath), zap.String("function", msg.Fnc), zap.Any("args", msg.Args))
	}

	s.lock.RLock()
	if s.activeAccount == nil {
		s.lock.RUnlock()
		return api_gen.ErrCode_ErrNoActiveAccount
	}
	s.lock.RUnlock()

	cfg, msgs := convertCallRequest(req.Msg)

	bres, err := s.client.Call(*cfg, msgs...)
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

func convertCallRequest(req *api_gen.CallRequest) (*gnoclient.BaseTxCfg, []gnoclient.MsgCall) {
	var callerAddress crypto.Address
	if req.CallerAddress != nil {
		callerAddress = crypto.AddressFromBytes(req.CallerAddress)
	}
	cfg := &gnoclient.BaseTxCfg{
		GasFee:        req.GasFee,
		GasWanted:     req.GasWanted,
		Memo:          req.Memo,
		CallerAddress: callerAddress,
	}

	msgs := make([]gnoclient.MsgCall, 0)

	for _, msg := range req.Msgs {
		msgs = append(msgs, gnoclient.MsgCall{
			PkgPath:  msg.PackagePath,
			FuncName: msg.Fnc,
			Args:     msg.Args,
			Send:     msg.Send,
		})
	}

	return cfg, msgs
}

func (s *gnoNativeService) Send(ctx context.Context, req *connect.Request[api_gen.SendRequest], stream *connect.ServerStream[api_gen.SendResponse]) error {
	for _, msg := range req.Msg.Msgs {
		s.logger.Debug("Send", zap.String("toAddress", crypto.AddressToBech32(crypto.AddressFromBytes(msg.ToAddress))), zap.String("send", msg.Send))
	}

	s.lock.RLock()
	if s.activeAccount == nil {
		s.lock.RUnlock()
		return api_gen.ErrCode_ErrNoActiveAccount
	}
	s.lock.RUnlock()

	cfg, msgs := convertSendRequest(req.Msg)

	_, err := s.client.Send(*cfg, msgs...)
	if err != nil {
		return getGrpcError(err)
	}

	if err := stream.Send(&api_gen.SendResponse{}); err != nil {
		s.logger.Error("Send stream.Send returned error", zap.Error(err))
		return err
	}

	return nil
}

func convertSendRequest(req *api_gen.SendRequest) (*gnoclient.BaseTxCfg, []gnoclient.MsgSend) {
	var callerAddress crypto.Address
	if req.CallerAddress != nil {
		callerAddress = crypto.AddressFromBytes(req.CallerAddress)
	}
	cfg := &gnoclient.BaseTxCfg{
		GasFee:        req.GasFee,
		GasWanted:     req.GasWanted,
		Memo:          req.Memo,
		CallerAddress: callerAddress,
	}

	msgs := make([]gnoclient.MsgSend, 0)

	for _, msg := range req.Msgs {
		msgs = append(msgs, gnoclient.MsgSend{
			ToAddress: crypto.AddressFromBytes(msg.ToAddress),
			Send:      msg.Send,
		})
	}

	return cfg, msgs
}

func (s *gnoNativeService) Run(ctx context.Context, req *connect.Request[api_gen.RunRequest], stream *connect.ServerStream[api_gen.RunResponse]) error {
	s.lock.RLock()
	if s.activeAccount == nil {
		s.lock.RUnlock()
		return api_gen.ErrCode_ErrNoActiveAccount
	}
	s.lock.RUnlock()

	cfg, msgs := convertRunRequest(req.Msg)

	bres, err := s.client.Run(*cfg, msgs...)
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

func convertRunRequest(req *api_gen.RunRequest) (*gnoclient.BaseTxCfg, []gnoclient.MsgRun) {
	var callerAddress crypto.Address
	if req.CallerAddress != nil {
		callerAddress = crypto.AddressFromBytes(req.CallerAddress)
	}
	cfg := &gnoclient.BaseTxCfg{
		GasFee:        req.GasFee,
		GasWanted:     req.GasWanted,
		Memo:          req.Memo,
		CallerAddress: callerAddress,
	}

	msgs := make([]gnoclient.MsgRun, 0)

	for _, msg := range req.Msgs {
		memPkg := &std.MemPackage{
			Files: []*std.MemFile{
				{
					Name: "main.gno",
					Body: msg.Package,
				},
			},
		}
		msgs = append(msgs, gnoclient.MsgRun{
			Package: memPkg,
			Send:    msg.Send,
		})
	}

	return cfg, msgs
}

func (s *gnoNativeService) MakeCallTx(ctx context.Context, req *connect.Request[api_gen.CallRequest]) (*connect.Response[api_gen.MakeTxResponse], error) {
	for _, msg := range req.Msg.Msgs {
		s.logger.Debug("MakeCallTx", zap.String("package", msg.PackagePath), zap.String("function", msg.Fnc), zap.Any("args", msg.Args))
	}

	cfg, msgs := convertCallRequest(req.Msg)
	tx, err := s.client.MakeCallTx(*cfg, msgs...)
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
	cfg, msgs := convertSendRequest(req.Msg)
	tx, err := s.client.MakeSendTx(*cfg, msgs...)
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
	cfg, msgs := convertRunRequest(req.Msg)
	tx, err := s.client.MakeRunTx(*cfg, msgs...)
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
	s.lock.RLock()
	if s.activeAccount == nil {
		s.lock.RUnlock()
		return nil, api_gen.ErrCode_ErrNoActiveAccount
	}
	s.lock.RUnlock()

	var tx std.Tx
	if err := amino.UnmarshalJSON([]byte(req.Msg.TxJson), &tx); err != nil {
		return nil, err
	}

	signedTx, err := s.client.SignTx(tx, req.Msg.AccountNumber, req.Msg.SequenceNumber)
	if err != nil {
		return nil, getGrpcError(err)
	}

	signedTxJSON, err := amino.MarshalJSON(signedTx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&api_gen.SignTxResponse{SignedTxJson: string(signedTxJSON)}), nil
}

func (s *gnoNativeService) BroadcastTxCommit(ctx context.Context, req *connect.Request[api_gen.BroadcastTxCommitRequest],
	stream *connect.ServerStream[api_gen.BroadcastTxCommitResponse]) error {
	signedTx := &std.Tx{}
	if err := amino.UnmarshalJSON([]byte(req.Msg.SignedTxJson), signedTx); err != nil {
		return err
	}

	bres, err := s.client.BroadcastTxCommit(signedTx)
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
