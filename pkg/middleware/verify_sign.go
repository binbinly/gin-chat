package middleware

import (
	"bytes"
	"io"
	"strconv"
	"strings"

	"gin-chat/pkg/app"

	"github.com/binbinly/pkg/errno"
	"github.com/binbinly/pkg/signature"
	"github.com/gin-gonic/gin"
)

// VerifySign 验证签名
func VerifySign(c *gin.Context) {
	if c.Request.URL.Path == "/v1/upload/file" {
		c.Next()
		return
	}

	// 签名信息
	authorization := c.GetHeader(app.HeaderSignToken)
	if authorization == "" {
		app.Error(c, errno.ErrSignParam)
		return
	}

	// 时间信息
	timestamp, _ := strconv.ParseInt(c.GetHeader(app.HeaderSignTokenDate), 10, 64)
	if timestamp == 0 {
		app.Error(c, errno.ErrSignParam)
		return
	}

	// 通过签名信息获取 key
	authorizationSplit := strings.Split(authorization, " ")
	if len(authorizationSplit) < 2 {
		app.Error(c, errno.ErrSignParam)
		return
	}
	key := authorizationSplit[0]

	data, _ := c.GetRawData()
	// 这里防止body只能读一次
	c.Request.Body = io.NopCloser(bytes.NewBuffer(data))
	params := app.RequestParams(c)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(data))

	st := signature.New(key, app.SignSecretKey, app.HeaderSignTokenTimeout)
	ok, err := st.Verify(authorization, timestamp, params)
	if err != nil || !ok {
		app.Error(c, errno.ErrSignParam)
		return
	}

	c.Next()
}
