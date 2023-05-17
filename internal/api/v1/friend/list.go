package friend

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// List 好友列表
// @Summary 好友列表
// @Description 好友列表
// @Tags 好友
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @success 0 {object} app.Response{data=[]model.User} "调用成功结构"
// @Router /friend/list [get]
func List(c *gin.Context) {
	list, err := service.Svc.FriendMyAll(c.Request.Context(), api.GetUserID(c))
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, list)
}

// TagList 标签好友列表
// @Summary 标签好友列表
// @Description 标签好友列表
// @Tags 好友
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param id query int true "标签ID"
// @success 0 {object} app.Response{data=[]model.User} "调用成功结构"
// @Router /friend/tag_list [get]
func TagList(c *gin.Context) {
	id := cast.ToInt(c.Query("id"))
	if id == 0 {
		app.Error(c, errno.ErrInvalidParam)
		return
	}

	list, err := service.Svc.FriendMyListByTagID(c.Request.Context(), api.GetUserID(c), id)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, list)
}
