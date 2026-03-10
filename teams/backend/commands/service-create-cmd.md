# Service Create Command

## Description
Command template for scaffolding new Go microservices with best practices.

## Usage
```
/service-create <ServiceName> [options]
```

## Parameters

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| ServiceName | string | true | - | Name of the service (kebab-case) |
| --port | int | false | 8080 | Service port number |
| --with-db | boolean | false | true | Include database setup |
| --with-cache | boolean | false | true | include Redis cache |
| --with-queue | boolean | false | false | Include message queue |
| --with-metrics | boolean | false | true | Include Prometheus metrics |
| --path | string | false | services/ | Target directory path |

## Examples

### Basic service
```bash
/service-create user-service
```
Creates: `services/user-service/`

### Service with options
```bash
/service-create order-service --with-queue=true --port=8081
```
Creates: `services/order-service/` with queue support

### Service in specific location
```bash
/service-create payment-service --path=internal/services/
```
Creates: `internal/services/payment-service/`

## Generated Structure

```
ServiceName/
├── cmd/
│   └── service/
│       └── main.go              # Entry point
├── internal/
│   ├── config/                  # Configuration
│   │   └── config.go
│   ├── handlers/                # HTTP handlers
│   │   └── handler.go
│   ├── models/                  # Data models
│   │   └── model.go
│   ├── repository/              # Database access
│   │   └── repository.go
│   ├── services/                # Business logic
│   │   └── service.go
│   └── middleware/              # HTTP middleware
│       └── middleware.go
├── pkg/
│   ├── database/                # Database connection
│   │   └── database.go
│   └── server/                  # Server setup
│       └── server.go
├── api/                         # OpenAPI specs
│   └── openapi.yaml
├── go.mod
├── go.sum
├── Dockerfile
├── Makefile
└── README.md
```

## Main Template

The generated main.go follows this structure:

```go
package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"

    "{project}/internal/config"
    "{project}/internal/handlers"
    "{project}/pkg/server"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    srv := server.New(cfg, handlers.Register)

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Handle graceful shutdown
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        <-sigChan
        cancel()
    }()

    if err := srv.Start(ctx); err != nil {
        log.Fatalf("Server error: %v", err)
    }
}
```

## Tags
`backend`, `go`, `service`, `scaffold`, `microservice`, `cli`
