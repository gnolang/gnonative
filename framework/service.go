package gnomobile

import (
	"io"
	"net"
	"os"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/gnolang/gnomobile/framework/gnomobiletypes"
)

type GnomobileService interface {
	gnomobiletypes.GnomobileServiceServer
}

type gnomobileService struct {
	logger *zap.Logger

	gnomobiletypes.UnimplementedGnomobileServiceServer
}

var _ GnomobileService = (*gnomobileService)(nil)

type GnomobileServiceServer interface {
	io.Closer
}

type gnomobileServiceServer struct {
	opts     []GnomobileOption
	listener net.Listener
	service  GnomobileService
	server   *grpc.Server
}

type GnomobileServerOption struct {
	ServiceOpts []GnomobileOption
	SockAddr    string
}

func (o *GnomobileServerOption) applyDefaults() error {
	if o.ServiceOpts == nil {
		o.ServiceOpts = []GnomobileOption{}
	}

	if o.SockAddr == "" {
		o.SockAddr = "/tmp/gnomobile.sock"
	}

	return nil
}

type GnomobileOption func(*gnomobileService) error

// NewGnomobileServiceServer creates a new Gnomobile protocol service and runs a gRPC server.
// When finished, you must call Close().
func NewGnomobileServiceServer(opts GnomobileServerOption) (GnomobileServiceServer, error) {
	if err := opts.applyDefaults(); err != nil {
		return nil, err
	}

	svc, err := NewGnomobileService(opts.ServiceOpts...)
	if err != nil {
		return nil, err
	}

	// delete socket if it already exists
	if _, err := os.Stat(opts.SockAddr); !os.IsNotExist(err) {
		if err := os.RemoveAll(opts.SockAddr); err != nil {
			return nil, err
		}
	}

	listener, err := net.Listen("unix", opts.SockAddr)
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer()

	gnomobiletypes.RegisterGnomobileServiceServer(s, svc)
	go func() {
		// we dont need to log the error
		_ = s.Serve(listener)
	}()

	return &gnomobileServiceServer{
		listener: listener,
		server:   s,
		service:  svc,
	}, nil
}

func (s *gnomobileServiceServer) Close() error {
	return s.listener.Close()
}

func NewGnomobileService(opts ...GnomobileOption) (GnomobileService, error) {
	svc := &gnomobileService{}

	withDefaultOpts := make([]GnomobileOption, len(opts))
	copy(withDefaultOpts, opts)
	withDefaultOpts = append(withDefaultOpts, WithFallbackDefaults)
	for _, opt := range withDefaultOpts {
		if err := opt(svc); err != nil {
			return nil, err
		}
	}

	return svc, nil
}

// FallBackOption is a structure that permit to fallback to a default option if the option is not set.
type FallBackOption struct {
	fallback func(s *gnomobileService) bool
	opt      GnomobileOption
}

// WithLogger set the given logger.
var WithLogger = func(l *zap.Logger) GnomobileOption {
	return func(s *gnomobileService) error {
		s.logger = l
		return nil
	}
}

// WithDefaultLogger init a noop logger.
var WithDefaultLogger GnomobileOption = func(s *gnomobileService) error {
	s.logger = zap.NewNop()
	return nil
}

var fallbackLogger = FallBackOption{
	fallback: func(s *gnomobileService) bool { return s.logger == nil },
	opt:      WithDefaultLogger,
}

// WithFallbackLogger set the logger if no logger is set.
var WithFallbackLogger GnomobileOption = func(s *gnomobileService) error {
	if fallbackLogger.fallback(s) {
		return fallbackLogger.opt(s)
	}
	return nil
}

var defaults = []FallBackOption{
	fallbackLogger,
}

// WithFallbackDefaults set the default options if no option is set.
var WithFallbackDefaults GnomobileOption = func(s *gnomobileService) error {
	for _, def := range defaults {
		if !def.fallback(s) {
			continue
		}
		if err := def.opt(s); err != nil {
			return err
		}
	}
	return nil
}
