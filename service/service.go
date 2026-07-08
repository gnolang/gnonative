package service

import (
	"io"
	"sync"

	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/tm2/pkg/bech32"
	rpcclient "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	"github.com/gnolang/gno/tm2/pkg/crypto"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys"
	crypto_keys "github.com/gnolang/gno/tm2/pkg/crypto/keys"
	api_gen "github.com/gnolang/gnonative/v5/api"
	"go.uber.org/zap"
)

// GnoNativeService is the exported service surface. It is the plain-Go
// (connect/gRPC-free) API plus Close(). The mobile bridge dispatcher and pure-Go
// consumers both call these methods directly.
type GnoNativeService interface {
	GnoNativeApi

	io.Closer
}

type userAccount struct {
	keyInfo keys.Info
	signer  *gnoclient.SignerFromKeybase
}

type gnoNativeService struct {
	logger    *zap.Logger
	keybase   crypto_keys.Keybase
	rpcClient *rpcclient.RPCClient
	lock      sync.RWMutex
	// The remote node address used to create client.RPCClient. We need to save this
	// here because the remote is a private member of the HTTP struct.
	remote string
	// TODO: Allow each userAccount to have its own chain ID
	chainID string

	// Map of key bech32 to userAccount.
	userAccounts map[string]*userAccount

	closeFunc func()
}

var _ GnoNativeService = (*gnoNativeService)(nil)

// NewGnoNativeService creates a new GnoNative service. There is no server: callers
// invoke the API methods directly (or via the mobile bridge dispatcher).
func NewGnoNativeService(opts ...GnoNativeOption) (GnoNativeService, error) {
	cfg := &Config{}
	if err := cfg.applyOptions(append(opts, WithFallbackDefaults)...); err != nil {
		return nil, err
	}

	svc, err := initService(cfg)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func initService(cfg *Config) (*gnoNativeService, error) {
	svc := &gnoNativeService{
		logger:       cfg.Logger,
		userAccounts: make(map[string]*userAccount),
		closeFunc:    func() {},
	}

	if err := cfg.checkDirs(); err != nil {
		return nil, err
	}

	if cfg.NativeDB != nil {
		cfg.Logger.Debug("using nativeDB for keybase")
		svc.keybase = keys.NewDBKeybase(cfg.NativeDB)
	} else {
		var err error
		cfg.Logger.Debug("using filesystem for keybase", zap.String("rootdir", cfg.RootDir))
		svc.keybase, err = keys.NewKeyBaseFromDir(cfg.RootDir)
		if err != nil {
			return nil, err
		}
	}

	var err error
	svc.rpcClient, err = rpcclient.NewHTTPClient(cfg.Remote)
	if err != nil {
		return nil, err
	}
	svc.remote = cfg.Remote
	svc.chainID = cfg.ChainID

	return svc, nil
}

// Get a gnoclient.Client with the RPCClient and the given signer.
func (s *gnoNativeService) getClient(signer gnoclient.Signer) (*gnoclient.Client, error) {
	return &gnoclient.Client{
		Signer:    signer,
		RPCClient: s.rpcClient,
	}, nil
}

// Look up addr in s.userAccounts and return the signer.
// (Also set the signer.ChainID to s.chainID. This may change if we allow each userAccount to have its own chain ID.)
// If there is no active account with the given address, return ErrCode_ErrNoActiveAccount.
func (s *gnoNativeService) getSigner(addr []byte) (*gnoclient.SignerFromKeybase, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	b32, err := bech32.Encode(crypto.Bech32AddrPrefix(), addr)
	if err != nil {
		return nil, getGrpcError(err)
	}
	account, ok := s.userAccounts[b32]
	if !ok {
		return nil, api_gen.ErrCode_ErrNoActiveAccount
	}

	account.signer.ChainID = s.chainID
	return account.signer, nil
}

func (s *gnoNativeService) Close() error {
	if s.closeFunc != nil {
		s.closeFunc()
	}
	return nil
}
