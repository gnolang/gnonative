package service

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	rpcclient "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys"
	"github.com/gnolang/gnomobile/gnoclient"
	"github.com/gnolang/gnomobile/service/rpc"
	"github.com/gnolang/gnomobile/service/rpc/rpcconnect"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const SOCKET_SUBDIR = "s"
const SOCKET_FILE = "gno"

type GnomobileService interface {
	GetSocketPath() string
	GetTcpPort() int

	io.Closer
}

type gnomobileService struct {
	logger     *zap.Logger
	client     *gnoclient.Client
	tcpPort    int
	socketPath string
	lock       sync.RWMutex

	activeAccount keys.Info

	listener net.Listener
	server   *http.Server
}

var _ GnomobileService = (*gnomobileService)(nil)

func NewGnomobileService(opts ...GnomobileOption) (GnomobileService, error) {
	cfg := &Config{}
	if err := cfg.applyOptions(append(opts, WithFallbackDefaults)...); err != nil {
		return nil, err
	}

	svc, err := initService(cfg)
	if err != nil {
		return nil, err
	}

	if cfg.UseTcpListener {
		if err := svc.createTcpListener(); err != nil {
			return nil, err
		}
	} else {
		if err := svc.createUDSListener(cfg); err != nil {
			return nil, err
		}
	}

	if err := svc.runGRPCServer(); err != nil {
		return nil, err
	}

	return svc, nil
}

func initService(cfg *Config) (*gnomobileService, error) {
	svc := &gnomobileService{
		logger:  cfg.Logger,
		tcpPort: cfg.TcpPort,
	}

	if err := cfg.checkDirs(); err != nil {
		return nil, err
	}

	kb, _ := keys.NewKeyBaseFromDir(cfg.RootDir)
	signer := &gnoclient.SignerFromKeybase{
		Keybase: kb,
		ChainID: cfg.ChainID,
	}

	rpcClient := rpcclient.NewHTTP(cfg.Remote, "/websocket")

	svc.client = &gnoclient.Client{
		Signer:    signer,
		RPCClient: rpcClient,
	}

	return svc, nil
}

func (cfg *Config) applyOptions(opts ...GnomobileOption) error {
	withDefaultOpts := make([]GnomobileOption, len(opts))
	copy(withDefaultOpts, opts)
	withDefaultOpts = append(withDefaultOpts, WithFallbackDefaults)
	for _, opt := range withDefaultOpts {
		if err := opt(cfg); err != nil {
			return err
		}
	}
	return nil
}

// Get s.client.Signer as a SignerFromKeybase.
func (s *gnomobileService) getSigner() *gnoclient.SignerFromKeybase {
	signer, ok := s.client.Signer.(*gnoclient.SignerFromKeybase)
	if !ok {
		// We only set s.client.Signer in initService, so this shouldn't happen.
		panic("signer is not gnoclient.SignerFromKeybase")
	}
	return signer
}

func (cfg *Config) checkDirs() error {
	// check if rootDir exists
	{
		_, err := os.Stat(cfg.RootDir)
		if os.IsNotExist(err) {
			return errors.Wrap(err, "rootDir folder doesn't exist")
		}
	}

	// check if tmpDir exists
	{
		_, err := os.Stat(cfg.TmpDir)
		if os.IsNotExist(err) {
			return errors.Wrap(err, "tmpDir folder doesn't exist")
		}
	}

	return nil
}

func (s *gnomobileService) createUDSListener(cfg *Config) error {
	s.logger.Debug("createUDSListener called")

	// create a socket subdirectory
	sockDir := filepath.Join(cfg.TmpDir, SOCKET_SUBDIR)
	if err := os.MkdirAll(sockDir, 0700); err != nil {
		s.logger.Debug("createUDSListener error", zap.Error(err))
		return rpc.ErrCode_ErrRunGRPCServer.Wrap(err)
	}

	s.socketPath = filepath.Join(sockDir, SOCKET_FILE)

	// delete socket if it already exists
	if _, err := os.Stat(s.socketPath); !os.IsNotExist(err) {
		if err := os.RemoveAll(s.socketPath); err != nil {
			s.logger.Debug("createUDSListener error", zap.Error(err))
			return rpc.ErrCode_ErrRunGRPCServer.Wrap(err)
		}
	}

	listener, err := net.Listen("unix", s.socketPath)
	if err != nil {
		s.logger.Debug("createUDSListener error", zap.Error(err))
		return rpc.ErrCode_ErrRunGRPCServer.Wrap(err)
	}

	s.listener = listener
	s.logger.Info("createUDSListener", zap.String("path", s.socketPath))

	return nil
}

