package utils

import "time"

type config struct {
	timeout time.Duration
	retry   int
}

type Option interface {
	apply(*config)
}

type funcOption struct {
	f func(cfg *config)
}

func (fdo *funcOption) apply(cfg *config) {
	fdo.f(cfg)
}

func newFuncOption(f func(*config)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func OptRetry(retry int) Option {
	return newFuncOption(func(cfg *config) {
		cfg.retry = retry
	})
}
func OptTimeout(timeout time.Duration) Option {
	return newFuncOption(func(cfg *config) {
		cfg.timeout = timeout
	})
}
