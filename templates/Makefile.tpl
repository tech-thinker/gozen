# Environments
include .env
export $(shell dotenv export)

test:
    go test -v ./...  -race -coverprofile=coverage.out -covermode=atomic

install:
	go mod tidy

run:
	go run main.go start

build:
	go build -o {{.AppName}}
