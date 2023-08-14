package app

import (
	"net/http"

	"github.com/binbinly/pkg/errno"
	"github.com/binbinly/pkg/util"
	"github.com/gin-gonic/gin"
)

// Response api的返回结构体
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// Success 成功返回
func Success(c *gin.Context, data any) {
	if data == nil {
		data = gin.H{}
	}

	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code: errno.Success.Code(),
		Msg:  errno.Success.Msg(),
		Data: data,
	})
}

// SuccessNil 成功返回，无数据
func SuccessNil(c *gin.Context) {
	Success(c, nil)
}

// Error 错误返回
func Error(c *gin.Context, err *errno.Error) {
	code, msg := errno.DecodeErr(err)
	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: gin.H{},
	})
}

// ErrorParamInvalid 参数错误
func ErrorParamInvalid(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code: errno.ErrInvalidParam.Code(),
		Msg:  err.Error(),
		Data: gin.H{},
	})
}

// RouteNotFound 未找到相关路由
func RouteNotFound(c *gin.Context) {
	c.String(http.StatusNotFound, "not found")
}

// healthCheckResponse 健康检查响应结构体
type healthCheckResponse struct {
	Status   string `json:"status"`
	Hostname string `json:"hostname"`
}

// HealthCheck will return OK if the underlying BoltDB is healthy.
// At least healthy enough for demoing purposes.
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, healthCheckResponse{Status: "UP", Hostname: util.Hostname()})
}
