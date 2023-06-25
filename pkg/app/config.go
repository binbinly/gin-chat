package app

import (
	"time"

	"github.com/spf13/viper"
)

// nolint: golint

var (
	Conf = &Config{}
)

func SetDefaultConf(v *viper.Viper) {
	v.SetDefault("Name", "gin-chat")
	v.SetDefault("Url", "http://127.0.0.1:9050")
	v.SetDefault("Mode", "debug")
	v.SetDefault("JwtSecret", "TQ2MNWIB2zK0z9JCqUC6WcTG9pMTnX12CLuVSop5Xr2owx4M9JTJIzBnMMYeWwRs")
	v.SetDefault("JwtTimeout", 86400)
	v.SetDefault("Debug", true)
	v.SetDefault("Proxy", false)
	v.SetDefault("HTTP", ServerConfig{
		Network:      "tcp",
		Addr:         ":9050",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})
	v.SetDefault("Websocket", ServerConfig{
		Network:      "tcp",
		Addr:         ":9060",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})
	v.BindEnv("name")
	v.BindEnv("url")
	v.BindEnv("debug")
	v.BindEnv("proxy")
}

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
