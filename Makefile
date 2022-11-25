.SILENT:

build:
	go mod download && GOOS=linux go build -o ./app ./app/cmd/main.go

run: build
	IS_DEBUG=false ADMIN_LOGIN=admin ADMIN_PASS=admin ./app/main

swag: 
	swag init -g ./app/cmd/main.go -o ./docs