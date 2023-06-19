package repository

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"gin-chat/internal/model"
)

const (
	_emoticonCatAllCacheKey = "emoticon:cat:all"
)

// GetEmoticonCatAll 获取表情所有分类
func (r *Repo) GetEmoticonCatAll(ctx context.Context) (list []*model.Emoticon, err error) {
	if err = r.QueryCache(ctx, _emoticonCatAllCacheKey, &list, 0, func(data any) error {
		// 从数据库中获取
		if err = r.DB.WithContext(ctx).Model(&model.EmoticonModel{}).Select("ANY_VALUE(id),ANY_VALUE(name),ANY_VALUE(url),category").
			Group("category").Scan(data).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.emoticon] query cache")
	}
	return
}

// GetEmoticonListByCat 获取分类下所有表情
func (r *Repo) GetEmoticonListByCat(ctx context.Context, cat string) (list []*model.Emoticon, err error) {
	if err = r.QueryCache(ctx, emoticonCacheKey(cat), &list, 0, func(data any) error {
		// 从数据库中获取
		if err = r.DB.WithContext(ctx).Model(&model.EmoticonModel{}).
			Where("category=?", cat).Scan(data).Error; err != nil {
			return errors.Wrap(err, "[repo.emoticon] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.emoticon] query cache")
	}
	return
}

func emoticonCacheKey(cat string) string {
	return fmt.Sprintf("emoticon:%s", cat)
}