func (s *gnomobileService) createTcpListener() error {
	s.logger.Debug("createTcpListener called")

	tcpAddr := fmt.Sprintf(":%d", s.tcpPort)
	listener, err := net.Listen("tcp", tcpAddr)
	if err != nil {
		s.logger.Debug("createTcpListener error", zap.Error(err))
		return rpc.ErrCode_ErrRunGRPCServer.Wrap(err)
	}

	s.listener = listener

	// update the tcpPort field

	addr := listener.Addr().String()

	_, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		s.logger.Debug("createTcpListener error", zap.Error(err))
		return rpc.ErrCode_ErrRunGRPCServer.Wrap(err)
	}

	portInt, err := net.LookupPort("tcp", portStr)
	if err != nil {
		s.logger.Debug("createTcpListener error", zap.Error(err))
		return rpc.ErrCode_ErrRunGRPCServer.Wrap(err)
	}

	s.logger.Info("createTcpListener: gRPC server listen on", zap.Int("port", portInt))
	s.tcpPort = portInt

	return nil
}

func newCORS() *cors.Cors {
	// To let web developers play with the demo service from browsers, we need a
	// very permissive CORS setup.
	return cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowOriginFunc: func(origin string) bool {
			// Allow all origins, which effectively disables CORS.
			return true
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{
			// Content-Type is in the default safelist.
			"Accept",
			"Accept-Encoding",
			"Accept-Post",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Content-Encoding",
			"Grpc-Accept-Encoding",
			"Grpc-Encoding",
			"Grpc-Message",
			"Grpc-Status",
			"Grpc-Status-Details-Bin",
		},
		// Let browsers cache CORS information for longer, which reduces the number
		// of preflight requests. Any changes to ExposedHeaders won't take effect
		// until the cached data expires. FF caps this value at 24h, and modern
		// Chrome caps it at 2h.
		MaxAge: int(2 * time.Hour / time.Second),
	})
}

func (s *gnomobileService) runGRPCServer() error {
	if s.listener == nil {
		return rpc.ErrCode_ErrRunGRPCServer.Wrap(errors.New("listener is not initialized"))
	}

	mux := http.NewServeMux()

	compress1KB := connect.WithCompressMinBytes(1024)
	mux.Handle(rpcconnect.NewGnomobileServiceHandler(
		s,
		compress1KB,
	))
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(rpcconnect.GnomobileServiceName),
		compress1KB,
	))
	mux.Handle(grpcreflect.NewHandlerV1(
		grpcreflect.NewStaticReflector(rpcconnect.GnomobileServiceName),
		compress1KB,
	))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(
		grpcreflect.NewStaticReflector(rpcconnect.GnomobileServiceName),
		compress1KB,
	))

	server := &http.Server{
		Handler: h2c.NewHandler(
			newCORS().Handler(mux),
			&http2.Server{},
		),
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}

	go func() {
		// we dont need to log the error
		err := s.server.Serve(s.listener)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("failed to serve the gRPC listener")
		}
	}()

	s.lock.Lock()
	s.server = server
	s.lock.Unlock()

	return nil
}

func (s *gnomobileService) GetSocketPath() string {
	return s.socketPath
}

func (s *gnomobileService) GetTcpPort() int {
	return s.tcpPort
}

func (s *gnomobileService) Close() error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.server != nil {
		if err := s.server.Shutdown(context.Background()); err != nil {
			s.logger.Error("HTTP shutdown: %v", zap.Error(err)) //nolint:gocritic
		}
		s.server = nil
	}

	return nil
}

// Config describes a set of settings for a GnomobileService
type Config struct {
	Logger         *zap.Logger
	Remote         string
	ChainID        string
	RootDir        string
	TmpDir         string
	TcpPort        int
	UseTcpListener bool
}

type GnomobileOption func(cfg *Config) error

// FallBackOption is a structure that permits to fallback to a default option if the option is not set.
type FallBackOption struct {
	fallback func(cfg *Config) bool
	opt      GnomobileOption
}

// --- Logger options ---

// WithLogger set the given logger.
var WithLogger = func(l *zap.Logger) GnomobileOption {
	return func(cfg *Config) error {
		cfg.Logger = l
		return nil
	}
}

// WithDefaultLogger init a noop logger.
var WithDefaultLogger GnomobileOption = func(cfg *Config) error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}

	cfg.Logger = logger

	return nil
}

var fallbackLogger = FallBackOption{
	fallback: func(cfg *Config) bool { return cfg.Logger == nil },
	opt:      WithDefaultLogger,
}

// WithFallbackLogger sets the logger if no logger is set.
var WithFallbackLogger GnomobileOption = func(cfg *Config) error {
	if fallbackLogger.fallback(cfg) {
		return fallbackLogger.opt(cfg)
	}
	return nil
}

