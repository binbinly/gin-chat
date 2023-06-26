package middleware

import (
	"gin-chat/internal/api"
	"gin-chat/internal/service"
	"gin-chat/pkg/app"

	"github.com/binbinly/pkg/logger"
	"github.com/gin-gonic/gin"
)

// JWT 认证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := service.Svc.UserTokenCheck(c, c.Request.Header.Get("Token"))
		if err != nil {
			app.Error(c, api.Error(err))
			return
		}
		logger.Debugf("context uid is: %v", uid)

		// set uid to context
		c.Set("uid", uid)

		c.Next()
	}
}
