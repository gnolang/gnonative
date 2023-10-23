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

type userAccount struct {
	keyInfo keys.Info
}

type gnomobileService struct {
	logger     *zap.Logger
	client     *gnoclient.Client
	tcpPort    int
	socketPath string
	lock       sync.RWMutex

	// Map of key name to userAccount.
	userAccounts map[string]*userAccount
	// The active account in userAccounts, or nil if none
	activeAccount *userAccount

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
		if err := svc.createTcpGrpcServer(); err != nil {
			return nil, err
		}
	}
	if cfg.UseUdsListener {
		if err := svc.createUdsGrpcServer(cfg); err != nil {
			return nil, err
		}
	}

	return svc, nil
}

func initService(cfg *Config) (*gnomobileService, error) {
	svc := &gnomobileService{
		logger:       cfg.Logger,
		tcpPort:      cfg.TcpPort,
		userAccounts: make(map[string]*userAccount),
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

func (s *gnomobileService) createUdsGrpcServer(cfg *Config) error {
	s.logger.Debug("createUdsGrpcServer called")

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

	if err := s.runGRPCServer(); err != nil {
		return err
	}

	s.logger.Info("createUDSListener: gRPC server listens to", zap.String("path", s.socketPath))

	return nil
}

func (s *gnomobileService) createTcpGrpcServer() error {
	s.logger.Debug("createTcpGrpcServer called")

	tcpAddr := fmt.Sprintf(":%d", s.tcpPort)
	listener, err := net.Listen("tcp", tcpAddr)
	if err != nil {
		s.logger.Debug("createTcpGrpcServer error", zap.Error(err))
		return rpc.ErrCode_ErrRunGRPCServer.Wrap(err)
	}

	s.listener = listener

	// update the tcpPort field

	addr := listener.Addr().String()

	_, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		s.logger.Debug("createTcpGrpcServer error", zap.Error(err))
		return rpc.ErrCode_ErrRunGRPCServer.Wrap(err)
	}

	portInt, err := net.LookupPort("tcp", portStr)
	if err != nil {
		s.logger.Debug("createTcpGrpcServer error", zap.Error(err))
		return rpc.ErrCode_ErrRunGRPCServer.Wrap(err)
	}

	s.tcpPort = portInt

	if err := s.runGRPCServer(); err != nil {
		return err
	}

	s.logger.Info("createTcpGrpcServer: gRPC server listens to", zap.Int("port", portInt))

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
