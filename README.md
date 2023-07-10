## 友情提示

**快速体验项目**：[在线访问地址](https://binbinly.github.io/chat-frontend)。

## 项目介绍

`gin-chat` 是一套仿微信ui的即时通讯全栈学习项目，项目UI出自 [uni-app实战仿微信app开发](https://study.163.com/course/introduction/1209487898.htm)

- 主要功能点如下
  ![功能点](/assets/img/app.png)
- 项目截图
  ![image](/assets/img/view.jpg)

## 📗 目录结构
- [project-layout](https://github.com/golang-standards/project-layout)

### 后端技术

- http框架使用 [Gin](https://github.com/gin-gonic/gin)
- websocket使用 [Websocket](https://github.com/gorilla/websocket)
- 数据库组件 [GORM](https://gorm.io) mysql连接
- redis组件 [go-redis](https://github.com/redis/go-redis) redis连接
- 命令行工具 [Cobra](https://github.com/spf13/cobra)
- 文档使用 [Swagger](https://swagger.io/) 生成
- 配置文件解析库 [Viper](https://github.com/spf13/viper)
- 使用 [JWT](https://jwt.io/) 进行身份鉴权认证
- 校验器使用 [validator](https://github.com/go-playground/validator)  也是 Gin 框架默认的校验器
- 包管理工具 [Go Modules](https://github.com/golang/go/wiki/Modules)
- 使用 [GolangCI-lint](https://golangci.com/) 进行代码检测
- 使用 make 来管理 Go 工程

### 前端技术
#### [代码入口](./web)

- 移动端 Vue 组件库 [vant](https://youzan.github.io/vant/#/zh-CN/)
- 脚手架 [vue-cli4 vant rem 移动端框架方案](https://github.com/sunniejs/vue-h5-template)
- 表情包 [ChineseBQB](https://github.com/zhaoolee/ChineseBQB)

### 开发环境

| 工具           | 版本号   | 下载                                         |
| ------------- |-------| ------------------------------------------ |
| golang        | 1.20  | https://golang.org/dl/                     |
| nodejs        | 18.15 | https://nodejs.org/zh-cn/download/         |
| mysql         | 5.7   | https://www.mysql.com/                     |
| redis         | 6.0   | https://redis.io/download                  |
| nginx         | 1.19  | http://nginx.org/en/download.html          |

### 项目部署

### 1,手动编译部署

> TIPS: 需要本地安装 MySQL and Redis 
```bash
# 下载安装
git clone https://github.com/binbinly/gin-chat

# 进入项目目录
cd gin-chat

# 编译
make build

# 修改 dbs redis 配置
cd build/configs

# 初始化数据结构，填充表情包
make init

# 运行
make run
```

### 2.Docker部署

[docker安装文档](https://docs.docker.com/engine/install/)
```shell
cd gin-chat
# 1. build image: 
docker build -t chat:latest -f Dockerfile .
# 2. start: 
docker run --rm -it -p 9050:9050 -p 9060:9060 chat:latest server
# 启动时设置 --rm 选项，这样在容器退出时就能够自动清理容器内部的文件系统
```

### 3. Railway部署 (✅推荐)

1. [![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/template/VTEHKI?referralCode=Aha8g3)
2. Set Deploy Start Command: /app/gin-chat server
> 前端打包后的代码部署在 [github-pages](https://github.com/binbinly/chat-frontend)

## 常用命令

- make help 查看帮助
- make init 初始化数据结构，并填充数据库表情包数据
- make build 编译项目
- make run 运行项目
- make test 运行测试用例
- make clean 清除编译文件
- make doc 生成接口文档
- make lint 代码检查
- make graph 生成交互式的可视化Go程序调用图
- make docker 生成docker镜像，确保已安装docker

## 📝 接口文档
> TIPS: 只有本地环境开启了文档
- [chat接口文档](http://127.0.0.1:9050/swagger/index.html)

## AI辅助
- [openai](https://platform.openai.com)
- [codeium](https://codeium.com/)
- [opencommit](https://github.com/di-sukharev/opencommit)
- [chatgpt-next](https://chat-gpt-next-web-weld-nine-63.vercel.app)

## storage
- redis: [upstash](https://console.upstash.com/)
- mysql: [planetscale](https://app.planetscale.com/)
- postgres: [supabase](https://supabase.com/), [neon](https://console.neon.tech/)

## 其他

- 编码规范: [Uber Go 语言编码规范](https://github.com/xxjwxc/uber_go_guide_cn)
