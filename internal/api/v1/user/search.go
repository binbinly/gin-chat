package user

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/resource"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// searchParams 搜索
type searchParams struct {
	Keyword string `json:"keyword" binding:"required,max=18" example:"test"` //关键字
}

// Search 搜索用户
// @Summary 搜索用户
// @Description 搜索用户
// @Tags 用户
// @Accept json
// @Produce  json
// @Param Token header string true "用户令牌"
// @Param keyword body string true "搜索关键词"
// @success 0 {object} app.Response{data=[]resource.UserResponse} "调用成功结构"
// @Router /user/search [get]
func Search(c *gin.Context) {
	var req searchParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	list, err := service.Svc.UserSearch(c.Request.Context(), req.Keyword)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, resource.UserListResource(list))
}
