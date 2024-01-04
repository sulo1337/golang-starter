# Golang Simple HTTP Service Starter

This project is a starter I use for new Golang projects. This project roughly demonstrates implementation of clean architecture (WIP).

## Architecture

The code is structured into the following layers:
- **api** - layer for API routers, and handlers along with middlewares 
- **domain** - business entities
- **service** - business logic
- **dto** - data transfer models between layers
- **store** - repository layer for database/other service calls

## Features
- Routing with [chi router](https://github.com/go-chi/chi) 
- DB with [GORM](https://gorm.io)
- Dependency injection using [fx](go.uber.org/fx)
- Logging with [slog](https://pkg.go.dev/log/slog)
- Custom middlewares for logging, errors, etc

## WIP
- Unit testing
