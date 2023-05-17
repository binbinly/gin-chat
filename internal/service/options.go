package service

type Option func(o *options)

type options struct {
	smsReal    bool
	jwtTimeout int64
	jwtSecret  string
}

func WithSmsReal() Option {
	return func(o *options) {
		o.smsReal = true
	}
}

func WithJwtSecret(jwtSecret string) Option {
	return func(o *options) {
		o.jwtSecret = jwtSecret
	}
}

func WithJwtTimeout(t int64) Option {
	return func(o *options) {
		o.jwtTimeout = t
	}
}

func newOptions(opt ...Option) options {
	opts := options{
		jwtTimeout: 86400,
	}
	for _, o := range opt {
		o(&opts)
	}

	return opts
}
