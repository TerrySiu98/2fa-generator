# 使用官方的 Go 镜像作为基础镜像
FROM golang:1.20-alpine AS builder
# 设置工作目录
WORKDIR /app
# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./
# 下载依赖
RUN go mod download
# 复制项目文件
COPY . .
# 编译项目
RUN go build -o 2fa-generator .
# 使用轻量级的 Alpine 镜像作为运行环境
FROM alpine:latest
# 设置工作目录
WORKDIR /app
# 从 builder 阶段复制编译好的二进制文件
COPY --from=builder /app/2fa-generator .
# 复制模板文件
COPY templates ./templates
# 暴露端口
EXPOSE 5656
# 运行应用
CMD ["./2fa-generator"]
