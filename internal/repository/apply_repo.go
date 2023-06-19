package repository

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin-chat/internal/model"
)

// Apply 申请好友接口
type Apply interface {
	// ApplyCreate 创建
	ApplyCreate(ctx context.Context, apply model.ApplyModel) (id int, err error)
	// ApplyUpdateStatus 修改申请状态
	ApplyUpdateStatus(ctx context.Context, tx *gorm.DB, id, friendID int) (err error)
	// GetApplysByUserID 获取申请列表
	GetApplysByUserID(ctx context.Context, userID int, offset, limit int) (list []*model.ApplyModel, err error)
	// ApplyPendingCount 申请未处理数
	ApplyPendingCount(ctx context.Context, userID int) (c int64, err error)
	// GetApplyByFriendID 申请详情
	GetApplyByFriendID(ctx context.Context, userID, friendID int) (apply *model.ApplyModel, err error)
}

// ApplyCreate 创建申请记录
func (r *Repo) ApplyCreate(ctx context.Context, apply model.ApplyModel) (id int, err error) {
	if err = r.DB.WithContext(ctx).Create(&apply).Error; err != nil {
		return 0, errors.Wrap(err, "[repo.apply] create err")
	}
	return apply.ID, nil
}

// ApplyUpdateStatus 修改申请状态
func (r *Repo) ApplyUpdateStatus(ctx context.Context, tx *gorm.DB, id, friendID int) (err error) {
	if err = tx.WithContext(ctx).Model(&model.ApplyModel{}).Where("id=? && status=?",
		id, model.ApplyStatusPending).Update("status", model.ApplyStatusAgree).Error; err != nil {
		return errors.Wrapf(err, "[repo.apply] update err")
	}
	return nil
}

// GetApplysByUserID 获取申请好友列表
func (r *Repo) GetApplysByUserID(ctx context.Context, userID int, offset, limit int) (list []*model.ApplyModel, err error) {
	if err = r.DB.WithContext(ctx).Scopes(model.OffsetPage(offset, limit)).Where("friend_id = ? ", userID).
		Order(model.DefaultOrder).Find(&list).Error; err != nil {
		return nil, errors.Wrap(err, "[repo.apply] query db")
	}
	return
}

// ApplyPendingCount 待处理数量
func (r *Repo) ApplyPendingCount(ctx context.Context, userID int) (c int64, err error) {
	if err = r.DB.WithContext(ctx).Model(&model.ApplyModel{}).
		Where("friend_id=? && status=?", userID, model.ApplyStatusPending).Count(&c).Error; err != nil {
		return 0, errors.Wrapf(err, "[repo.apply] pending count db err, uid: %d", userID)
	}
	return c, nil
}

// GetApplyByFriendID 获取申请详情
func (r *Repo) GetApplyByFriendID(ctx context.Context, userID, friendID int) (apply *model.ApplyModel, err error) {
	apply = new(model.ApplyModel)
	if err = r.DB.WithContext(ctx).Where("user_id=? && friend_id=?", userID, friendID).
		Order(model.DefaultOrder).First(apply).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(err, "[repo.apply] query db err")
	}
	return apply, nil
}
