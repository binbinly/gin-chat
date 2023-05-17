package repository

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin-chat/internal/model"
)

// Group 群组接口
type Group interface {
	// GroupCreate 创建群组
	GroupCreate(ctx context.Context, tx *gorm.DB, group *model.GroupModel) (id int, err error)
	// GroupSave 保存群组
	GroupSave(ctx context.Context, group *model.GroupModel) error
	// GroupDelete 删除群组
	GroupDelete(ctx context.Context, tx *gorm.DB, group *model.GroupModel) (err error)
	// GetGroupByID 获取群组信息
	GetGroupByID(ctx context.Context, id int) (info *model.GroupModel, err error)
	// GetGroupsByUserID 获取我的群组列表
	GetGroupsByUserID(ctx context.Context, userID int) (list []*model.GroupList, err error)
}

// GroupCreate 创建群组
func (r *Repo) GroupCreate(ctx context.Context, tx *gorm.DB, group *model.GroupModel) (id int, err error) {
	if err = tx.WithContext(ctx).Create(&group).Error; err != nil {
		return 0, errors.Wrapf(err, "[repo.group] create")
	}
	r.delCache(ctx, groupAllCacheKey(group.UserID))
	return group.ID, nil
}

// GroupSave 保存群组信息
func (r *Repo) GroupSave(ctx context.Context, group *model.GroupModel) (err error) {
	if err = r.db.WithContext(ctx).Save(group).Error; err != nil {
		return errors.Wrapf(err, "[repo.group] save")
	}
	r.delCache(ctx, groupAllCacheKey(group.UserID))
	r.delCache(ctx, groupCacheKey(group.ID))
	return nil
}

// GroupDelete 删除群
func (r *Repo) GroupDelete(ctx context.Context, tx *gorm.DB, group *model.GroupModel) (err error) {
	if err = tx.WithContext(ctx).Delete(group).Error; err != nil {
		return errors.Wrapf(err, "[repo.group] delete")
	}
	r.delCache(ctx, groupAllCacheKey(group.UserID))
	r.delCache(ctx, groupCacheKey(group.ID))
	return err
}

// GetGroupByID 获取群组信息
func (r *Repo) GetGroupByID(ctx context.Context, id int) (group *model.GroupModel, err error) {
	if err = r.queryCache(ctx, groupCacheKey(id), &group, func(data any) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).First(data, id).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.group] query cache")
	}
	return
}

// GetGroupsByUserID 群组列表
func (r *Repo) GetGroupsByUserID(ctx context.Context, userID int) (list []*model.GroupList, err error) {
	if err = r.queryCache(ctx, groupAllCacheKey(userID), &list, func(data any) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).Model(&model.GroupUserModel{}).Distinct().Select("`group`.id, `group`.name, `group`.avatar").
			Joins("left join `group` on `group`.id = group_user.group_id").
			Where("group_user.user_id=?", userID).Scan(&data).Error; err != nil {
			return err
		}
		if len(list) == 0 {
			return gorm.ErrEmptySlice
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.group] query cache")
	}
	return
}

func groupCacheKey(id int) string {
	return fmt.Sprintf("group:%d", id)
}

func groupAllCacheKey(uid int) string {
	return fmt.Sprintf("group:all:%d", uid)
}
