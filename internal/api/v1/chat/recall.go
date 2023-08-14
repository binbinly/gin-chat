package chat

import (
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// recallParams 撤回消息
type recallParams struct {
	ID       string `json:"id" binding:"required,max=20" example:"1111"`        // 消息id
	ToID     int    `json:"to_id" binding:"required,numeric" example:"1"`       // 用户/群组ID
	ChatType int    `json:"chat_type" binding:"required,oneof=1 2" example:"1"` // 聊天类型，1=用户，2=群组
}

// Recall 消息撤回
// @Summary 消息撤回
// @Description 消息撤回
// @Tags 聊天
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param req body recallParams true "recall"
// @success 0 {object} app.Response "调用成功结构"
// @Router /chat/recall [post]
func Recall(c *gin.Context) {
	var req recallParams
	if err := api.BindJSON(c, &req); err != nil {
		app.ErrorParamInvalid(c, err)
		return
	}

	var err error
	if req.ChatType == typeUser {
		err = service.Svc.ChatUserRecall(c.Request.Context(), api.GetUserID(c), req.ToID, req.ID)
	} else {
		err = service.Svc.ChatGroupRecall(c.Request.Context(), api.GetUserID(c), req.ToID, req.ID)
	}
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}
