.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t telegram-bot-go:v0.1 .

start-container:
	docker run --name telegram-bot -p 80:80 --env-file .env telegram-bot-go:v0.1 