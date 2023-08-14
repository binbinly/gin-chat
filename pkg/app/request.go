package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strings"
)

// RequestParams 获取请求参数
func RequestParams(c *gin.Context) map[string]any {
	params := map[string]any{}
	if c.Request.Method == "POST" {
		contextType := c.Request.Header.Get("Content-Type")
		if strings.Index(contextType, "json") >= 0 {
			if err := c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
				return nil
			}
		} else {
			_ = c.Request.ParseMultipartForm(32 << 20)
			if len(c.Request.PostForm) > 0 {
				for k, v := range c.Request.PostForm {
					params[k] = v[0]
				}
			}
		}
	} else {
		var tmpParams = make(map[string]string)
		if err := c.ShouldBind(&tmpParams); err != nil {
			return nil
		}
		for k, v := range tmpParams {
			params[k] = v
		}
	}
	params["method"] = c.Request.Method
	params["path"] = c.Request.URL.Path
	return params
}
