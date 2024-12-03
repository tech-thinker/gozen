# Environments

test:
	go test -v ./...  -race -coverprofile=coverage.out -covermode=atomic

setup:
	go mod tidy

run:
	export $(cat .env | xargs) && go run main.go start

build:
	go build -o {{.AppName}}

clean:
	rm -rf {{.AppName}}

