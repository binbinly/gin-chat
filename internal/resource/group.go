package resource

import (
	"gin-chat/internal/model"
	"gin-chat/pkg/app"
)

// GroupResponse 群组响应结构
type GroupResponse struct {
	Nickname string        `json:"nickname"`
	Info     *Group        `json:"info"`
	Users    []*model.User `json:"users"`
}

// Group 群详情结构
type Group struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	InviteConfirm int8   `json:"invite_confirm"`
	Type          int8   `json:"type"`
	Name          string `json:"name"`
	Avatar        string `json:"avatar"`
	Remark        string `json:"remark"`
}

// GroupResource 群组信息转换
func GroupResource(group *model.GroupModel, users []*model.UserModel, gUsers []*model.GroupUserModel,
	my *model.GroupUserModel) *GroupResponse {
	if group.Type == model.GroupTypeRoom {
		return &GroupResponse{
			Info:  groupInfo(group),
			Users: make([]*model.User, 0),
		}
	}
	return &GroupResponse{
		Info:     groupInfo(group),
		Nickname: my.Nickname,
		Users:    GroupUsersResource(users, gUsers, 4),
	}
}

// GroupUsersResource 群成员列表转换
func GroupUsersResource(users []*model.UserModel, gUsers []*model.GroupUserModel, max int) []*model.User {
	us := make([]*model.User, 0)
	uMap := gUserToMap(gUsers)
	for i, user := range users {
		if max > 0 && i >= max { // 最多返回用户数
			break
		}
		name := user.Username
		if user.Nickname != "" { // 设置了昵称
			name = user.Nickname
		}
		if nick, ok := uMap[user.ID]; ok {
			name = nick
		}
		us = append(us, &model.User{
			ID:     user.ID,
			Name:   name,
			Avatar: app.BuildResUrl(user.Avatar),
		})
	}
	return us
}

// gUserToMap 转换群成员map结构 user_id => nickname
func gUserToMap(gUsers []*model.GroupUserModel) (m map[int]string) {
	m = make(map[int]string, len(gUsers))
	for _, user := range gUsers {
		if user.Nickname != "" {
			m[user.UserID] = user.Nickname
		}
	}
	return
}

func groupInfo(group *model.GroupModel) *Group {
	return &Group{
		ID:            group.ID,
		UserID:        group.UserID,
		InviteConfirm: group.InviteConfirm,
		Name:          group.Name,
		Avatar:        group.Avatar,
		Remark:        group.Remark,
		Type:          group.Type,
	}
}
