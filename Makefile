# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

PROJECT_NAME := "gin-chat"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /examples)

all: test build
dev:
	$(GOBUILD) run main.go server

build:
	cp -r configs build/configs
	$(GOBUILD) -o build/gin-chat main.go

test:
	$(GOTEST) -v ${PKG_LIST}

clean:
	rm -rf build/gin-chat
	rm -rf nohup.out

# 运行服务
run:
	nohup build/gin-chat server -c configs &

# 停止服务
stop:
	pkill -f build/gin-chat

#生成docker镜像，请确保已安装docker
docker:
	docker build -t chat:latest -f Dockerfile .

# 生成api文档
doc:
	go install github.com/swaggo/swag/cmd/swag@latest
	@swag init
	echo "docs done"
	echo "see docs by: http://127.0.0.1:9050/swagger/index.html"

# 生成交互式的可视化Go程序调用图
graph:
	@echo "downloading go-callvis"
	@echo "generating graph"
	@go get -u github.com/ofabry/go-callvis
	@go-callvis ${PROJECT_NAME}

# 生成ca证书
ca:
	openssl req -new -nodes -x509 -out build/cert/server.crt -keyout build/cert/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"

#检查代码规范
lint:
	@go get -u golang.org/x/lint/golint
	@golint -set_exit_status ${PKG_LIST}

#查看帮助
help:
	target/chat --help
