package service

import (
	"context"
	"gin-chat/pkg/app"

	"github.com/binbinly/pkg/util"
	"github.com/pkg/errors"

	"gin-chat/internal/model"
	"gin-chat/internal/websocket"
	"gin-chat/pkg/dbs"
)

// MomentList 朋友圈列表结构
type MomentList struct {
	Moments  []*model.MomentModel
	Users    []*model.UserModel
	Likes    map[int][]int
	Comments map[int][]*model.MomentCommentModel
}

// Moment 朋友圈接口
type Moment interface {
	// MomentCreate 发布朋友圈
	MomentCreate(ctx context.Context, uid int, content, image, video, location string,
		t, sType int8, remind, see []int) error
	// MomentTimeline 我的朋友圈
	MomentTimeline(ctx context.Context, uid, offset, limit int) (*MomentList, error)
	// MomentList 好友朋友圈
	MomentList(ctx context.Context, mid, uid, offset, limit int) (*MomentList, error)
	// MomentLike 点赞
	MomentLike(ctx context.Context, uid, id int) error
	// MomentComment 评论
	MomentComment(ctx context.Context, uid, rid, id int, content string) error
}

// MomentCreate 发布朋友圈
func (s *Service) MomentCreate(ctx context.Context, uid int, content, image, video, location string,
	t, sType int8, remind, see []int) (err error) {
	// 我的好友列表
	friends, err := s.repo.GetFriendAll(ctx, uid)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] friends")
	}
	// 过滤非好友用户
	newRemind := util.SliceSmallFilter(friendIds(friends), func(v int) bool {
		return util.InSliceInt(v, remind)
	})
	m := &model.MomentModel{
		UID:      model.UID{UserID: uid},
		Content:  content,
		Image:    image,
		Video:    video,
		Location: location,
		Remind:   util.SliceIntJoin(newRemind, ","),
		Type:     t,
		SeeType:  sType,
		See:      util.SliceIntJoin(see, ","),
	}
	// 开启事务
	tx := dbs.DB.Begin()
	id, err := s.repo.MomentCreate(ctx, tx, m)
	if err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.moment] create err")
	}
	// 时间线 - 自己
	lines := []*model.MomentTimelineModel{
		{
			UID:      model.UID{UserID: uid},
			MomentID: id,
			IsOwn:    1,
		},
	}
	for _, f := range friends {
		if sType == model.MomentSeeTypeNone {
			continue
		}
		line := &model.MomentTimelineModel{
			UID:      model.UID{UserID: f.FriendID},
			MomentID: id,
			IsOwn:    0,
		}
		if sType == model.MomentSeeTypeAll {
			lines = append(lines, line)
		} else if sType == model.MomentSeeTypeOnly {
			if util.InSliceInt(f.FriendID, see) {
				lines = append(lines, line)
			}
		} else if sType == model.MomentSeeTypeExcept {
			if !util.InSliceInt(f.FriendID, see) {
				lines = append(lines, line)
			}
		}
	}
	// 朋友圈时间线创建
	if err = s.repo.TimelineBatchCreate(ctx, tx, lines); err != nil {
		tx.Rollback()
		return errors.Wrapf(err, "[service.moment] timeline batch create err")
	}
	// 提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[service.moment] tx commit err")
	}
	return s.pushMessage(ctx, uid, lines, newRemind)
}

// MomentTimeline 我的朋友圈动态
func (s *Service) MomentTimeline(ctx context.Context, uid, offset, limit int) (*MomentList, error) {
	// 朋友圈动态
	moments, err := s.repo.GetMyMoments(ctx, uid, offset, limit)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] my timeline by uid:%d", uid)
	}
	mIds, uIds := momentIds(moments)
	likes, luIds, err := s.momentLikes(ctx, mIds, nil)
	if err != nil {
		return nil, err
	}
	comments, cuIds, err := s.momentComments(ctx, mIds, nil)
	if err != nil {
		return nil, err
	}
	// 合并去重所有用户id
	userIds := util.SliceIntDeduplication(append(append(uIds, luIds...), cuIds...))
	// 批量获取用户信息
	users, err := s.batchUserinfo(ctx, userIds)
	if err != nil {
		return nil, err
	}
	return &MomentList{
		Moments:  moments,
		Users:    users,
		Likes:    likes,
		Comments: comments,
	}, nil
}

