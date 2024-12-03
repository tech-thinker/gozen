# gozen
Simplified Golang MVC code generator.

Read project documentation for more information here: [Documentation](docs/README.md).

## Installation
Download and install executable binary from GitHub releases page.

### Linux Installation
```sh
curl -sL https://github.com/tech-thinker/gozen/releases/download/v0.1.1/gozen-linux-amd64 -o gozen
chmod +x gozen
sudo mv gozen /usr/bin
```

### MacOS Installation
```sh
curl -sL https://github.com/tech-thinker/gozen/releases/download/v0.1.1/gozen-darwin-amd64 -o gozen
chmod +x gozen
sudo mv gozen /usr/bin
```

- Or using homebrew
```sh
brew tap tech-thinker/tap
brew install gozen
```

### Windows Installation
```sh
curl -sL https://github.com/tech-thinker/gozen/releases/download/v0.1.1/gozen-windows-amd64.exe -o gozen.exe
gozen.exe
```

## CLI Guide
- `gozen` help
```sh
gozen -h
```

- Create new project
```sh
gozen create -d <database-driver> -p github.com/<username>/<appname> <appname>
```
Note: `database-driver` is one of `postgres`, `mysql`, `sqlite`.

- Chenge directory to new project directory
```sh
cd <appname>
```

- Install dependencies
```sh
go mod tidy
```

- Change `.env` file as per your requirements and export `.env`.
```sh
export $(cat .env | xargs)
```

- Run project
```sh
go run main.go start
```
