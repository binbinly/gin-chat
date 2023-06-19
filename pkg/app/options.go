package app

import (
	"net/url"

	"github.com/binbinly/pkg/transport"
)

type Option func(o *options)

type options struct {
	id        string
	name      string
	endpoints []*url.URL
	servers   []transport.Server
}

func WithID(id string) Option {
	return func(o *options) {
		o.id = id
	}
}

func WithName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

func WithEndpoint(endpoints ...*url.URL) Option {
	return func(o *options) { o.endpoints = endpoints }
}

func WithServer(srv ...transport.Server) Option {
	return func(o *options) {
		o.servers = srv
	}
}
