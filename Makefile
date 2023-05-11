.DEFAULT_GOAL := run

run:
	go run ./cmd/api

build:
	set CGO_ENABLED=0
	go build -o target/greenlight.exe ./cmd/api
	set GOOS=linux
	go build -o target/greenlight ./cmd/api
lint:
	go fmt ./...
	go fix
