package repository

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin-chat/internal/model"
)

// MomentTimeline 朋友圈时间线数据仓库
type MomentTimeline interface {
	// TimelineBatchCreate 批量创建
	TimelineBatchCreate(ctx context.Context, tx *gorm.DB, models []*model.MomentTimelineModel) (err error)
	// TimelineExist 记录是否存在
	TimelineExist(ctx context.Context, userID, momentID int) (bool, error)
}

// TimelineBatchCreate 批量创建
func (r *Repo) TimelineBatchCreate(ctx context.Context, tx *gorm.DB, models []*model.MomentTimelineModel) (err error) {
	if err = tx.WithContext(ctx).CreateInBatches(&models, 500).Error; err != nil {
		return errors.Wrapf(err, "[repo.moment_timeline] batch create err")
	}
	return nil
}

// TimelineExist 记录是否存在
func (r *Repo) TimelineExist(ctx context.Context, userID, momentID int) (is bool, err error) {
	var c int64
	if err = r.DB.WithContext(ctx).Model(&model.MomentTimelineModel{}).
		Where("user_id=? and moment_id=?", userID, momentID).Count(&c).Error; err != nil {
		return false, errors.Wrap(err, "[repo.moment_timeline] query db")
	}
	return c > 0, nil
}