// MomentList 指定用户的动态
func (s *Service) MomentList(ctx context.Context, mid, uid, offset, limit int) (*MomentList, error) {
	// 朋友圈动态
	moments, err := s.repo.GetMomentsByUserID(ctx, mid, uid, offset, limit)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] list by uid:%d", uid)
	}
	mIds, uIds := momentIds(moments)

	// 我的好友列表
	friends, err := s.repo.GetFriendAll(ctx, mid)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.moment] friends err")
	}
	// 好友id列表
	fIds := friendIds(friends)

	likes, luIds, err := s.momentLikes(ctx, mIds, fIds)
	if err != nil {
		return nil, err
	}
	comments, cuIds, err := s.momentComments(ctx, mIds, fIds)
	if err != nil {
		return nil, err
	}

	// 合并去重所有用户id
	userIds := util.SliceIntDeduplication(append(append(uIds, luIds...), cuIds...))
	// 批量获取用户信息
	users, err := s.batchUserinfo(ctx, userIds)
	if err != nil {
		return nil, err
	}
	return &MomentList{
		Moments:  moments,
		Users:    users,
		Likes:    likes,
		Comments: comments,
	}, nil
}

// MomentLike 点赞
func (s *Service) MomentLike(ctx context.Context, uid, id int) error {
	u, authorID, err := s.momentCheck(ctx, uid, id)
	if err != nil {
		return err
	}
	// 已经点赞的用户列表
	likeIds, err := s.repo.GetLikeUserIdsByMomentID(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] like users for id:%d", id)
	}
	// 已点赞用户列表
	uIds, err := s.repo.GetLikeUserIdsByMomentID(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] exist like err uid:%d mid:%d", uid, id)
	}
	if util.InSliceInt(uid, uIds) { // 已点赞，即取消
		if err = s.repo.LikeDelete(ctx, uid, id); err != nil {
			return errors.Wrapf(err, "[service.moment] delete err uid:%d mid:%d", uid, id)
		}
		return nil
	}
	// 创建点赞记录
	mLike := &model.MomentLikeModel{
		UserID:   uid,
		MomentID: id,
	}

	if _, err = s.repo.LikeCreate(ctx, mLike); err != nil {
		return errors.Wrapf(err, "[service.moment] create like err uid:%d mid:%d", uid, id)
	}
	// 通知作者
	userIds := []int{authorID}
	// 发送其他点赞好友
	for _, i := range likeIds {
		userIds = append(userIds, i)
	}
	cs, err := s.BatchUserConn(ctx, userIds)
	if err != nil {
		return err
	}
	if err = s.ws.BatchSendConn(ctx, cs, websocket.EventMoment, &websocket.Moment{
		UserID: uid,
		Avatar: app.BuildResUrl(u.Avatar),
		Type:   "like",
	}); err != nil {
		return errors.Wrapf(err, "[service.moment] ws send to new")
	}
	return nil
}

// MomentComment 评论
func (s *Service) MomentComment(ctx context.Context, uid, rid, id int, content string) error {
	u, authorID, err := s.momentCheck(ctx, uid, id)
	if err != nil {
		return err
	}
	// 已评论用户列表
	comments, err := s.repo.GetCommentsByMomentID(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] comment users err id:%d", id)
	}
	comment := &model.MomentCommentModel{
		UID:      model.UID{UserID: uid},
		ReplyID:  rid,
		MomentID: id,
		Content:  content,
	}
	_, err = s.repo.CommentCreate(ctx, comment)
	if err != nil {
		return errors.Wrapf(err, "[service.moment] create comment err")
	}
	// 通知作者
	userIds := []int{authorID}
	// 发送其他评论好友
	for _, commentModel := range comments {
		userIds = append(userIds, commentModel.UserID)
	}
	cs, err := s.BatchUserConn(ctx, userIds)
	if err != nil {
		return err
	}
	if err = s.ws.BatchSendConn(ctx, cs, websocket.EventMoment, &websocket.Moment{
		UserID: uid,
		Avatar: app.BuildResUrl(u.Avatar),
		Type:   "comment",
	}); err != nil {
		return errors.Wrapf(err, "[service.moment] ws send to new")
	}
	return nil
}