// --- Remote options ---

// WithRemote sets the given remote node address.
var WithRemote = func(remote string) GnomobileOption {
	return func(cfg *Config) error {
		cfg.Remote = remote
		return nil
	}
}

// WithDefaultRemote inits a default remote node address.
var WithDefaultRemote GnomobileOption = func(cfg *Config) error {
	cfg.Remote = "testnet.gno.berty.io:26657"
	return nil
}

var fallbackRemote = FallBackOption{
	fallback: func(cfg *Config) bool { return cfg.Remote == "" },
	opt:      WithDefaultRemote,
}

// WithFallbackRemote sets the remote node address if no address is set.
var WithFallbacRemote GnomobileOption = func(cfg *Config) error {
	if fallbackRemote.fallback(cfg) {
		return fallbackRemote.opt(cfg)
	}
	return nil
}

// --- ChainID options ---

// WithChainID sets the given chain ID.
var WithChainID = func(chainID string) GnomobileOption {
	return func(cfg *Config) error {
		cfg.ChainID = chainID
		return nil
	}
}

// WithDefaultChainID sets a default chain ID.
var WithDefaultChainID GnomobileOption = func(cfg *Config) error {
	cfg.ChainID = "dev"

	return nil
}

var fallbackChainID = FallBackOption{
	fallback: func(cfg *Config) bool { return cfg.ChainID == "" },
	opt:      WithDefaultChainID,
}

// WithFallbackChainID sets the chain ID if no chain ID is set.
var WithFallbacChainID GnomobileOption = func(cfg *Config) error {
	if fallbackChainID.fallback(cfg) {
		return fallbackChainID.opt(cfg)
	}
	return nil
}

// --- RootDir options ---

// WithRootDir sets the given root directory path.
var WithRootDir = func(rootDir string) GnomobileOption {
	return func(cfg *Config) error {
		cfg.RootDir = rootDir
		return nil
	}
}

// WithDefaultRootDir sets a default root directory in a temporary folder.
var WithDefaultRootDir GnomobileOption = func(cfg *Config) error {
	rootDir, err := os.MkdirTemp("", "gnomobile")
	if err != nil {
		return err
	}

	cfg.RootDir = rootDir

	return nil
}

var fallbackRootDir = FallBackOption{
	fallback: func(cfg *Config) bool { return cfg.RootDir == "" },
	opt:      WithDefaultRootDir,
}

// WithFallbackRootDir sets the default root directory if no directory is set.
var WithFallbackRootDir GnomobileOption = func(cfg *Config) error {
	if fallbackRootDir.fallback(cfg) {
		return fallbackRootDir.opt(cfg)
	}
	return nil
}

// --- tmpDir options ---

// WithTmpDir sets the given temporary path.
var WithTmpDir = func(path string) GnomobileOption {
	return func(cfg *Config) error {
		cfg.TmpDir = path
		return nil
	}
}

// WithDefaultTmpDir sets a default temporary path.
var WithDefaultTmpDir GnomobileOption = func(cfg *Config) error {
	// dependency
	if err := WithFallbackRootDir(cfg); err != nil {
		return err
	}

	cfg.TmpDir = cfg.RootDir

	return nil
}

var fallbackTmpDir = FallBackOption{
	fallback: func(cfg *Config) bool { return cfg.TmpDir == "" },
	opt:      WithDefaultTmpDir,
}

// WithFallbackTmpDir sets the default temporary path if no path is set.
var WithFallbackTmpDir GnomobileOption = func(cfg *Config) error {
	if fallbackTmpDir.fallback(cfg) {
		return fallbackTmpDir.opt(cfg)
	}
	return nil
}

// --- tcpPort options ---

// WithTcpPort sets the given tcp port to serve the gRPC server.
var WithTcpPort = func(port int) GnomobileOption {
	return func(cfg *Config) error {
		cfg.TcpPort = port
		return nil
	}
}

// --- useTcpListener options ---

// WithUseTcpListener sets the given tcp port to serve the gRPC server.
var WithUseTcpListener = func(choice bool) GnomobileOption {
	return func(cfg *Config) error {
		cfg.UseTcpListener = choice
		return nil
	}
}

// --- Fallback options ---

var defaults = []FallBackOption{
	fallbackLogger,
	fallbackRemote,
	fallbackChainID,
	fallbackRootDir,
	fallbackTmpDir,
}

// WithFallbackDefaults sets the default options if no option is set.
var WithFallbackDefaults GnomobileOption = func(cfg *Config) error {
	for _, def := range defaults {
		if !def.fallback(cfg) {
			continue
		}
		if err := def.opt(cfg); err != nil {
			return err
		}
	}
	return nil
}
