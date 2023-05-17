package group

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// idsParams 创建群组，用户id列表
type idsParams struct {
	Ids []int `json:"ids" binding:"gt=0,dive,required" example:"[1,2,3]" swaggertype:"string"` // 用户id列表
}

// Create 创建
// @Summary 创建群组
// @Description 创建群组
// @Tags 群组
// @Accept  json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param req body idsParams true "ids"
// @success 0 {object} app.Response "调用成功结构"
// @Router /group/create [post]
func Create(c *gin.Context) {
	var req idsParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	err := service.Svc.GroupCreate(c.Request.Context(), api.GetUserID(c), req.Ids)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}
