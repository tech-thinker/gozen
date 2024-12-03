# Environments

# Load .env file
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

test:
	go test -v ./...  -race -coverprofile=coverage.out -covermode=atomic

setup:
	go mod tidy

run:
	go run main.go start

build:
	go build -o {{.AppName}}

clean:
	rm -rf {{.AppName}}

