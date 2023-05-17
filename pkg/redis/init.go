package redis

import (
	"log"

	redis2 "github.com/binbinly/pkg/storage/redis"
	"github.com/redis/go-redis/v9"

	"gin-chat/pkg/config"
)

const (
	// Nil redis nil
	Nil = redis.Nil
	// Success redis成功标识
	Success = 1
)

var Client *redis.Client

// New redis
func New() *redis.Client {
	cfg, err := loadConf("")
	if err != nil {
		log.Fatalf("load orm conf err: %v", err)
	}

	c, err := redis2.NewClient(cfg)
	if err != nil {
		log.Fatalf("new redis client err: %v", err)
	}
	Client = c
	return c
}

// loadConf load redis config
func loadConf(name string) (*redis2.Config, error) {
	if name == "" {
		name = "default"
	}
	v, err := config.LoadWithType("redis")
	if err != nil {
		return nil, err
	}

	var cfg redis2.Config
	if err = v.UnmarshalKey(name, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
