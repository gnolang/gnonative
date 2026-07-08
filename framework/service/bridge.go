package gnonative

import (
	"context"
	"sync"
	"time"

	"github.com/gnolang/gno/tm2/pkg/errors"
	"github.com/oklog/run"
	"go.uber.org/multierr"

	api_gen "github.com/gnolang/gnonative/v5/api"
	"github.com/gnolang/gnonative/v5/service"
)

type BridgeConfig struct {
	NativeDB NativeDB
	RootDir  string
	TmpDir   string
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

	ServiceDispatcher
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

	// start service
	{
		if config.NativeDB != nil {
			// use provided NativeDB
			svcOpts = append(svcOpts,
				service.WithNativeDB(&db{NativeDB: config.NativeDB}),
			)
		}

		svcOpts = append(svcOpts,
			service.WithRootDir(config.RootDir),
			service.WithTmpDir(config.TmpDir),
		)

		serviceServer, err := service.NewGnoNativeService(svcOpts...)
		if err != nil {
			return nil, errors.Wrap(err, "unable to create bridge service")
		}
		b.serviceServer = serviceServer
	}

	// The dispatcher serves JS calls by calling the plain service API directly.
	b.ServiceDispatcher = newServiceDispatcher(b.serviceServer)

	// start Bridge
	go func() {
		b.errc <- b.workers.Run()
	}()

	return b, nil
}

func (b *Bridge) Close() error {
	var errs error

	// close bridge
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
