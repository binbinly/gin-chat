package emoticon

import (
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// Cat 表情包所有分类
// @Summary 表情包
// @Description 表情包
// @Tags 表情包
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @success 0 {object} app.Response{data=[]model.Emoticon} "调用成功结构"
// @Router /emoticon/cat [get]
func Cat(c *gin.Context) {
	list, err := service.Svc.EmoticonCat(c.Request.Context())
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.Success(c, list)
}
