# Development Guide

## Project Directory Structure
```
.
├── Makefile
├── app
│   ├── controllers
│   │   └── health.go
│   ├── initializer
│   │   └── service.go
│   └── router
│       └── router.go
├── config
│   └── config.go
├── constants
│   └── app.go
├── go.mod
├── go.sum
├── gozen.json
├── instance
│   ├── instance.go
│   └── registry
│       └── models.go
├── logger
│   └── logger.go
├── main.go
├── models
│   └── health.go
├── repository
│   └── health.go
├── runner
│   └── api.go
├── service
│   └── health.go
└── utils
    └── utils.go
```

- `Makefile`: Makefile that will be used to build the project.
- `app`: This is the package that contains all the controllers, initializers and routers.
    - `controllers`: This is the package that contains all the controllers.
    - `initializer`: This is the package that contains all the initializers.
    - `router`: This is the package that contains all the routers.
- `config`: This is the package that contains the configuration for the application.
- `constants`: This is the package that contains the constants for the application.
- `go.mod`: This is the go.mod file that contains the dependencies for the application.
- `go.sum`: This is the go.sum file that contains the dependencies for the application.
- `gozen.json`: This is the gozen.json file that contains the configuration for the application. It will is being used by `gozen` cli tool.
- `instance`: This is the package that contains the instance for the application.
    - `instance.go`: This is the instance package that will be used to store object that required all over the application.
    - `registry`: This is the package that contains the registry for the application.
        - `models.go`: This is the registry package for models that will be used to register models for migration.
- `logger`: This is the package that contains the logger for the application.
- `main.go`: This is the main package that will be used to run the application.
- `models`: This is the package that contains the models for the application.
- `repository`: This is the package that contains the repository for the application.
- `runner`: This is the package that contains the runner for the application.
    - `api.go`: This is the runner package that will be used to run api.
- `service`: This is the package that contains the service, that will handle business logic of application.
- `utils`: This is the package that contains common utility functions.

### Request flow diagram
```
+-----------+                                +-------------------+
|           +------------------------------->|                   |
|   Router  |                                |    Controllers    |
|           |<-------------------------------+                   |
+-----------+                                +-------------+-----+
                                                     ^     |      
                                                     |     |      
                                                     |     |      
                                                     |     |      
                                                     |     v      
+------------------+                           +-----+-----------+
|                  +-------------------------->|                 |
|    Repository    |                           |     Service     |
|                  |<--------------------------+                 |
+------------------+                           +-----------------+
```
