package collect

import (
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/resource"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// List 收藏列表
// @Summary 收藏列表
// @Description 收藏列表
// @Tags 用户收藏
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param p query int false "页码"
// @success 0 {object} app.Response{data=[]resource.CollectListResponse} "调用成功结构"
// @Router /collect/list [get]
func List(c *gin.Context) {
	offset, limit := api.GetPage(c)
	list, err := service.Svc.CollectGetList(c.Request.Context(), api.GetUserID(c), offset, limit)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, resource.CollectListResource(list))
}
