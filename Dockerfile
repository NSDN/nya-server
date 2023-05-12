FROM golang:1.19

WORKDIR /app

# 安装依赖
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# 打包程序
COPY . .
RUN go build -v -o forum-backend .

# 执行程序
CMD ["./forum-backend"]
