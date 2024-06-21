# Install dependencies
install:
	go mod tidy

test:
	go test -v ./...  -race -coverprofile=coverage.out -covermode=atomic

build:
	go build -o gozen
