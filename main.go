package main

import (
	"gin-chat/cmd"
)

// @title swagger 接口文档
// @version 2.0
// @description

// @contact.name binbinly
// @contact.url https://github.com/binbinly
// @contact.email binbinly@hotmail.com

// @license.name MIT
// @license.url

// @securityDefinitions.apikey  Token
// @in                          header
// @name                        Token

// @host http://127.0.0.1:9050
// @BasePath /v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	cmd.Execute()
}
