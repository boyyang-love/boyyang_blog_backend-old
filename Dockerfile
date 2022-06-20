FROM golang:1.17 as builder

# ENV GO111MODULE=on \
#     GOPROXY=https://goproxy.cn,direct \
#     CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN GO111MODULE=on GOPROXY=https://goproxy.cn,direct CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# RUN go env -w GOPROXY=https://goproxy.io

# 打包
# RUN go build -o main ./main.go

# 端口号
# EXPOSE 80

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/config.yaml .

CMD ["/app/main"]


