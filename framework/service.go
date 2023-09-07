package gnomobile

import (
	"github.com/gnolang/gnomobile/framework/gnomobiletypes"
	"go.uber.org/zap"
)

type GnomobileService interface {
	gnomobiletypes.GnomobileServiceServer
}

type gnomobileService struct {
	logger *zap.Logger

	gnomobiletypes.UnimplementedGnomobileServiceServer
}

var _ GnomobileService = (*gnomobileService)(nil)

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

type GnomobileOption func(*gnomobileService) error

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
