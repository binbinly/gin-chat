# build stage
# 拉取最小体积的node环境容器, 并安装cnpm (加快容器构建速度,npm比较慢)
FROM node:lts-alpine as build-stage

# 进入容器的app目录, 然后拷贝当前目录(根目录)的所有文件到容器的当前目录中(/app)
WORKDIR /app
COPY . .

ENV VUE_APP_HTTP_URL='http://127.0.0.1:9050/v1' \
  VUE_APP_WS_URL='ws://127.0.0.1:9070/ws'

#在容器内build  --unsafe-perm 用高权限用户运行npm
RUN npm i --unsafe-perm --allow-root && npm run build

# production stage
# 最后通过nginx部署build出来的文件(/dist)
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
