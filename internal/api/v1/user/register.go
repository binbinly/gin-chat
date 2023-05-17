package user

import (
	"github.com/binbinly/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"gin-chat/internal/api"
	"gin-chat/internal/ecode"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"
)

// registerParams 注册
type registerParams struct {
	Phone           string `json:"phone" form:"phone" binding:"required" example:"13333333333"`                                   //手机号
	Username        string `json:"username" form:"username" binding:"required,min=4,max=18" example:"test"`                       //用户名
	Password        string `json:"password" form:"password" binding:"required,min=6,max=20" example:"123456"`                     //密码
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required,eqfield=Password" example:"123456"` //确认密码
}

// Register 注册
// @Summary 注册
// @Description 用户注册
// @Tags 用户
// @Accept json
// @Produce  json
// @Param req body registerParams true "register"
// @success 0 {object} app.Response "调用成功结构"
// @Router /reg [post]
func Register(c *gin.Context) {
	var req registerParams
	if v := api.BindJSON(c, &req); !v {
		app.Error(c, errno.ErrInvalidParam)
		return
	}
	is := api.ValidateMobile(req.Phone)
	phone := cast.ToInt64(req.Phone)
	if !is || phone == 0 {
		app.Error(c, ecode.ErrPhoneValid)
		return
	}
	_, err := service.Svc.UserRegister(c.Request.Context(), req.Username, req.Password, phone)
	if e := api.Error(err); e != nil {
		app.Error(c, e)
		return
	}
	app.SuccessNil(c)
}
