package user

import (
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// Tag 标签列表
// @Summary 标签列表
// @Description 标签列表
// @Tags 用户
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @success 0 {object} app.Response{data=[]model.UserTag} "调用成功结构"
// @Router /user/tag [get]
func Tag(c *gin.Context) {
	list, err := service.Svc.UserTagAll(c.Request.Context(), api.GetUserID(c))
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, list)
}
