# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run

PROJECT_NAME := "gin-chat"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /examples)

all: build run

.PHONY: dev
dev:
	$(GORUN) main.go server -c configs

.PHONY: build
build:
	mkdir -p build/configs
	cp -r configs/app.yaml build/configs/
	cp -r configs/database.yaml build/configs/
	cp -r configs/redis.yaml build/configs/
	@go build -o build/gin-chat main.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: clean
clean:
	rm -rf build/gin-chat
	rm -rf build/configs
	rm -rf nohup.out

# 运行服务
.PHONY: run
run:
	nohup build/gin-chat server -c configs &

# 初始化数据结构，并填充数据库表情包数据
.PHONY: init
init:
	$(GORUN) main.go migrate -d "root:root@tcp(127.0.0.1:3306)/chat"
	$(GORUN) main.go seed -d "root:root@tcp(127.0.0.1:3306)/chat"

# 停止服务
.PHONY: stop
stop:
	pkill -f build/gin-chat

#生成docker镜像，请确保已安装docker
.PHONY: docker
docker:
	docker build -t chat:latest -f Dockerfile .

# 生成api文档
.PHONY: doc
doc:
	go install github.com/swaggo/swag/cmd/swag@latest
	@swag init
	echo "docs done"
	echo "see docs by: http://127.0.0.1:9050/swagger/index.html"

# 生成交互式的可视化Go程序调用图
.PHONY: graph
graph:
	@echo "downloading go-callvis"
	@echo "generating graph"
	@go get -u github.com/ofabry/go-callvis
	@go-callvis ${PROJECT_NAME}

# 生成ca证书
.PHONY: ca
ca:
	openssl req -new -nodes -x509 -out build/cert/server.crt -keyout build/cert/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"

#检查代码规范
.PHONY: lint
lint:
	@go install golang.org/x/lint/golint@latest
	@golint -set_exit_status ${PKG_LIST}

#查看帮助
.PHONY: help
help:
	target/chat --help
