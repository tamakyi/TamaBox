FROM golang:1.20-alpine as builder

WORKDIR /app

ENV CGO_ENABLED=0

ARG GITHUB_SHA=dev

COPY . .

RUN go mod tidy
RUN go build -v -ldflags "-w -s -extldflags '-static' -X 'github.com/tamakyi/TamaBox/internal/conf.BuildCommit=$GITHUB_SHA'" -o TamaBox ./cmd/

FROM alpine:latest

RUN apk update && apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Asia/Shanghai" > /etc/timezone

WORKDIR /home/app

COPY --from=builder /app/TamaBox .

RUN chmod 777 /home/app/TamaBox

ENTRYPOINT ["./TamaBox", "web"]
EXPOSE 8080
