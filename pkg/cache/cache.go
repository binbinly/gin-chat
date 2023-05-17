package cache

import (
	"context"
	"time"

	"github.com/binbinly/pkg/cache"
	"github.com/binbinly/pkg/codec"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

// Cache 缓存结构
type Cache struct {
	Redis cache.Cache
	//localCache cache.Driver

	opts Options
}

// NewCache 实例化cache
func NewCache(rdb *redis.Client, opts ...Option) *Cache {
	o := NewOptions(opts...)
	return &Cache{
		Redis: cache.NewRedisCache(rdb, o.prefix, codec.JSONEncoding{}),
		opts:  o,
	}
}

// Set 写入缓存
func (c *Cache) Set(ctx context.Context, key string, data any) error {
	return c.Redis.Set(ctx, key, data, c.opts.expire)
}

// SetEX 写入缓存设置过期时间
func (c *Cache) SetEX(ctx context.Context, key string, data any, d time.Duration) error {
	return c.Redis.Set(ctx, key, data, d)
}

// Get 获取缓存
func (c *Cache) Get(ctx context.Context, key string, data any) (err error) {
	return c.Redis.Get(ctx, key, data)
}

// MultiGet 批量获取用户cache
func (c *Cache) MultiGet(ctx context.Context, keys []string, valueMap any, obj func() any) (err error) {
	return c.Redis.MultiGet(ctx, keys, valueMap, obj)
}

// Del 删除缓存
func (c *Cache) Del(ctx context.Context, key string) error {
	return c.Redis.Del(ctx, key)
}

// SetNotFound 设置空,防缓存穿透
func (c *Cache) SetNotFound(ctx context.Context, key string) error {
	return c.Redis.SetCacheWithNotFound(ctx, key)
}

// IsNotFound 是否存在
func (c *Cache) IsNotFound(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
