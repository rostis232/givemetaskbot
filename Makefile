.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go
db-start:
	sudo docker start givemetaskbot
db-stop:
	sudo docker stop givemetaskbot
run: build
	./.bin/bot