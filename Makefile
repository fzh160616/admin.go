APP_NAME=admin-server

.PHONY: run build test

run:
	go run ./cmd/server

build:
	go build -o ./bin/$(APP_NAME) ./cmd/server

test:
	go test ./...
