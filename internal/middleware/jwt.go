package middleware

import (
	"errors"

	"github.com/binbinly/pkg/errno"
	"github.com/binbinly/pkg/logger"
	"github.com/gin-gonic/gin"

	"gin-chat/pkg/app"
	"gin-chat/pkg/auth"
)

// JWT 认证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		payload, err := parse(c, app.Conf.JwtSecret)
		if err != nil {
			app.Error(c, errno.ErrInvalidToken)
			return
		}
		logger.Debugf("context is: %+v", payload)

		// set uid to context
		c.Set("uid", payload.UserID)

		c.Next()
	}
}

func parse(c *gin.Context, secret string) (*auth.Payload, error) {
	token := c.Request.Header.Get("Token")

	if len(token) == 0 {
		return nil, errors.New("the length of the `Authorization` header is zero")
	}
	return auth.Parse(token, secret)
}
