package repository

import (
	"context"
	"fmt"

	"github.com/binbinly/pkg/logger"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin-chat/internal/model"
)

// MomentLike 朋友圈点赞
type MomentLike interface {
	// LikeCreate 创建
	LikeCreate(ctx context.Context, model *model.MomentLikeModel) (id int, err error)
	// LikeDelete 删除数据
	LikeDelete(ctx context.Context, user, moment int) error
	// GetLikeUserIdsByMomentID 朋友圈动态点赞用户列表
	GetLikeUserIdsByMomentID(ctx context.Context, momentID int) ([]int, error)
	// GetLikesByMomentIds 朋友圈多条动态的点赞列表
	GetLikesByMomentIds(ctx context.Context, mIds []int) (map[int][]int, error)
}

// LikeCreate 创建-点赞
func (r *Repo) LikeCreate(ctx context.Context, model *model.MomentLikeModel) (id int, err error) {
	if err = r.db.WithContext(ctx).Create(model).Error; err != nil {
		return 0, errors.Wrapf(err, "[repo.moment_like] create")
	}
	r.delCache(ctx, likeCacheKey(model.MomentID))
	return model.ID, nil
}

// LikeDelete 删除-取消点赞
func (r *Repo) LikeDelete(ctx context.Context, userID, momentID int) error {
	if err := r.db.WithContext(ctx).Where("user_id=? and moment_id=?", userID, momentID).
		Delete(&model.MomentLikeModel{}).Error; err != nil {
		return errors.Wrapf(err, "[repo.moment_like] delete")
	}
	r.delCache(ctx, likeCacheKey(momentID))
	return nil
}

// GetLikeUserIdsByMomentID 获取动态的所有点赞用户id列表
func (r *Repo) GetLikeUserIdsByMomentID(ctx context.Context, momentID int) (userIds []int, err error) {
	if err = r.queryCache(ctx, likeCacheKey(momentID), &userIds, func(data any) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Model(&model.MomentLikeModel{}).Select("user_id").
			Where("moment_id=?", momentID).Order("id asc").Pluck("user_id", data).Error; err != nil {
			return err
		}
		if len(userIds) == 0 {
			return gorm.ErrEmptySlice
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.moment_like] query cache")
	}
	return
}

// GetLikesByMomentIds 朋友圈动态点赞列表
func (r *Repo) GetLikesByMomentIds(ctx context.Context, mIds []int) (likes map[int][]int, err error) {
	likes = make(map[int][]int, len(mIds))

	keys := make([]string, 0, len(mIds))
	for _, id := range mIds {
		keys = append(keys, likeCacheKey(id))
	}
	// 从cache批量获取
	cacheMap := make(map[string]*[]int)
	if err = r.cache.MultiGet(ctx, keys, cacheMap, func() any {
		return &[]int{}
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.moment_like] multi get cache data")
	}

	// 查询未命中
	for _, id := range mIds {
		userIds, ok := cacheMap[likeCacheKey(id)]
		if !ok {
			uIds, err := r.GetLikeUserIdsByMomentID(ctx, id)
			if err != nil {
				logger.Warnf("[repo.moment_like] get like err: %v", err)
				continue
			}
			userIds = &uIds
		}
		if len(*userIds) == 0 {
			continue
		}
		likes[id] = *userIds
	}
	return likes, nil
}

func likeCacheKey(mid int) string {
	return fmt.Sprintf("moment:like:%d", mid)
}
