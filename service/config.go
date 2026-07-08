package service

import (
	"os"

	"github.com/gnolang/gno/tm2/pkg/db"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Config describes a set of settings for a GnoNativeService
type Config struct {
	Logger   *zap.Logger
	Remote   string
	ChainID  string
	NativeDB db.DB
	RootDir  string
	TmpDir   string
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

// --- NativeDB options ---

// WithNativeDB sets the given native DB.
var WithNativeDB = func(db db.DB) GnoNativeOption {
	return func(cfg *Config) error {
		cfg.NativeDB = db
		return nil
	}
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

// --- Fallback options ---

var defaults = []FallBackOption{
	fallbackLogger,
	fallbackRemote,
	fallbackChainID,
	fallbackRootDir,
	fallbackTmpDir,
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
