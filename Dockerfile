FROM golang:1.15-alpine3.12 AS builder

COPY . /github.com/nikandpro/telegram-bot/
WORKDIR /github.com/nikandpro/telegram-bot/

RUN go mod download
RUN go build -o ./bin/bot/ cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/nikandpro/telegram-bot/bin/bot .
COPY --from=0 /github.com/nikandpro/telegram-bot/bin/configs configs/

EXPOSE 80

CMD ["./bot "]
