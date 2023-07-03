package service

import (
	"context"
	"gin-chat/pkg/app"
	"strconv"
	"strings"

	"github.com/binbinly/pkg/util"
	"github.com/pkg/errors"

	"gin-chat/internal/model"
)

// Friend 好友服务接口
type Friend interface {
	// FriendInfo 好友信息
	FriendInfo(ctx context.Context, uid, fid int) (*model.FriendModel, *model.UserModel, error)
	// FriendMyAll 我的好友
	FriendMyAll(ctx context.Context, uid int) ([]*model.User, error)
	// FriendMyListByIds 我的指定好友
	FriendMyListByIds(ctx context.Context, uid int, ids []int) ([]*model.User, error)
	// FriendMyListByTagID 我的标签好友
	FriendMyListByTagID(ctx context.Context, uid, tagID int) ([]*model.User, error)
	// FriendSetBlack 设置黑名单
	FriendSetBlack(ctx context.Context, uid, fid int, isBlack int8) error
	// FriendSetStar 设置星标
	FriendSetStar(ctx context.Context, uid, fid int, isStar int8) error
	// FriendSetMomentAuth 设置朋友圈权限
	FriendSetMomentAuth(ctx context.Context, uid, fid int, me, him int8) error
	// FriendSetRemarkTag 设置备注标签
	FriendSetRemarkTag(ctx context.Context, uid, fid int, nickname string, tags []string) error
	// FriendDestroy 删除好友
	FriendDestroy(ctx context.Context, uid, fid int) error
}

// FriendInfo 好友信息
func (s *Service) FriendInfo(ctx context.Context, uid, fid int) (f *model.FriendModel, u *model.UserModel, err error) {
	// 好友用户详情
	u, err = s.userinfo(ctx, fid)
	if err != nil {
		return
	}
	// 好友关系详情
	f, err = s.friendInfo(ctx, uid, fid)
	if err != nil {
		return
	}
	return
}

// FriendMyAll 我的好友列表
func (s *Service) FriendMyAll(ctx context.Context, uid int) (list []*model.User, err error) {
	l, err := s.repo.GetFriendAll(ctx, uid)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.friend] friend all to uid:%d", uid)
	}
	return s.friendUserList(ctx, l)
}

// FriendMyListByIds 我的好友选中列表
func (s *Service) FriendMyListByIds(ctx context.Context, uid int, ids []int) ([]*model.User, error) {
	friends, err := s.friendsByIds(ctx, uid, ids)
	if err != nil {
		return nil, err
	}
	return s.friendUserList(ctx, friends)
}

// FriendMyListByTagID 我的标签好友
func (s *Service) FriendMyListByTagID(ctx context.Context, uid, tagID int) ([]*model.User, error) {
	friends, err := s.repo.GetFriendAll(ctx, uid)
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.friend] friend all by uid:%v", uid)
	}
	list := make([]*model.FriendModel, 0)
	for _, friend := range friends {
		if friend.Tags != "" {
			tags := strings.Split(friend.Tags, ",")
			if util.InSliceStr(strconv.Itoa(tagID), tags) {
				list = append(list, friend)
			}
		}
	}
	return s.friendUserList(ctx, list)
}

// FriendSetBlack 设置加入/移除黑名单
func (s *Service) FriendSetBlack(ctx context.Context, uid, fid int, isBlack int8) error {
	f, err := s.friendInfo(ctx, uid, fid)
	if err != nil {
		return err
	}
	f.IsBlack = isBlack
	return s.repo.FriendSave(ctx, f)
}

// FriendSetStar 设置加入/移除星标
func (s *Service) FriendSetStar(ctx context.Context, uid, fid int, isStar int8) error {
	f, err := s.friendInfo(ctx, uid, fid)
	if err != nil {
		return err
	}
	f.IsStar = isStar
	return s.repo.FriendSave(ctx, f)
}

