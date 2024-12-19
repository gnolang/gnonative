package service

import (
	"os"
	"path/filepath"

	api_gen "github.com/gnolang/gnonative/api/gen/go"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const DEFAULT_TCP_ADDR = ":26658"
const DEFAULT_SOCKET_SUBDIR = "s"
const DEFAULT_SOCKET_FILE = "gno"

// Config describes a set of settings for a GnoNativeService
type Config struct {
	Logger             *zap.Logger
	Remote             string
	ChainID            string
	RootDir            string
	TmpDir             string
	TcpAddr            string
	UdsPath            string
	UseTcpListener     bool
	DisableUdsListener bool
}

type GnoNativeOption func(cfg *Config) error

func (cfg *Config) applyOptions(opts ...GnoNativeOption) error {
	withDefaultOpts := make([]GnoNativeOption, len(opts))
	copy(withDefaultOpts, opts)
	withDefaultOpts = append(withDefaultOpts, WithFallbackDefaults)
	for _, opt := range withDefaultOpts {
		if err := opt(cfg); err != nil {
			return err
		}
	}
	return nil
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

// FallBackOption is a structure that permits to fallback to a default option if the option is not set.
type FallBackOption struct {
	fallback func(cfg *Config) bool
	opt      GnoNativeOption
}

// --- Logger options ---

// WithLogger set the given logger.
var WithLogger = func(l *zap.Logger) GnoNativeOption {
	return func(cfg *Config) error {
		cfg.Logger = l
		return nil
	}
}

// WithDefaultLogger init a noop logger.
var WithDefaultLogger GnoNativeOption = func(cfg *Config) error {
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
var WithFallbackLogger GnoNativeOption = func(cfg *Config) error {
	if fallbackLogger.fallback(cfg) {
		return fallbackLogger.opt(cfg)
	}
	return nil
}

// --- Remote options ---

// WithRemote sets the given remote node address.
var WithRemote = func(remote string) GnoNativeOption {
	return func(cfg *Config) error {
		cfg.Remote = remote
		return nil
	}
}

// WithDefaultRemote inits a default remote node address.
var WithDefaultRemote GnoNativeOption = func(cfg *Config) error {
	cfg.Remote = "127.0.0.1:26657"
	return nil
}

var fallbackRemote = FallBackOption{
	fallback: func(cfg *Config) bool { return cfg.Remote == "" },
	opt:      WithDefaultRemote,
}

// WithFallbackRemote sets the remote node address if no address is set.
var WithFallbacRemote GnoNativeOption = func(cfg *Config) error {
	if fallbackRemote.fallback(cfg) {
		return fallbackRemote.opt(cfg)
	}
	return nil
}

// --- ChainID options ---

// WithChainID sets the given chain ID.
var WithChainID = func(chainID string) GnoNativeOption {
	return func(cfg *Config) error {
		cfg.ChainID = chainID
		return nil
	}
}

// WithDefaultChainID sets a default chain ID.
var WithDefaultChainID GnoNativeOption = func(cfg *Config) error {
	cfg.ChainID = "dev"

	return nil
}

var fallbackChainID = FallBackOption{
	fallback: func(cfg *Config) bool { return cfg.ChainID == "" },
	opt:      WithDefaultChainID,
}

// WithFallbackChainID sets the chain ID if no chain ID is set.
var WithFallbacChainID GnoNativeOption = func(cfg *Config) error {
	if fallbackChainID.fallback(cfg) {
		return fallbackChainID.opt(cfg)
	}
	return nil
}

// --- RootDir options ---

// WithRootDir sets the given root directory path.
var WithRootDir = func(rootDir string) GnoNativeOption {
	return func(cfg *Config) error {
		cfg.RootDir = rootDir
		return nil
	}
}

// WithDefaultRootDir sets a default root directory in a temporary folder.
var WithDefaultRootDir GnoNativeOption = func(cfg *Config) error {
	rootDir, err := os.MkdirTemp("", "gnonative")
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
var WithFallbackRootDir GnoNativeOption = func(cfg *Config) error {
	if fallbackRootDir.fallback(cfg) {
		return fallbackRootDir.opt(cfg)
	}
	return nil
}

// --- tmpDir options ---

// WithTmpDir sets the given temporary path.
var WithTmpDir = func(path string) GnoNativeOption {
	return func(cfg *Config) error {
		cfg.TmpDir = path
		return nil
	}
}

// WithDefaultTmpDir sets a default temporary path.
var WithDefaultTmpDir GnoNativeOption = func(cfg *Config) error {
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
var WithFallbackTmpDir GnoNativeOption = func(cfg *Config) error {
	if fallbackTmpDir.fallback(cfg) {
		return fallbackTmpDir.opt(cfg)
	}
	return nil
}

// --- tcpAddr options ---

// WithTcpAddr sets the given TCP address to serve the gRPC server.
// If no TCP address is defined, a default will be used.
// If the TCP port is set to 0, a random port number will be chosen.
var WithTcpAddr = func(addr string) GnoNativeOption {
	return func(cfg *Config) error {
		cfg.TcpAddr = addr
		return nil
	}
}

// WithDefaultTcpAddr sets a default TCP addr to listen to.
var WithDefaultTcpAddr GnoNativeOption = func(cfg *Config) error {
	cfg.TcpAddr = DEFAULT_TCP_ADDR

	return nil
}

var fallbackTcpAddr = FallBackOption{
	fallback: func(cfg *Config) bool { return cfg.TcpAddr == "" },
	opt:      WithDefaultTcpAddr,
}

// WithDefaultTcpAddr sets a default TCP addr to listen to if no address is set.
var WithFallbackTcpAddr GnoNativeOption = func(cfg *Config) error {
	if fallbackTcpAddr.fallback(cfg) {
		return fallbackTcpAddr.opt(cfg)
	}
	return nil
}

// --- udsPath options ---

// WithUdsPath sets the given Unix Domain Socket path to serve the gRPC server.
// If no UDS socket is defined, a default will be used.
var WithUdsPath = func(path string) GnoNativeOption {
	return func(cfg *Config) error {
		absPath, err := filepath.Abs(path)
		if err != nil {
			return err
		}

		cfg.UdsPath = absPath
		return nil
	}
}

// WithDefaultUdsPath sets a default UDS path to listen to.
var WithDefaultUdsPath GnoNativeOption = func(cfg *Config) error {
	// dependency
	if err := WithFallbackTmpDir(cfg); err != nil {
		return err
	}

	// create a socket subdirectory
	sockDir := filepath.Join(cfg.TmpDir, DEFAULT_SOCKET_SUBDIR)
	if err := os.MkdirAll(sockDir, 0700); err != nil {
		return api_gen.ErrCode_ErrInitService.Wrap(err)
	}

	cfg.UdsPath = filepath.Join(sockDir, DEFAULT_SOCKET_FILE)

	return nil
}

var fallbackUdsPath = FallBackOption{
	fallback: func(cfg *Config) bool { return cfg.UdsPath == "" },
	opt:      WithDefaultUdsPath,
}

// WithDefaultUdsPath sets a default UDS path to listen to if no path is set.
var WithFallbackUdsPath GnoNativeOption = func(cfg *Config) error {
	if fallbackUdsPath.fallback(cfg) {
		return fallbackUdsPath.opt(cfg)
	}
	return nil
}

// --- listener options ---

// WithUseTcpListener sets the gRPC server to serve on a TCP listener.
var WithUseTcpListener = func() GnoNativeOption {
	return func(cfg *Config) error {
		cfg.UseTcpListener = true
		return nil
	}
}

// WithDisableUdsListener sets the gRPC server to serve on a TCP listener.
var WithDisableUdsListener = func() GnoNativeOption {
	return func(cfg *Config) error {
		cfg.DisableUdsListener = true
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
	fallbackTcpAddr,
	fallbackUdsPath,
}

// WithFallbackDefaults sets the default options if no option is set.
var WithFallbackDefaults GnoNativeOption = func(cfg *Config) error {
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
