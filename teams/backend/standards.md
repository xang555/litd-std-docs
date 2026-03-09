# Backend Team Standards

## Overview

This document outlines the coding standards, best practices, and conventions for the Backend Engineering team.

## Technology Stack

- **Language**: Go 1.21+
- **Web Framework**: Chi / Gin
- **Database**: PostgreSQL 15+
- **ORM**: sqlc / GORM
- **Testing**: testify
- **API Documentation**: OpenAPI 3.0

## Code Standards

### File Naming

- Packages: `lowercase` (e.g., `handlers`, `models`)
- Tests: `*_test.go` (e.g., `handlers_test.go`)
- Interfaces: `*.go` with clear naming
- Mocks: `mock_*.go` (generated)

### Package Structure

```
project/
├── cmd/           # Application entry points
├── internal/      # Private application code
│   ├── handlers/  # HTTP handlers
│   ├── services/  # Business logic
│   ├── models/    # Data models
│   └── repo/      # Database access
├── pkg/           # Public reusable packages
└── api/           # API definitions
```

### Go Best Practices

- Handle errors explicitly - don't ignore them
- Use contexts for all I/O operations
- Prefer interfaces at boundaries
- Keep goroutines constrained and observable
- Use structured logging with slog

## API Design

- RESTful resource naming
- Consistent response format
- Proper HTTP status codes
- Request validation
- Rate limiting

## Related Documents

- [API Design](./api-design.md)
- [Database Patterns](./database-patterns.md)
