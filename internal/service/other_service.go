package service

import (
	"context"

	"github.com/pkg/errors"

	"gin-chat/internal/model"
)

// ReportCreate 举报好友/群
func (s *Service) ReportCreate(ctx context.Context, UserID, friendID int, cType int8, cat, content string) error {
	is, err := s.repo.ReportExistPending(ctx, friendID)
	if err != nil {
		return errors.Wrapf(err, "[service.report] exist id:%d", friendID)
	}
	if is { // 已举报过
		return ErrReportExisted
	}
	report := &model.ReportModel{
		UID:        model.UID{UserID: UserID},
		TargetID:   friendID,
		TargetType: cType,
		Content:    content,
		Category:   cat,
		Status:     model.ReportStatusPending,
	}
	_, err = s.repo.ReportCreate(ctx, report)
	if err != nil {
		return errors.Wrapf(err, "[service.report] create report")
	}
	return nil
}
