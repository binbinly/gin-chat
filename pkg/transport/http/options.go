package http

import (
	"time"

	"gin-chat/pkg/transport"
)

var _ transport.Server = (*Server)(nil)

// Option is HTTP server option
type Option func(*options)

type options struct {
	network      string
	address      string
	readTimeout  time.Duration
	writeTimeout time.Duration
}

// WithAddress with server address.
func WithAddress(addr string) Option {
	return func(s *options) {
		s.address = addr
	}
}

// WithReadTimeout with read timeout.
func WithReadTimeout(timeout time.Duration) Option {
	return func(s *options) {
		s.readTimeout = timeout
	}
}

// WithWriteTimeout with write timeout.
func WithWriteTimeout(timeout time.Duration) Option {
	return func(s *options) {
		s.writeTimeout = timeout
	}
}

func newOptions(opt ...Option) options {
	opts := options{
		network:      "tcp",
		address:      ":9050",
		readTimeout:  5 * time.Second,
		writeTimeout: 5 * time.Second,
	}
	for _, o := range opt {
		o(&opts)
	}

	return opts
}
