package middleware

import (
	"net/http"

	"github.com/binbinly/pkg/errno"
	"github.com/binbinly/pkg/logger"
	"github.com/gin-gonic/gin"
)

// HandleErrors 异常捕获处理
func HandleErrors(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("[exception] err:%v", err)

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  errno.ErrInternalServer.Msg(),
			})
		}
	}()
	c.Next()
}
