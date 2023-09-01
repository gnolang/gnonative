package gnomobile

import (
	"context"
	"sync"
	"time"

	"github.com/gnolang/gno/tm2/pkg/errors"
	"github.com/oklog/run"
	"go.uber.org/multierr"

	gnotypes "github.com/gnolang/gnomobile/framework/gnomobiletypes"
)

type BridgeConfig struct {
	RootDir string
}

func NewBridgeConfig() *BridgeConfig {
	return &BridgeConfig{}
}

// func (c *BridgeConfig) SetRootDir(rootDir string) {
// 	c.RootDir = rootDir
// }

type Bridge struct {
	errc   chan error
	closec chan struct{}

	onceCloser sync.Once
	workers    run.Group

	serviceServer GnomobileServiceServer
}

func NewBridge(config *BridgeConfig) (*Bridge, error) {
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
			return gnotypes.ErrCode_ErrBridgeInterrupted
		}, func(error) {
			b.onceCloser.Do(func() { close(b.closec) })
		})
	}

	// setup native bridge client
	{
		opts := GnomobileServerOptions{
			SockAddr: config.RootDir + "/gnomobile.sock",
		}

		serviceServer, err := NewGnomobileServiceServer(opts)
		if err != nil {
			return nil, errors.Wrap(err, "unable to create bridge service")
		}
		b.serviceServer = serviceServer
	}

	// start Bridge
	go func() {
		b.errc <- b.workers.Run()
	}()

	return b, nil
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

		if !gnotypes.Is(err, gnotypes.ErrCode_ErrBridgeInterrupted) {
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
