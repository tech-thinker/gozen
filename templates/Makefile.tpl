# Environments
include .env
export $(shell dotenv export)

# Install dependencies
install:
	go mod tidy

run:
	go run main.go start

build:
	go build -o {{.AppName}}
