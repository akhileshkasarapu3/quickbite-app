# QUICK BITE

# QuickBite

QuickBite is a real-time food delivery and live tracking web application built to learn production-style backend and system design using Golang.

## Tech Stack
- Golang
- REST APIs
- WebSockets
- PostgreSQL
- Redis
- Docker
- GitHub Actions
- React frontend

## Project Goals
- Build restaurant, customer, and delivery workflows
- Support real-time order tracking
- Learn scalable backend architecture
- Practice production-style folder structure and coding standards

## Initial Structure
- `cmd/api` - application entrypoint
- `internal/config` - configuration
- `internal/handler` - HTTP handlers
- `internal/service` - business logic
- `internal/repository` - data access layer
- `internal/model` - domain models
- `docs` - project documentation
- `web` - frontend application

## Run
```bash
go run ./cmd/api