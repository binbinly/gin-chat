package app

import (
	"time"
)

// nolint: golint

var (
	Conf *Config
)

type Config struct {
	Name       string
	Url        string
	Mode       string
	JwtSecret  string
	JwtTimeout int64
	CtxTimeout time.Duration
	Debug      bool
	Proxy      bool // 是否开启代理 http://[host]/ws -> ws://[host]
	HTTP       ServerConfig
	Websocket  ServerConfig
}

type ServerConfig struct {
	Network      string
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
