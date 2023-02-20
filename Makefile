.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go
db-start:
	sudo docker start givemetaskbot
db-stop:
	sudo docker stop givemetaskbot
migrate-up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
migrate-down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down
run: build
	./.bin/bot