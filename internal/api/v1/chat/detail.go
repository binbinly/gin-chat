package chat

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/internal/websocket"
	"gin-chat/pkg/app"
)

const (
	typeUser = iota + 1
	typeGroup
)

// detailParams 聊天详情
type detailParams struct {
	ID   int `json:"id" binding:"required,numeric" example:"1"`     // 用户/群组ID
	Type int `json:"type" binding:"required,oneof=1 2" example:"1"` // 聊天类型，1=用户，2=群组
}

// Detail 获取聊天信息
// @Summary 获取聊天信息
// @Description 获取聊天信息
// @Tags 聊天
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body detailParams true "chat detail"
// @success 0 {object} app.Response{data=websocket.Sender} "调用成功结构"
// @Router /chat/detail [post]
func Detail(c *gin.Context) {
	var req detailParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	var info *websocket.Sender
	var err error
	if req.Type == typeUser {
		info, err = service.Svc.ChatUserDetail(c.Request.Context(), api.GetUserID(c), req.ID)
	} else {
		info, err = service.Svc.ChatGroupDetail(c.Request.Context(), api.GetUserID(c), req.ID)
	}
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, info)
}
