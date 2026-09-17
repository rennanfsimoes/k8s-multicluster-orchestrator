.PHONY: build test

build:
	go build -o bin/manager cmd/manager/main.go

test:
	go test ./...
