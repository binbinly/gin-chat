package collect

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// destroyParams 删除收藏
type destroyParams struct {
	ID int `json:"id" binding:"required,numeric" example:"1"` // 收藏id
}

// Destroy 删除收藏
// @Summary 删除收藏
// @Description 删除收藏
// @Tags 用户收藏
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param req body destroyParams true "destroy"
// @success 0 {object} app.Response "调用成功结构"
// @Router /collect/destroy [post]
func Destroy(c *gin.Context) {
	var req destroyParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	err := service.Svc.CollectDestroy(c.Request.Context(), api.GetUserID(c), req.ID)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}
