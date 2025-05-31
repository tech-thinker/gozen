# 🧭 Project Structure Documentation

This Go project was scaffolded using [`gozen`](https://github.com/tech-thinker/gozen). It follows a clean and modular architecture suited for microservices and REST/gRPC APIs.

## 🗂️ Root-Level Files and Directories

| Path                        | Description                                                             |
|-----------------------------|-------------------------------------------------------------------------|
| `main.go`                   | Entry point that initializes and runs the application.                  |
| `go.mod`, `go.sum`          | Go modules files for dependency management.                             |
| `Makefile`                  | Contains commands to build, run, and manage the project.                |
| `gozen.json`                | GoZen metadata and configuration file.                                  |
| `docker-compose*.yml`       | Docker Compose files for different environments (dev, debug, prod).     |
| `Dockerfile*`               | Docker configurations for building images in various contexts.          |
| `modd-*.conf`               | Modd configs for live reloading.                                        |
| `.env`, `.env.example`      | Environment variables configuration files.                              |
| `.gitignore`                | Git ignore file for project-specific files and directories.             |
| `.editorconfig`             | Editor configuration file for consistent coding styles.                 |

### Environment Variables
| Environment                 | Description                                                             |
|-----------------------------|-------------------------------------------------------------------------|
| VERSION                     | Application version.                                                    |
| APP_NAME                    | Application name.                                                       |
| APP_ENV                     | Environment name (development, staging, production).                    |
| API_PORT                    | Port number for the rest api to listen on.                              |
| GRPC_PORT                   | Port number for the gRPC api to listen on.                              |
| DB_DRIVER                   | Database driver (mysql, postgres, sqlite etc.).                         |
| DB_HOST                     | Database host address.                                                  |
| DB_PORT                     | Database port number.                                                   |
| DB_USER                     | Database username.                                                      |
| DB_PASS                     | Database password.                                                      |
| DB_NAME                     | Database name.                                                          |


## 📦 Application Structure

### `app/`
Main application logic divided into REST and gRPC interfaces.

#### `app/init.go`
- Main service registry for the application.


### `app/rest/`
Handles REST API endpoints.

| Path                        | Description                                     |
|-----------------------------|-------------------------------------------------|
| `controllers/health.go`     | Health check handler implementation (REST).     |
| `router/router.go`          | REST route registration.                        |


### `app/grpc/`
Handles gRPC endpoints and server setup.

| Path                            | Description                                      |
|---------------------------------|--------------------------------------------------|
| `handlers/health.go`            | Health check handler implementation (gRPC).     |
| `proto/health.proto`            | Protobuf definitions for health service.        |
| `proto/*.pb.go`                 | Generated Go code from `.proto`.                |
| `router/router.go`              | gRPC server and service registration.           |


## ⚙️ Configuration and Constants

### `config/config.go`
- Centralized configuration handling environments.

### `constants/app.go`
- Defines constants used throughout the application.


## 🧱 Core Domain Layers

### `models/health.go`
- Domain model for the health resource.

### `repository/health.go`
- Database or external service access logic for health module.

### `service/health.go`
- Business logic layer for the health service.


## 🧰 Supporting Infrastructure

### `instance/`

| Path                         | Description                                                      |
|------------------------------|------------------------------------------------------------------|
| `instance.go`                | Initializes shared instances (DB, cache, etc).                   |
| `registry/models.go`         | Structures for model registry or instance metadata.            |


### `logger/logger.go`
- Logger setup and utilities (e.g., wrapper for Zap or Logrus).


### `runner/`

| Path            | Description                        |
|-----------------|------------------------------------|
| `api.go`        | Initializes and runs REST server.  |
| `grpc.go`       | Initializes and runs gRPC server.  |


### `utils/utils.go`
- General-purpose utility functions.


## 🐳 Docker Support

| Path                   | Description                                       |
|------------------------|---------------------------------------------------|
| `Dockerfile*`          | Various Dockerfiles for dev/debug/production.     |
| `modd-*.conf`          | Live-reload configs using `modd` utility.         |


## ✅ Summary

This project structure promotes:

- **Separation of Concerns**: Clear division between transport, logic, and infrastructure.
- **Scalability**: Easy to add new modules or services.
- **Maintainability**: Decoupled layers and modular design.
