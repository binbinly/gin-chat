package repository

import (
	"context"

	"github.com/pkg/errors"

	"gin-chat/internal/model"
)

// ReportCreate 创建举报
func (r *Repo) ReportCreate(ctx context.Context, report *model.ReportModel) (int, error) {
	if err := r.db.WithContext(ctx).Create(&report).Error; err != nil {
		return 0, errors.Wrap(err, "[repo.report] create report")
	}
	return report.ID, nil
}

// ReportExistPending 待处理记录是否存在
func (r *Repo) ReportExistPending(ctx context.Context, targetID int) (bool, error) {
	var c int64
	if err := r.db.WithContext(ctx).Model(&model.ReportModel{}).
		Where("target_id=? && status=?", targetID, model.ReportStatusPending).Count(&c).Error; err != nil {
		return false, errors.Wrapf(err, "[repo.report] exist")
	}
	return c > 0, nil
}
