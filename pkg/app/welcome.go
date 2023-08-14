package app

import (
	"context"
	"encoding/json"

	"gin-chat/pkg/client/http"
)

// 在线聊天室进入欢迎语
const defText = "welcome"

type result struct {
	Code int  `json:"code"`
	Data data `json:"data"`
}

type data struct {
	Message string `json:"message"`
}

// WelcomeText 这里调用心灵鸡汤文为欢迎语 see: https://collect.xmwxxc.com/index/doc/sign/djt.html
func WelcomeText() string {
	url := "https://collect.xmwxxc.com/collect/djt/?type=0"
	client := http.NewRestyClient()
	rsp, err := client.Get(context.Background(), url)
	if err != nil {
		return defText
	}
	var rs result
	if err = json.Unmarshal(rsp, &rs); err != nil {
		return defText
	}
	if rs.Data.Message == "" {
		return defText
	}

	return rs.Data.Message
}
