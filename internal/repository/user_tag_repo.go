package repository

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin-chat/internal/model"
)

// UserTag 用户标签
type UserTag interface {
	// GetTagsByUserID 用户所有标签
	GetTagsByUserID(ctx context.Context, userID int) (list []*model.UserTag, err error)
	// TagBatchCreate 批量创建
	TagBatchCreate(ctx context.Context, tags []*model.UserTagModel) (ids []int, err error)
}

// GetTagsByUserID 用户所有标签
func (r *Repo) GetTagsByUserID(ctx context.Context, userID int) (list []*model.UserTag, err error) {
	if err = r.QueryCache(ctx, tagCacheKey(userID), &list, 0, func(data any) error {
		// 从数据库中获取
		if err = r.DB.WithContext(ctx).Model(&model.UserTagModel{}).
			Where("user_id = ? ", userID).Order(model.DefaultOrder).Scan(data).Error; err != nil {
			return err
		}
		if len(list) == 0 {
			return gorm.ErrEmptySlice
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.tag] query cache")
	}
	return
}

// TagBatchCreate 批量创建标签
func (r *Repo) TagBatchCreate(ctx context.Context, tags []*model.UserTagModel) (ids []int, err error) {
	if err = r.DB.WithContext(ctx).Create(&tags).Error; err != nil {
		return nil, errors.Wrapf(err, "[repo.tag] batch create")
	}
	r.DelCache(ctx, tagCacheKey(tags[0].UserID))
	for _, tag := range tags {
		ids = append(ids, tag.ID)
	}
	return ids, nil
}

func tagCacheKey(uid int) string {
	return fmt.Sprintf("tag:all:%d", uid)
}
