package repository

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin-chat/internal/model"
)

// Friend 好友接口
type Friend interface {
	// FriendCreate 创建
	FriendCreate(ctx context.Context, tx *gorm.DB, friend *model.FriendModel) (err error)
	// FriendBatchCreate 批量创建
	FriendBatchCreate(ctx context.Context, tx *gorm.DB, friends []*model.FriendModel) (err error)
	// GetFriendInfo 好友信息
	GetFriendInfo(ctx context.Context, userID, friendID int) (friend *model.FriendModel, err error)
	// GetFriendAll 我的全部好友
	GetFriendAll(ctx context.Context, userID int) (list []*model.FriendModel, err error)
	// FriendSave 保存信息
	FriendSave(ctx context.Context, friend *model.FriendModel) error
	// FriendDelete 删除好友
	FriendDelete(ctx context.Context, friend *model.FriendModel) error
}

// FriendCreate 创建好友关系
func (r *Repo) FriendCreate(ctx context.Context, tx *gorm.DB, friend *model.FriendModel) (err error) {
	if err = tx.WithContext(ctx).Create(&friend).Error; err != nil {
		return errors.Wrapf(err, "[repo.friend] create")
	}
	r.DelCache(ctx, friendAllCacheKey(friend.UserID))
	r.DelCache(ctx, friendCacheKey(friend.UserID, friend.FriendID))
	return err
}

// FriendBatchCreate 批量创建
func (r *Repo) FriendBatchCreate(ctx context.Context, tx *gorm.DB, friends []*model.FriendModel) (err error) {
	if err = tx.WithContext(ctx).Model(&model.FriendModel{}).Create(&friends).Error; err != nil {
		return errors.Wrapf(err, "[repo.friend] batch create")
	}
	for _, friend := range friends {
		r.DelCache(ctx, friendAllCacheKey(friend.UserID))
		r.DelCache(ctx, friendCacheKey(friend.UserID, friend.FriendID))
	}
	return err
}

// GetFriendInfo 好友信息
func (r *Repo) GetFriendInfo(ctx context.Context, userID, friendID int) (friend *model.FriendModel, err error) {
	if err = r.QueryCache(ctx, friendCacheKey(userID, friendID), &friend, 0, func(data any) error {
		// 从数据库中获取
		if err = r.DB.WithContext(ctx).Where("user_id=? && friend_id=?", userID, friendID).
			First(data).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.friend] query cache")
	}
	return
}

// GetFriendAll 好友列表
func (r *Repo) GetFriendAll(ctx context.Context, userID int) (list []*model.FriendModel, err error) {
	if err = r.QueryCache(ctx, friendAllCacheKey(userID), &list, 0, func(data any) error {
		// 从数据库中获取
		if err = r.DB.WithContext(ctx).Model(&model.FriendModel{}).
			Where("user_id=? and is_black=0", userID).Limit(5000).Find(data).Error; err != nil {
			return err
		}
		if len(list) == 0 {
			return gorm.ErrEmptySlice
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.friend] query cache")
	}
	return
}

// FriendSave 保存好友信息
func (r *Repo) FriendSave(ctx context.Context, friend *model.FriendModel) error {
	if err := r.DB.WithContext(ctx).Save(friend).Error; err != nil {
		return errors.Wrapf(err, "[repo.friend] save err")
	}
	r.DelCache(ctx, friendAllCacheKey(friend.UserID))
	r.DelCache(ctx, friendCacheKey(friend.UserID, friend.FriendID))
	return nil
}

// FriendDelete 删除记录
func (r *Repo) FriendDelete(ctx context.Context, friend *model.FriendModel) error {
	if err := r.DB.WithContext(ctx).Delete(friend).Error; err != nil {
		return errors.Wrapf(err, "[repo.friend] delete err")
	}
	r.DelCache(ctx, friendAllCacheKey(friend.UserID))
	r.DelCache(ctx, friendCacheKey(friend.UserID, friend.FriendID))
	return nil
}

func friendAllCacheKey(uid int) string {
	return fmt.Sprintf("friend:all:%d", uid)
}

func friendCacheKey(uid, fid int) string {
	return fmt.Sprintf("friend:%d_%d", uid, fid)
}
