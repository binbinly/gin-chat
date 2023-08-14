package user

import (
	"github.com/binbinly/pkg/util/validator"
	"github.com/gin-gonic/gin"

	"gin-chat/internal/api"
	"gin-chat/internal/ecode"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// sendCodeParams 发送短信验证码
type sendCodeParams struct {
	Phone string `json:"phone" binding:"required" example:"13333333333"` //手机号
}

// SendCode 获取验证码
// @Summary 根据手机号获取校验码
// @Description 根据手机号获取校验码
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param req body sendCodeParams true "手机号"
// @success 0 {object} app.Response "调用成功结构"
// @Router /send_code [get]
func SendCode(c *gin.Context) {
	var req sendCodeParams
	if err := api.BindJSON(c, &req); err != nil {
		app.ErrorParamInvalid(c, err)
		return
	}

	if is := validator.RegexMatch(req.Phone, validator.ChineseMobileMatcher); !is {
		app.Error(c, ecode.ErrPhoneValid)
		return
	}
	code, err := service.Svc.SendSMS(c.Request.Context(), req.Phone)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	if code != "" { //测试时，直接返回验证码，方便调试
		app.Success(c, code)
		return
	}
	app.Success(c, nil)
}
