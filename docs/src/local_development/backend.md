---
icon: dot
---

# Backend

### Requirements

To set up your local development environment, you need to install the following tools:

- Git
- Docker
- Go (1.24.2)

!!!warning
If you want to work on the AI-Worker, you also need ollama installed.
!!!

### Databases

Almost all services need a database. In the `docker/e2e` directory, you can find a `docker-compose.yml` file that contains the configuration for all databases.
It's recommended to run the whole `docker-compose.yml` file, as it will set up all the databases you need for local development. 
Optionally you can also uncomment the Grafana-Stack if you want to use Grafana for monitoring.

### Run a service

All services are written in Go and can be run using the `go run` command.

To start a service, run the following commands:
```bash
cd services/<service>
go run cmd/e2e/main.go
```

### Run tests for a service

To run the tests for a service, run the following commands:
```bash
cd services/<service>
go test ./...
```