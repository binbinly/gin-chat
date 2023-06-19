package repository

import (
	"context"

	"github.com/pkg/errors"

	"gin-chat/internal/model"
)

// CreateMessage 创建聊天消息
func (r *Repo) CreateMessage(ctx context.Context, message model.MessageModel) (id int, err error) {
	if err = r.DB.WithContext(ctx).Create(&message).Error; err != nil {
		return 0, errors.Wrap(err, "[repo.message] create message")
	}
	return message.ID, nil
}
