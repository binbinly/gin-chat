package apply

import (
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/resource"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// List 申请列表
// @Summary 我的申请列表
// @Description 我的申请列表
// @Tags 好友申请
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param p query int false "页码"
// @success 0 {object} app.Response{data=[]resource.ApplyListResponse} "调用成功结构"
// @Router /apply/list [get]
func List(c *gin.Context) {
	offset, limit := api.GetPage(c)
	list, users, err := service.Svc.ApplyMyList(c.Request.Context(), api.GetUserID(c), offset, limit)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, resource.ApplyListResource(list, users))
}

// Count 申请数量
// @Summary 待处理申请数量
// @Description 待处理申请数量
// @Tags 好友申请
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Success 200 {string} json "{"code":0,"message":"OK","data":1}"
// @Router /apply/count [get]
func Count(c *gin.Context) {
	count, err := service.Svc.ApplyPendingCount(c.Request.Context(), api.GetUserID(c))
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, count)
}
