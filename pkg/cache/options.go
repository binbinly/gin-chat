package cache

import "time"

const (
	// _prefixKey 缓存前缀
	_prefixKey = "cache:"
	// _expireTime 缓存过期时间
	_expireTime = time.Hour * 24
)

type Option func(*Options)

type Options struct {
	expire time.Duration
	prefix string
}

func NewOptions(opt ...Option) Options {
	opts := Options{
		expire: _expireTime,
		prefix: _prefixKey,
	}

	for _, o := range opt {
		o(&opts)
	}

	return opts
}

func WithExpire(d time.Duration) Option {
	return func(o *Options) {
		o.expire = d
	}
}

func WithPrefix(prefix string) Option {
	return func(o *Options) {
		o.prefix = prefix
	}
}
