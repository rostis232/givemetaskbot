FROM golang:1.19.6-alpine3.16 AS builder

COPY . /github.com/rostis232/givemetaskbot/
WORKDIR /github.com/rostis232/givemetaskbot/

RUN go mod download
RUN go build -o /bin/bot cmd/bot/main.go

FROM alpine:3.16

WORKDIR /root/

COPY --from=0 /github.com/rostis232/givemetaskbot/bin/bot .
COPY --from=0 /github.com/rostis232/givemetaskbot/configs configs/

EXPOSE 80

CMD ["./bot"]