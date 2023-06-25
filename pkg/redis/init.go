package redis

import (
	"gin-chat/pkg/config"
	redis2 "github.com/binbinly/pkg/storage/redis"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"log"
	"time"
)

const (
	// Nil redis nil
	Nil = redis.Nil
	// Success redis成功标识
	Success = 1
)

var Client *redis.Client

var cfg = &Config{}

type Config struct {
	Default redis2.Config
}

// New redis
func New() *redis.Client {
	if err := loadConf(); err != nil {
		log.Fatalf("load orm conf err: %v", err)
	}

	c, err := redis2.NewClient(&cfg.Default)
	if err != nil {
		log.Fatalf("new redis client err: %v", err)
	}
	Client = c
	return c
}

// loadConf load redis config
func loadConf() error {
	if err := config.Load("redis", cfg, func(v *viper.Viper) {
		v.SetDefault("default", map[string]any{
			"Addr":         "127.0.0.1:6379",
			"Password":     "",
			"DB":           0,
			"MinIdleConn":  200,
			"DialTimeout":  60 * time.Second,
			"ReadTimeout":  500 * time.Millisecond,
			"WriteTimeout": 500 * time.Millisecond,
			"PoolSize":     100,
			"PoolTimeout":  240 * time.Second,
		})
		v.BindEnv("default.addr", "CHAT_REDIS_ADDR")
		v.BindEnv("default.password", "CHAT_REDIS_PASSWORD")
		v.BindEnv("default.DB", "CHAT_REDIS_DB")
	}); err != nil {
		return err
	}

	return nil
}
