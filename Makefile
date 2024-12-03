VERSION := $(or $(AppVersion), "v0.0.0")
COMMIT := $(or $(shell git rev-parse --short HEAD), "unknown")
BUILDDATE := $(shell date +%Y-%m-%d)

LDFLAGS := -X 'main.AppVersion=$(VERSION)' -X 'main.CommitHash=$(COMMIT)' -X 'main.BuildDate=$(BUILDDATE)'

all: build

setup:
	go mod tidy

test:
	go test -v ./...  -race -coverprofile=coverage.out -covermode=atomic

coverage: test
	go tool cover -func=coverage.out

coverage-html: test
	mkdir -p coverage
	go tool cover -html=coverage.out -o coverage/index.html

coverage-serve: coverage-html
	python3 -m http.server 8080 -d coverage

install: build
	cp gozen /usr/local/bin/gozen
	cp man/gozen.1 /usr/local/share/man/man1/gozen.1

uninstall:
	rm /usr/local/bin/gozen
	rm /usr/local/share/man/man1/gozen.1

build:
	go build -ldflags="$(LDFLAGS)" -o gozen

dist:
	cp man/gozen.1 man/gozen.old
	sed -e "s|BUILDDATE|$(BUILDDATE)|g" -e "s|VERSION|$(VERSION)|g" man/gozen.old > man/gozen.1

	GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o build/gozen-linux-amd64
	cp build/gozen-linux-amd64 build/gozen
	tar -zcvf build/gozen-linux-amd64.tar.gz build/gozen man/gozen.1

	GOOS=linux GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o build/gozen-linux-arm64
	cp build/gozen-linux-arm64 build/gozen
	tar -zcvf build/gozen-linux-arm64.tar.gz build/gozen man/gozen.1

	GOOS=linux GOARCH=arm go build -ldflags="$(LDFLAGS)" -o build/gozen-linux-arm
	cp build/gozen-linux-arm build/gozen
	tar -zcvf build/gozen-linux-arm.tar.gz build/gozen man/gozen.1
	
	GOOS=darwin GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o build/gozen-darwin-amd64
	cp build/gozen-darwin-amd64 build/gozen
	tar -zcvf build/gozen-darwin-amd64.tar.gz build/gozen man/gozen.1
	
	GOOS=darwin GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o build/gozen-darwin-arm64
	cp build/gozen-darwin-arm64 build/gozen
	tar -zcvf build/gozen-darwin-arm64.tar.gz build/gozen man/gozen.1
	rm build/gozen
	
	GOOS=windows GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o build/gozen-windows-amd64.exe
	GOOS=windows GOARCH=386 go build -ldflags="$(LDFLAGS)" -o build/gozen-windows-i386.exe
	
	# Generating checksum
	cd build && sha256sum * >> checksum-sha256sum.txt
	cd build && md5sum * >> checksum-md5sum.txt

	# Cleaning
	mv man/gozen.old man/gozen.1

clean:
	rm -rf gozen*
	rm -rf build
