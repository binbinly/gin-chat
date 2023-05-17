package group

import (
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// List 群组列表
// @Summary 群组列表
// @Description 群组列表
// @Tags 群组
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @success 0 {object} app.Response{data=[]model.GroupList} "调用成功结构"
// @Router /group/list [get]
func List(c *gin.Context) {
	list, err := service.Svc.GroupMyList(c.Request.Context(), api.GetUserID(c))
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, list)
}
