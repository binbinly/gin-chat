package resource

import (
	"gin-chat/internal/model"
)

// UserResponse 用户响应结构
type UserResponse struct {
	ID       int    `json:"id"`
	Phone    int64  `json:"phone"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Sign     string `json:"sign"`
	Area     string `json:"area"`
	Gender   int8   `json:"gender"`
}

// UserResource 用户信息转换
func UserResource(user *model.UserModel) *UserResponse {
	return &UserResponse{
		ID:       user.ID,
		Phone:    user.Phone,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Sign:     user.Sign,
		Area:     user.Area,
		Gender:   user.Gender,
	}
}

// UserBasicResource 用户基础信息转换
func UserBasicResource(user *model.UserModel) *model.User {
	return userBasic(user)
}

// UserListResource 用户列表转换
func UserListResource(users []*model.UserModel) []*UserResponse {
	if len(users) == 0 {
		return []*UserResponse{}
	}

	us := make([]*UserResponse, 0, len(users))
	for _, u := range users {
		us = append(us, UserResource(u))
	}
	return us
}

// usersToMap 用户列表转换成map结构
func usersToMap(users []*model.UserModel) (m map[int]*model.User) {
	m = make(map[int]*model.User, len(users))
	for _, user := range users {
		m[user.ID] = userBasic(user)
	}
	return m
}

// userBasic 用户基础信息
func userBasic(user *model.UserModel) *model.User {
	name := user.Username
	if user.Nickname != "" {
		name = user.Nickname
	}
	return &model.User{
		ID:     user.ID,
		Name:   name,
		Avatar: user.Avatar,
	}
}
