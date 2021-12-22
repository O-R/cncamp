FROM golang:1.17.1-alpine3.14 AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/* && go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /httpserver

COPY . .
RUN go mod download && go build -o ./app .

FROM alpine:3.14

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
&& apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=builder /httpserver/app .

EXPOSE 8081
EXPOSE 8082

ENTRYPOINT ["./app"]
