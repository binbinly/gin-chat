package chat

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/internal/websocket"
	"gin-chat/pkg/app"
)

// sendParams 发送消息
type sendParams struct {
	ToID     int             `json:"to_id" binding:"required,numeric" example:"1"`          // 用户/群组ID
	ChatType int             `json:"chat_type" binding:"required,oneof=1 2" example:"1"`    // 聊天类型，1=用户，2=群组
	Type     int             `json:"type" binding:"required,oneof=2 3 4 5 6 7" example:"1"` // 聊天信息类型
	Content  string          `json:"content" binding:"required" example:"test"`             // 内容
	Options  json.RawMessage `json:"options" example:"test" swaggertype:"string"`           // 额外选项
}

// Send 发送消息
// @Summary 发送消息
// @Description 发送消息
// @Tags 聊天
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param req body sendParams true "send"
// @success 0 {object} app.Response{data=websocket.Chat} "调用成功结构"
// @Router /chat/send [post]
func Send(c *gin.Context) {
	var req sendParams
	if err := api.BindJSON(c, &req); err != nil {
		app.ErrorParamInvalid(c, err)
		return
	}

	var msg *websocket.Chat
	var err error
	if req.ChatType == typeUser {
		msg, err = service.Svc.ChatSendUser(c.Request.Context(), api.GetUserID(c), req.ToID, req.Type, req.Content, req.Options)
	} else {
		msg, err = service.Svc.ChatSendGroup(c.Request.Context(), api.GetUserID(c), req.ToID, req.Type, req.Content, req.Options)
	}
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, msg)
}
