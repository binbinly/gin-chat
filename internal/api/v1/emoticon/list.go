package emoticon

import (
	"strings"

	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// List 表情资源列表
// @Summary 表情包
// @Description 表情包
// @Tags 表情包
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param cat query string true "分类"
// @success 0 {object} app.Response{data=[]model.Emoticon} "调用成功结构"
// @Router /emoticon/list [get]
func List(c *gin.Context) {
	cat := strings.TrimSpace(c.Query("cat"))
	if cat == "" {
		app.Error(c, errno.ErrValidation)
		return
	}
	list, err := service.Svc.Emoticon(c.Request.Context(), cat)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, list)
}
