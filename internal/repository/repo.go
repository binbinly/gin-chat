package repository

import (
	"context"
	"reflect"

	"github.com/binbinly/pkg/logger"
	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"gin-chat/internal/model"
	"gin-chat/pkg/cache"
	"gin-chat/pkg/redis"
)

var g singleflight.Group

var _ IRepo = (*Repo)(nil)

// IRepo 数据仓库接口
type IRepo interface {
	User
	Apply
	Collect
	Friend
	Group
	GroupUser
	Moment
	MomentComment
	MomentTimeline
	MomentLike
	UserTag

	GetEmoticonCatAll(ctx context.Context) (list []*model.Emoticon, err error)
	GetEmoticonListByCat(ctx context.Context, cat string) (list []*model.Emoticon, err error)
	ReportCreate(ctx context.Context, report *model.ReportModel) (int, error)
	ReportExistPending(ctx context.Context, targetID int) (bool, error)
	CreateMessage(ctx context.Context, message model.MessageModel) (id int, err error)
	Close()
}

// Repo mysql struct
type Repo struct {
	db    *gorm.DB
	cache *cache.Cache
}

// New new a Dao and return
func New(db *gorm.DB, cache *cache.Cache) IRepo {
	return &Repo{
		db:    db,
		cache: cache,
	}
}

// Close release mysql connection
func (r *Repo) Close() {

}

// queryCache 查询启用缓存
// 缓存的更新策略使用 Cache Aside Pattern
// see: https://coolshell.cn/articles/17416.html
func (r *Repo) queryCache(ctx context.Context, key string, data any, query func(any) error) (err error) {
	// 从cache获取
	err = r.cache.Get(ctx, key, data)
	if r.cache.IsNotFound(err) {
		// 空数据也需要返回空的数据结构，实例化结构
		reflectValue := reflect.ValueOf(data)
		for reflectValue.Kind() == reflect.Ptr {
			if reflectValue.IsNil() && reflectValue.CanAddr() {
				reflectValue.Set(reflect.New(reflectValue.Type().Elem()))
			}

			reflectValue = reflectValue.Elem()
		}
		logger.Infof("[repo] key %v is empty", key)
		return nil
	} else if err != nil && err != redis.Nil {
		return errors.Wrapf(err, "[repo] get cache by key: %s", key)
	}

	// 检查数据类型
	elem := reflect.ValueOf(data).Elem()
	switch elem.Kind() {
	case reflect.String:
		if elem.String() != "" {
			logger.Infof("[repo] get from string cache, key: %v", key)
			return
		}
	default:
		if !elem.IsNil() { // 已经从缓存中取到了数据，直接返回
			logger.Infof("[repo] get from obj cache, key: %v, kind:%v", key, elem.Kind())
			return
		}
	}

	// use sync/singleflight mode to get data
	// why not use redis lock? see this topic: https://redis.io/topics/distlock
	// demo see: https://github.com/go-demo/singleflight-demo/blob/master/main.go
	// https://juejin.cn/post/6844904084445593613
	_, err, _ = g.Do(key, func() (any, error) {
		// 从数据库中获取
		err = query(data)
		// if data is empty, set not found cache to prevent cache penetration(缓存穿透)
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, gorm.ErrEmptySlice) {
			if err = r.cache.SetNotFound(ctx, key); err != nil {
				logger.Warnf("[repo] SetCacheWithNotFound err, key: %s", key)
			}
			return data, nil
		} else if err != nil {
			return nil, errors.Wrapf(err, "[repo] query db")
		}

		// set cache
		if err = r.cache.Set(ctx, key, data); err != nil {
			return nil, errors.Wrapf(err, "[repo] set data to cache key: %s", key)
		}
		return data, nil
	})
	if err != nil {
		return errors.Wrapf(err, "[repo] get err via single flight do key: %s", key)
	}

	return nil
}

// delCache 删除缓存
func (r *Repo) delCache(ctx context.Context, key string) {
	if err := r.cache.Del(ctx, key); err != nil {
		logger.Warnf("[repo] del cache key: %v", key)
	}
}
