package repository

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin-chat/internal/model"
)

// GroupUser 群成员数据仓库
type GroupUser interface {
	// GroupUserCreate 创建群组
	GroupUserCreate(ctx context.Context, user *model.GroupUserModel) (err error)
	// GroupUserBatchCreate 批量创建
	GroupUserBatchCreate(ctx context.Context, tx *gorm.DB, users []*model.GroupUserModel) (err error)
	// GroupUserUpdateNickname 修改群成员昵称
	GroupUserUpdateNickname(ctx context.Context, userID, groupID int, nickname string) error
	// GroupUserDelete 删除成员
	GroupUserDelete(ctx context.Context, user *model.GroupUserModel) error
	// GroupUserDeleteByGroupID 删除群所有成员
	GroupUserDeleteByGroupID(ctx context.Context, tx *gorm.DB, groupID int) error
	// GetGroupUserByID 获取成员信息
	GetGroupUserByID(ctx context.Context, userID, groupID int) (info *model.GroupUserModel, err error)
	// GroupUserAll 获取所有群成员
	GroupUserAll(ctx context.Context, groupID int) (all []*model.GroupUserModel, err error)
}

// GroupUserCreate 创建群成员
func (r *Repo) GroupUserCreate(ctx context.Context, user *model.GroupUserModel) (err error) {
	if err = r.DB.WithContext(ctx).Create(user).Error; err != nil {
		return errors.Wrapf(err, "[repo.group_user] create")
	}
	r.DelCache(ctx, groupUserAllCacheKey(user.GroupID))
	return nil
}

// GroupUserBatchCreate 批量创建群成员
func (r *Repo) GroupUserBatchCreate(ctx context.Context, tx *gorm.DB, users []*model.GroupUserModel) (err error) {
	if err = tx.WithContext(ctx).Create(&users).Error; err != nil {
		return errors.Wrapf(err, "[repo.group_user] batch create")
	}
	return nil
}

// GroupUserUpdateNickname 修改群昵称
func (r *Repo) GroupUserUpdateNickname(ctx context.Context, userID, groupID int, nickname string) error {
	if err := r.DB.WithContext(ctx).Model(&model.GroupUserModel{}).
		Where("user_id=? and group_id=?", userID, groupID).
		Update("nickname", nickname).Error; err != nil {
		return errors.Wrapf(err, "[repo.group_user] update nickname")
	}
	r.DelCache(ctx, groupUserAllCacheKey(groupID))
	r.DelCache(ctx, groupUserCacheKey(userID, groupID))
	return nil
}

// GroupUserDelete 删除群成员
func (r *Repo) GroupUserDelete(ctx context.Context, user *model.GroupUserModel) error {
	if err := r.DB.WithContext(ctx).Delete(user).Error; err != nil {
		return errors.Wrapf(err, "[repo.group_user] quit group")
	}
	r.DelCache(ctx, groupUserAllCacheKey(user.GroupID))
	return nil
}

// GroupUserDeleteByGroupID 删除群下所有成员
func (r *Repo) GroupUserDeleteByGroupID(ctx context.Context, tx *gorm.DB, groupID int) (err error) {
	if err = tx.WithContext(ctx).Where("group_id=?", groupID).Delete(&model.GroupUserModel{}).Error; err != nil {
		return errors.Wrapf(err, "[repo.group_user] delete users")
	}
	r.DelCache(ctx, groupUserAllCacheKey(groupID))
	return nil
}

// GetGroupUserByID 获取群成员信息
func (r *Repo) GetGroupUserByID(ctx context.Context, userID, groupID int) (info *model.GroupUserModel, err error) {
	if err = r.QueryCache(ctx, groupUserCacheKey(userID, groupID), &info, 0, func(data any) error {
		// 从数据库中获取
		if err = r.DB.WithContext(ctx).Where("user_id=? and group_id=?", userID, groupID).
			First(data).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.group_user] query cache")
	}
	return
}

// GroupUserAll 获取群所有成员
func (r *Repo) GroupUserAll(ctx context.Context, groupID int) (list []*model.GroupUserModel, err error) {
	if err = r.QueryCache(ctx, groupUserAllCacheKey(groupID), &list, 0, func(data any) error {
		// 从数据库中获取
		if err = r.DB.WithContext(ctx).Model(&model.GroupUserModel{}).
			Where("group_id=?", groupID).Find(data).Error; err != nil {
			return err
		}
		if len(list) == 0 {
			return gorm.ErrEmptySlice
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.group_user] query cache")
	}
	return
}

func groupUserCacheKey(uid, gid int) string {
	return fmt.Sprintf("group_user:%d_%d", uid, gid)
}

func groupUserAllCacheKey(id int) string {
	return fmt.Sprintf("group_user:all:%d", id)
}