// momentCheck 验证动态id是否合法
func (s *Service) momentCheck(ctx context.Context, uid, id int) (user *model.UserModel, authorID int, err error) {
	user, err = s.userinfo(ctx, uid)
	if err != nil {
		return nil, 0, err
	}
	// 此条动态发布者
	moment, err := s.repo.GetMomentByID(ctx, id)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "[service.moment] auther err mid:%d", id)
	}
	if moment.ID == 0 {
		return nil, 0, ErrMomentNotFound
	}
	if moment.SeeType != model.MomentSeeTypeAll { //非公开动态进一步判断权限
		// 是否存在或是否有权限
		exist, err := s.repo.TimelineExist(ctx, uid, id)
		if err != nil {
			return nil, 0, errors.Wrapf(err, "[service.moment] exist err uid:%d mid:%d", uid, id)
		}
		if !exist {
			return nil, 0, ErrMomentNotFound
		}
	}
	return user, moment.UserID, nil
}

// pushMessage 推送 朋友圈新动态消息
func (s *Service) pushMessage(ctx context.Context, uid int, lines []*model.MomentTimelineModel, remind []int) (err error) {
	u, err := s.userinfo(ctx, uid)
	if err != nil {
		return err
	}
	userIds := make([]int, len(lines))
	for i, line := range lines {
		if line.UserID != uid { //不需要给自己发送 新动态通知
			userIds[i] = line.UserID
		}
	}
	cs, err := s.BatchUserConn(ctx, userIds)
	if err != nil {
		return err
	}
	if err = s.ws.BatchSendConn(ctx, cs, websocket.EventMoment, &websocket.Moment{
		UserID: uid,
		Avatar: app.BuildResUrl(u.Avatar),
		Type:   "new",
	}); err != nil {
		return errors.Wrapf(err, "[service.moment] ws send to new")
	}
	if len(remind) > 0 { // 是否需要提醒好友
		cs, err = s.BatchUserConn(ctx, remind)
		if err != nil {
			return err
		}
		if err = s.ws.BatchSendConn(ctx, cs, websocket.EventMoment, &websocket.Moment{
			UserID: uid,
			Avatar: app.BuildResUrl(u.Avatar),
			Type:   "remind",
		}); err != nil {
			return errors.Wrapf(err, "[service.moment] ws send to remind")
		}
	}
	return nil
}

// momentLikes 朋友圈点赞信息
func (s *Service) momentLikes(ctx context.Context, mIds []int, fIds []int) (map[int][]int, []int, error) {
	likes, err := s.repo.GetLikesByMomentIds(ctx, mIds)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[service.moment] likes err mids:%v", mIds)
	}
	// 收集点赞用户id
	uIds := make([]int, 0)
	for _, like := range likes {
		for _, uid := range like {
			// 如果好友id存在，并且当前不用不在好友列表内，不需要收集
			if fIds != nil && !util.InSliceInt(uid, fIds) {
				continue
			}
			uIds = append(uIds, uid)
		}
	}
	return likes, uIds, nil
}

// momentComments 朋友圈评论信息
func (s *Service) momentComments(ctx context.Context, mIds []int, fIds []int) (map[int][]*model.MomentCommentModel, []int, error) {
	comments, err := s.repo.GetCommentsByMomentIds(ctx, mIds)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "[service.moment] comments err mids:%v", mIds)
	}
	// 收集评论，回复的用户id
	uIds := make([]int, 0)
	for _, list := range comments {
		for _, comment := range list {
			// 如果好友id存在，必须要是好友才收集用户
			if fIds == nil || util.InSliceInt(comment.UserID, fIds) {
				uIds = append(uIds, comment.UserID)
			}
			if comment.ReplyID > 0 {
				if fIds == nil || util.InSliceInt(comment.ReplyID, fIds) {
					uIds = append(uIds, comment.ReplyID)
				}
			}
		}
	}
	return comments, uIds, nil
}

// momentIds 朋友圈所有id,所有用户id
func momentIds(moments []*model.MomentModel) ([]int, []int) {
	ids := make([]int, 0, len(moments))
	uIds := make([]int, 0, len(moments))
	for _, moment := range moments {
		ids = append(ids, moment.ID)
		uIds = append(uIds, moment.UserID)
	}
	return ids, uIds
}
