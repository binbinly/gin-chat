package friend

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/ecode"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// blackParams 移入/移除黑名单
type blackParams struct {
	UserID int  `json:"user_id" binding:"required,numeric" example:"1"` // 用户ID
	Black  int8 `json:"black" binding:"oneof=0 1" example:"1"`          // 是否拉黑
}

// starParams 设置/取消星标好友
type starParams struct {
	UserID int  `json:"user_id" binding:"required,numeric" example:"1"` // 用户ID
	Star   int8 `json:"star" binding:"oneof=0 1" example:"1"`           // 是否星标用户
}

// authParams 设置朋友圈权限
type authParams struct {
	UserID  int  `json:"user_id" binding:"required,numeric" example:"1"` // 用户ID
	LookMe  int8 `json:"look_me"  binding:"oneof=0 1" example:"1"`       // 看我
	LookHim int8 `json:"look_him" binding:"oneof=0 1" example:"1"`       // 看他
}

// remarkParams 设置好友备注标签
type remarkParams struct {
	UserID   int      `json:"user_id" binding:"required,numeric" example:"1"`           // 用户ID
	Nickname string   `json:"nickname"  binding:"required,min=1,max=30" example:"test"` // 备注内侧
	Tags     []string `json:"tags" example:"test,test1"`                                // 标签
}

// Black 加入/移除黑名单
// @Summary 加入/移除黑名单
// @Description 加入/移除黑名单
// @Tags 好友
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body blackParams true "black"
// @success 0 {object} app.Response "调用成功结构"
// @Router /friend/black [post]
func Black(c *gin.Context) {
	var req blackParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	uid := api.GetUserID(c)
	if uid == req.UserID {
		app.Error(c, ecode.ErrUserNoSelf)
		return
	}
	err := service.Svc.FriendSetBlack(c.Request.Context(), uid, req.UserID, req.Black)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}

// Star 加入/移除星标
// @Summary 加入/移除星标
// @Description 加入/移除星标
// @Tags 好友
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body starParams true "star"
// @success 0 {object} app.Response "调用成功结构"
// @Router /friend/star [post]
func Star(c *gin.Context) {
	var req starParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	uid := api.GetUserID(c)
	if uid == req.UserID {
		app.Error(c, ecode.ErrUserNoSelf)
		return
	}
	err := service.Svc.FriendSetStar(c.Request.Context(), uid, req.UserID, req.Star)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}

// Remark 设置备注标签
// @Summary 设置备注标签
// @Description 设置备注标签
// @Tags 好友
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body remarkParams true "remark"
// @success 0 {object} app.Response "调用成功结构"
// @Router /friend/remark [post]
func Remark(c *gin.Context) {
	var req remarkParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	uid := api.GetUserID(c)
	if uid == req.UserID {
		app.Error(c, ecode.ErrUserNoSelf)
		return
	}
	err := service.Svc.FriendSetRemarkTag(c.Request.Context(), uid, req.UserID, req.Nickname, req.Tags)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}

// Auth 设置朋友圈权限
// @Summary 设置朋友圈权限
// @Description 设置朋友圈权限
// @Tags 好友
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body authParams true "auth"
// @success 0 {object} app.Response "调用成功结构"
// @Router /friend/auth [post]
func Auth(c *gin.Context) {
	var req authParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	uid := api.GetUserID(c)
	if uid == req.UserID {
		app.Error(c, ecode.ErrUserNoSelf)
		return
	}
	err := service.Svc.FriendSetMomentAuth(c.Request.Context(), uid, req.UserID, req.LookMe, req.LookHim)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}
