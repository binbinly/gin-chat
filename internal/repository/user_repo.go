package repository

import (
	"context"
	"fmt"

	"github.com/binbinly/pkg/logger"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin-chat/internal/model"
)

// User 会员接口定义
type User interface {
	// UserCreate 创建用户
	UserCreate(ctx context.Context, user *model.UserModel) (int, error)
	// UserUpdate 更新用户信息
	UserUpdate(ctx context.Context, id int, um map[string]any) error
	// UserUpdatePwd 修改用户密码
	UserUpdatePwd(ctx context.Context, user *model.UserModel) error
	// GetUserByID 获取用户信息
	GetUserByID(ctx context.Context, id int) (*model.UserModel, error)
	// GetUserByUsername 根据用户名获取用户信息
	GetUserByUsername(ctx context.Context, username string) (*model.UserModel, error)
	// GetUsersByKeyword 关键字搜索用户
	GetUsersByKeyword(ctx context.Context, keyword string) ([]*model.UserModel, error)
	// GetUserByPhone 根据手机号获取用户信息
	GetUserByPhone(ctx context.Context, phone int64) (*model.UserModel, error)
	// GetUsersByIds 批量获取用户信息
	GetUsersByIds(ctx context.Context, ids []int) ([]*model.UserModel, error)
	// UserExist 用户是否已存在
	UserExist(ctx context.Context, username string, phone int64) (bool, error)
}

// UserCreate 创建用户
func (r *Repo) UserCreate(ctx context.Context, user *model.UserModel) (id int, err error) {
	if err = r.db.WithContext(ctx).Create(user).Error; err != nil {
		return 0, errors.Wrap(err, "[repo.user] Create")
	}
	r.delCache(ctx, userCacheKey(user.ID))

	return user.ID, nil
}

// UserUpdate 更新用户信息
func (r *Repo) UserUpdate(ctx context.Context, id int, um map[string]any) error {
	if err := r.db.WithContext(ctx).Model(&model.UserModel{}).Where("id=?", id).Updates(um).Error; err != nil {
		return errors.Wrapf(err, "[repo.user] update")
	}
	r.delCache(ctx, userCacheKey(id))

	return nil
}

// UserUpdatePwd 修改用户密码
func (r *Repo) UserUpdatePwd(ctx context.Context, user *model.UserModel) error {
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		return errors.Wrapf(err, "[repo.user] update pwd")
	}
	return nil
}

// GetUserByID 获取用户信息
func (r *Repo) GetUserByID(ctx context.Context, id int) (user *model.UserModel, err error) {

	if err = r.queryCache(ctx, userCacheKey(id), &user, func(data any) error {
		// 从数据库中获取
		if err = r.db.WithContext(ctx).First(data, id).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.user] query cache")
	}

	return
}

// GetUsersByIds 批量获取用户信息
func (r *Repo) GetUsersByIds(ctx context.Context, ids []int) (users []*model.UserModel, err error) {
	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, userCacheKey(id))
	}
	// 从cache批量获取
	cacheMap := make(map[string]*model.UserModel)
	if err = r.cache.MultiGet(ctx, keys, cacheMap, func() any {
		return &model.UserModel{}
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.user] multi get cache data")
	}

	// 查询未命中
	for _, id := range ids {
		user, ok := cacheMap[userCacheKey(id)]
		if !ok {
			user, err = r.GetUserByID(ctx, id)
			if err != nil {
				logger.Warnf("[repo.user] find id: %v err: %v", id, err)
				continue
			}
		}
		if user == nil || user.ID == 0 {
			continue
		}
		users = append(users, user)
	}
	return
}

// GetUserByUsername 根据用户名获取用户信息
func (r *Repo) GetUserByUsername(ctx context.Context, username string) (user *model.UserModel, err error) {
	user = new(model.UserModel)
	if err = r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "[repo.user] find by username")
	}
	return user, nil
}

// GetUsersByKeyword 关键字搜索用户
func (r *Repo) GetUsersByKeyword(ctx context.Context, keyword string) (users []*model.UserModel, err error) {
	//最多查询10个用户
	if err = r.db.WithContext(ctx).Where("username like ?", keyword+"%").Limit(10).Find(&users).Error; err != nil {
		return nil, errors.Wrapf(err, "[repo.user] find by keyword:%v", keyword)
	}
	return users, nil
}

// GetUserByPhone 根据手机号获取用户信息
func (r *Repo) GetUserByPhone(ctx context.Context, phone int64) (user *model.UserModel, err error) {
	user = new(model.UserModel)
	if err = r.db.WithContext(ctx).Where("phone = ?", phone).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "[repo.user] find by phone")
	}
	return user, nil
}

// UserExist 用户是否已存在
func (r *Repo) UserExist(ctx context.Context, username string, phone int64) (bool, error) {
	var c int64
	if err := r.db.WithContext(ctx).Model(&model.UserModel{}).
		Where("phone = ? or username=?", phone, username).Count(&c).Error; err != nil {
		return false, errors.Wrapf(err, "[repo.user] username %v or phone %v does not exist", username, phone)
	}
	return c > 0, nil
}

// userCacheKey 构建会员缓存key
func userCacheKey(id int) string {
	return fmt.Sprintf("user:%d", id)
}
