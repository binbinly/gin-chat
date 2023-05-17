package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"github.com/binbinly/pkg/errno"
	"github.com/binbinly/pkg/logger"
	"github.com/binbinly/pkg/util"
	"github.com/gin-gonic/gin"

	"gin-chat/pkg/app"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logging is a middleware function that logs the request.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		// Read the Body content
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
		}

		// Restore the io.ReadCloser to its original state
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		method := c.Request.Method
		ip := c.ClientIP()

		//log.Debugf("New request come in, path: %s, Method: %s, body `%s`", path, method, string(bodyBytes))
		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw

		// Continue.
		c.Next()

		// Calculates the latency.
		end := time.Now().UTC()
		latency := end.Sub(start)

		var code int
		var message string

		// get code and message
		var response app.Response
		if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
			logger.Warnf("response body can not unmarshal to model.Response struct, body: `%s`, err: %+v",
				blw.body.Bytes(), err)
			code = errno.ErrInternalServer.Code()
			message = err.Error()
		} else {
			code = response.Code
			message = response.Msg
		}
		if code != errno.Success.Code() {
			logger.Infof("%-13s | %-12s | %s %s | {code: %d, message: %s}", latency, ip,
				util.RightPad(method, " ", 5), path, code, message)
		}
	}
}
