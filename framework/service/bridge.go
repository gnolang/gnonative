package gnonative

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gnolang/gno/tm2/pkg/errors"
	"github.com/oklog/run"
	"github.com/peterbourgon/unixtransport"
	"go.uber.org/multierr"

	api_gen "github.com/gnolang/gnonative/api/gen/go"
	"github.com/gnolang/gnonative/api/gen/go/_goconnect"
	"github.com/gnolang/gnonative/service"
)

type BridgeConfig struct {
	RootDir            string
	TmpDir             string
	UseTcpListener     bool
	DisableUdsListener bool
}

func NewBridgeConfig() *BridgeConfig {
	return &BridgeConfig{}
}

type Bridge struct {
	errc   chan error
	closec chan struct{}

	onceCloser sync.Once
	workers    run.Group

	serviceServer service.GnoNativeService

	ServiceClient
}

func NewBridge(config *BridgeConfig) (*Bridge, error) {
	svcOpts := []service.GnoNativeOption{}

	// create bridge instance
	b := &Bridge{
		errc:   make(chan error),
		closec: make(chan struct{}),
	}

	// create cancel service
	{
		b.workers.Add(func() error {
			// wait for closing signal
			<-b.closec
			return api_gen.ErrCode_ErrBridgeInterrupted
		}, func(error) {
			b.onceCloser.Do(func() { close(b.closec) })
		})
	}

	// start gRPC service
	{
		svcOpts = append(svcOpts,
			service.WithRootDir(config.RootDir),
			service.WithTmpDir(config.TmpDir),
		)

		if config.UseTcpListener {
			svcOpts = append(svcOpts, service.WithUseTcpListener())
		}

		if config.DisableUdsListener {
			svcOpts = append(svcOpts, service.WithDisableUdsListener())
		}

		serviceServer, err := service.NewGnoNativeService(svcOpts...)
		if err != nil {
			return nil, errors.Wrap(err, "unable to create bridge service")
		}
		b.serviceServer = serviceServer
	}

	// create native bridge client
	{
		var httpClient *http.Client
		var address string

		// prefer a TCP connection if available
		// because iOS simulator devices cannot use UDS connections
		if config.UseTcpListener {
			httpClient = http.DefaultClient
			port := b.serviceServer.GetTcpPort()
			address = fmt.Sprintf("http://localhost:%d", port)
		} else {
			path := b.serviceServer.GetUDSPath()
			address = fmt.Sprintf("http+unix://%s:", path)

			t := &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					conn, err := net.DialTimeout(network, addr, time.Second*2)
					if err != nil {
						return nil, err
					}
					conn.SetDeadline(time.Now().Add(time.Second * 2))
					return conn, nil
				},
			}
			unixtransport.Register(t)
			httpClient = &http.Client{Transport: t}
		}

		client := _goconnect.NewGnoNativeServiceClient(
			httpClient,
			address,
		)
		b.ServiceClient = NewServiceClient(client)
	}

	// start Bridge
	go func() {
		b.errc <- b.workers.Run()
	}()

	return b, nil
}

func (b *Bridge) GetUDSPath() string {
	if b.serviceServer == nil {
		return ""
	}

	return b.serviceServer.GetUDSPath()
}

func (b *Bridge) GetTcpPort() int {
	if b.serviceServer == nil {
		return 0
	}

	return b.serviceServer.GetTcpPort()
}

func (b *Bridge) GetTcpAddr() string {
	if b.serviceServer == nil {
		return ""
	}

	return b.serviceServer.GetTcpAddr()
}

func (b *Bridge) Close() error {
	var errs error

	// close gRPC bridge
	if !b.isClosed() {
		// send close signal
		b.onceCloser.Do(func() { close(b.closec) })

		// set close timeout
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)

		// wait or die
		var err error
		select {
		case err = <-b.errc:
		case <-ctx.Done():
			err = ctx.Err()
		}

		b.serviceServer.Close()

		if !api_gen.Is(err, api_gen.ErrCode_ErrBridgeInterrupted) {
			errs = multierr.Append(errs, err)
		}

		cancel()
	}

	return errs
}

func (b *Bridge) isClosed() bool {
	select {
	case <-b.closec:
		return true
	default:
		return false
	}
}
