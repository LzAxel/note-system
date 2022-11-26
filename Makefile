.SILENT:

build:
	go mod download && GOOS=linux go build -o ./app/bin ./app/cmd/main.go

run: build
	IS_DEBUG=true ADMIN_LOGIN=admin ADMIN_PASS=admin \
	DB_HOST=localhost DB_PORT=5432 DB_USERNAME=postgres \
	DB_PASS=admin DB_NAME=postgres DB_SSLMODE=disable \
	JWT_SECRET=wRs0TXItEMUZcNanU6m9109SOfUm2I1P25YKKJsUbV8esBYGXUHELSwSbrFF \
	LOG_LEVEL=debug \
	./app/bin/main

swag: 
	swag init -g ./app/cmd/main.go -o ./docs
migrate:
	migrate -path ./schema -database postgres://postgres:admin@localhost:5432/postgres?sslmode=disable up

migrate-down:
	migrate -path ./schema -database postgres://postgres:admin@localhost:5432/postgres?sslmode=disable down