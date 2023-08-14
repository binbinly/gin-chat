package collect

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// createParams 创建收藏
type createParams struct {
	Type    int8            `json:"type" binding:"required,oneof=2 3 4 5 6 7" example:"1"` // 聊天信息类型
	Content string          `json:"content" binding:"required" example:"test"`             // 内容
	Options json.RawMessage `json:"options" example:"test" swaggertype:"string"`           // 额外选项
}

// Create 添加收藏
// @Summary 添加收藏
// @Description 添加收藏
// @Tags 用户收藏
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param req body createParams true "create"
// @success 0 {object} app.Response "调用成功结构"
// @Router /collect/create [post]
func Create(c *gin.Context) {
	var req createParams
	if err := api.BindJSON(c, &req); err != nil {
		app.ErrorParamInvalid(c, err)
		return
	}

	err := service.Svc.CollectCreate(c.Request.Context(), req.Content, req.Options, api.GetUserID(c), req.Type)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}
