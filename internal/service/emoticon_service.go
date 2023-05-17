package service

import (
	"context"

	"gin-chat/internal/model"
)

// EmoticonCat 表情包分类
func (s *Service) EmoticonCat(ctx context.Context) (list []*model.Emoticon, err error) {
	return s.repo.GetEmoticonCatAll(ctx)
}

// Emoticon 分类下的表情
func (s *Service) Emoticon(ctx context.Context, cat string) (list []*model.Emoticon, err error) {
	return s.repo.GetEmoticonListByCat(ctx, cat)
}