// FriendSetMomentAuth 设置朋友圈权限
func (s *Service) FriendSetMomentAuth(ctx context.Context, uid, fid int, me, him int8) error {
	info, err := s.friendInfo(ctx, uid, fid)
	if err != nil {
		return err
	}
	info.LookMe = me
	info.LookHim = him
	return s.repo.FriendSave(ctx, info)
}

// FriendSetRemarkTag 设置备注标签
func (s *Service) FriendSetRemarkTag(ctx context.Context, uid, fid int, nickname string, tags []string) error {
	info, err := s.friendInfo(ctx, uid, fid)
	if err != nil {
		return err
	}
	if len(tags) > 0 {
		tagIds, err := s.getTagIds(ctx, uid, tags)
		if err != nil {
			return err
		}
		info.Tags = util.SliceIntJoin(tagIds, ",")
	}
	info.Nickname = nickname
	return s.repo.FriendSave(ctx, info)
}

// FriendDestroy 删除好友
func (s *Service) FriendDestroy(ctx context.Context, uid, fid int) error {
	info, err := s.friendInfo(ctx, uid, fid)
	if err != nil {
		return err
	}
	return s.repo.FriendDelete(ctx, info)
}

// friendUserList 好友用户信息列表
func (s *Service) friendUserList(ctx context.Context, friends []*model.FriendModel) (list []*model.User, err error) {
	if len(friends) == 0 {
		return []*model.User{}, nil
	}
	// 批量获取用户信息
	users, err := s.batchUserinfo(ctx, friendIds(friends))
	if err != nil {
		return nil, err
	}
	list = make([]*model.User, 0, len(users))
	for _, u := range users {
		list = append(list, &model.User{
			ID:     u.ID,
			Name:   getNickname(u),
			Avatar: app.BuildResUrl(u.Avatar),
		})
	}
	return list, nil
}

// getTagIds 获取标签id列表
func (s *Service) getTagIds(ctx context.Context, uid int, tags []string) (tagIds []int, err error) {
	// 获取我的所有标签
	myTags, err := s.repo.GetTagsByUserID(ctx, uid)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.friend] tag all id:%d", uid)
	}
	newTags := make([]*model.UserTagModel, 0, len(tags))
	for _, tag := range tags {
		var id int
		for _, myTag := range myTags {
			if myTag.Name == tag { // 该标签已存在，不需要插入新标签，直接获取id返回
				id = myTag.ID
				break
			}
		}
		if id == 0 { //标签不存在，需要创建
			newTags = append(newTags, &model.UserTagModel{
				UID:  model.UID{UserID: uid},
				Name: tag,
			})
		} else {
			tagIds = append(tagIds, id)
		}
	}
	if len(newTags) > 0 {
		// 新标签批量入库
		ids, err := s.repo.TagBatchCreate(ctx, newTags)
		if err != nil {
			return nil, errors.Wrapf(err, "[service.firned] batch create")
		}
		tagIds = append(tagIds, ids...)
	}
	return tagIds, nil
}

// friendInfo
func (s *Service) friendInfo(ctx context.Context, uid, id int) (*model.FriendModel, error) {
	// 好友关系详情
	f, err := s.repo.GetFriendInfo(ctx, uid, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.friend] friend id:%d, fid:%d", uid, id)
	}
	if f == nil || f.ID == 0 {
		return nil, ErrFriendNotRecord
	}
	return f, nil
}

// friendIds 好友id列表
func friendIds(friends []*model.FriendModel) []int {
	ids := make([]int, 0, len(friends))
	for _, f := range friends {
		ids = append(ids, f.FriendID)
	}
	return ids
}

// friendsByIds 获取指定的好友信息
func (s *Service) friendsByIds(ctx context.Context, uid int, fIds []int) ([]*model.FriendModel, error) {
	//用户的全部好友
	friends, err := s.repo.GetFriendAll(ctx, uid)
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.friend] friend all by uid:%v", uid)
	}
	list := make([]*model.FriendModel, 0)
	for _, friend := range friends {
		if util.InSliceInt(friend.FriendID, fIds) { //过滤未选择的好友
			list = append(list, friend)
		}
	}
	return list, nil
}
