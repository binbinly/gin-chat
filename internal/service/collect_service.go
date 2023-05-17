package service

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	"gin-chat/internal/model"
)

// Collect 用户收藏服务接口
type Collect interface {
	// CollectCreate 添加收藏
	CollectCreate(ctx context.Context, content string, options json.RawMessage, uid int, t int8) error
	// CollectGetList 收藏列表
	CollectGetList(ctx context.Context, uid, offset, limit int) (list []*model.CollectModel, err error)
	// CollectDestroy 删除收藏
	CollectDestroy(ctx context.Context, uid, id int) error
}

// CollectCreate 创建收藏
func (s *Service) CollectCreate(ctx context.Context, content string, options json.RawMessage, uid int, t int8) error {
	collect := &model.CollectModel{
		UID:     model.UID{UserID: uid},
		Content: content,
		Type:    t,
		Options: string(options),
	}
	if _, err := s.repo.CollectCreate(ctx, collect); err != nil {
		return errors.Wrapf(err, "[service.collect] create err")
	}
	return nil
}

// CollectGetList 我的收藏列表
func (s *Service) CollectGetList(ctx context.Context, uid, offset, limit int) ([]*model.CollectModel, error) {
	return s.repo.GetCollectsByUserID(ctx, uid, offset, limit)
}

// CollectDestroy 删除收藏
func (s *Service) CollectDestroy(ctx context.Context, uid, id int) error {
	return s.repo.CollectDelete(ctx, uid, id)
}
