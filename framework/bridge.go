package gnomobile

import (
	"context"
	"sync"

	"github.com/oklog/run"
)

type Bridge struct {
	errc   chan error
	closec chan struct{}

	onceCloser sync.Once
	workers    run.Group
}

func NewBridge(config *BridgeConfig) (Bridge, error) {
	ctx := context.Background()

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
			return gnomobiletypes.ErrBridgeInterrupted
		}, func(error) {
			b.onceCloser.Do(func() { close(b.closec) })
		})
	}

	return b, nil
}
