package resource

import (
	"gin-chat/internal/model"
)

// MomentListResource 朋友圈列表结构
func MomentListResource(moments []*model.MomentModel, users []*model.UserModel, mLikes map[int][]int,
	mComments map[int][]*model.MomentCommentModel) []*model.MomentList {
	um := usersToMap(users)
	likes := likesToMap(mLikes, um)
	comments := commentsToMap(mComments, um)
	list := make([]*model.MomentList, 0)
	for _, moment := range moments {
		if user, ok := um[moment.UserID]; ok {
			m := &model.MomentList{
				ID:        moment.ID,
				Content:   moment.Content,
				Image:     moment.Image,
				Video:     moment.Video,
				Location:  moment.Location,
				Type:      moment.Type,
				CreatedAt: moment.CreatedAt,
				User:      user,
				Likes:     make([]*model.User, 0),
				Comments:  make([]*model.Comment, 0),
			}
			if l, o := likes[moment.ID]; o {
				m.Likes = l
			}
			if c, o := comments[moment.ID]; o {
				m.Comments = c
			}
			list = append(list, m)
		}
	}
	return list
}

// likesToMap 点赞列表数据格式化为map
func likesToMap(likes map[int][]int, users map[int]*model.User) map[int][]*model.User {
	ml := make(map[int][]*model.User)
	if len(likes) == 0 {
		return ml
	}
	for mid, like := range likes {
		for _, uid := range like {
			if user, ok := users[uid]; ok {
				if _, o := ml[mid]; o {
					ml[mid] = append(ml[mid], user)
				} else {
					ml[mid] = []*model.User{user}
				}
			}
		}
	}
	return ml
}

// commentsToMap 评论列表格式化为map
func commentsToMap(mComments map[int][]*model.MomentCommentModel, users map[int]*model.User) map[int][]*model.Comment {
	ml := make(map[int][]*model.Comment)
	if len(mComments) == 0 {
		return ml
	}
	for _, comments := range mComments {
		for _, comment := range comments {
			if user, ok := users[comment.UserID]; ok {
				ct := &model.Comment{
					Content: comment.Content,
					User:    user,
				}
				if comment.ReplyID > 0 { // 格式化回复者
					if u, o := users[comment.ReplyID]; o {
						ct.Reply = u
					}
				}
				if _, o := ml[comment.MomentID]; o {
					ml[comment.MomentID] = append(ml[comment.MomentID], ct)
				} else {
					ml[comment.MomentID] = []*model.Comment{ct}
				}
			}
		}
	}
	return ml
}
