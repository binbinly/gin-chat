package moment

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// likeParams 点赞
type likeParams struct {
	ID int `json:"id" binding:"required,numeric" example:"1"` // 动态ID
}

// Like 点赞
// @Summary 点赞
// @Description 点赞
// @Tags 朋友圈
// @Accept  json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body likeParams true "like"
// @success 0 {object} app.Response "调用成功结构"
// @Router /moment/like [post]
func Like(c *gin.Context) {
	var req likeParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	err := service.Svc.MomentLike(c.Request.Context(), api.GetUserID(c), req.ID)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}
