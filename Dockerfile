# 第一阶段：构建阶段
FROM golang:1.24 AS builder

WORKDIR /app

# 缓存模块依赖
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 🟢 关键：使用静态编译，避免动态依赖 GLIBC
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o dnslog_for_go

# 第二阶段：最小化运行环境
FROM debian:bullseye-slim

WORKDIR /app

# 只复制静态编译后的可执行文件
COPY --from=builder /app/dnslog_for_go .

EXPOSE 8080

CMD ["./dnslog_for_go"]


# 运行：docker-compose up