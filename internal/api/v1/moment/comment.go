package moment

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// commentParams 评论
type commentParams struct {
	ID      int    `json:"id" binding:"required,numeric" example:"1"`         // 动态ID
	ReplyID int    `json:"reply_id" binding:"omitempty,numeric" example:"1"`  // 回复者
	Content string `json:"content" binding:"required,max=500" example:"test"` // 内容
}

// Comment 评论
// @Summary 评论
// @Description 评论
// @Tags 朋友圈
// @Accept  json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body commentParams true "create"
// @success 0 {object} app.Response "调用成功结构"
// @Router /moment/comment [post]
func Comment(c *gin.Context) {
	var req commentParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	err := service.Svc.MomentComment(c.Request.Context(), api.GetUserID(c), req.ReplyID, req.ID, req.Content)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}
