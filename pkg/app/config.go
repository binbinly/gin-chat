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
	v.SetDefault("DfsUrl", "http://127.0.0.1:9050/group1/")
	v.SetDefault("Mode", "debug")
	v.SetDefault("JwtSecret", "TQ2MNWIB2zK0z9JCqUC6WcTG9pMTnX12CLuVSop5Xr2owx4M9JTJIzBnMMYeWwRs")
	v.SetDefault("JwtTimeout", 86400)
	v.SetDefault("LogLevel", "debug")
	v.SetDefault("LogDir", "./logs/")
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
	DfsUrl     string
	Mode       string
	LogLevel   string
	LogDir     string
	JwtSecret  string
	JwtTimeout int64
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

func BuildResUrl(path string) string {
	if len(path) == 0 {
		return ""
	}
	return Conf.DfsUrl + path
}
