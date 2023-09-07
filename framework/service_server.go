package gnomobile

import (
	"io"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/gnolang/gnomobile/framework/gnomobiletypes"
)

type GnomobileServiceServer interface {
	io.Closer
}

type gnomobileServiceServer struct {
	opts     []GnomobileOption
	listener net.Listener
	service  GnomobileService
	server   *grpc.Server
}

type GnomobileServerOptions struct {
	ServiceOpts []GnomobileOption
	SockPath    string
}

func (o *GnomobileServerOptions) applyDefaults() error {
	if o.ServiceOpts == nil {
		o.ServiceOpts = []GnomobileOption{}
	}

	if o.SockPath == "" {
		o.SockPath = "/tmp/gnomobile.sock"
	}

	return nil
}

// NewGnomobileServiceServer creates a new Gnomobile protocol service and runs a gRPC server.
// When finished, you must call Close().
func NewGnomobileServiceServer(opts GnomobileServerOptions) (GnomobileServiceServer, error) {
	if err := opts.applyDefaults(); err != nil {
		return nil, err
	}

	svc, err := NewGnomobileService(opts.ServiceOpts...)
	if err != nil {
		return nil, err
	}

	// delete socket if it already exists
	if _, err := os.Stat(opts.SockPath); !os.IsNotExist(err) {
		if err := os.RemoveAll(opts.SockPath); err != nil {
			return nil, err
		}
	}

	listener, err := net.Listen("unix", opts.SockPath)
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
