# Install dependencies

install:
	go mod tidy

test:
	go test -v ./...  -race -coverprofile=coverage.out -covermode=atomic

build:
	go build -o gozen

build-all:
	GOOS=linux GOARCH=amd64 go build -o build/gozen-linux-amd64
	GOOS=linux GOARCH=arm64 go build -o build/gozen-linux-arm64
	GOOS=linux GOARCH=arm go build -o build/gozen-linux-arm
	GOOS=darwin GOARCH=amd64 go build -o build/gozen-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -o build/gozen-darwin-arm64
	GOOS=windows GOARCH=amd64 go build -o build/gozen-windows-amd64.exe
	GOOS=windows GOARCH=386 go build -o build/gozen-windows-i386.exe
