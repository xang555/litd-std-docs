# Backend Team - Getting Started

## Welcome to the Backend Team

This guide helps you get started with backend development at our company.

## Tech Stack

- **Language**: Go
- **Framework**: Chi / Gin
- **Database**: PostgreSQL
- **Cache**: Redis
- **Message Queue**: RabbitMQ
- **Testing**: testify

## Setup

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Run database migrations:
   ```bash
   make migrate-up
   ```

3. Start development server:
   ```bash
   make run
   ```

4. Run tests:
   ```bash
   make test
   ```

## Project Structure

```
cmd/
├── api/            # API server
internal/
├── handlers/       # HTTP handlers
├── models/         # Data models
├── repository/     # Database access
├── services/       # Business logic
└── middleware/     # HTTP middleware
pkg/
├── config/         # Configuration
└── database/       # Database connection
```

## Code Style

Follow the coding standards defined in `teams/backend/standards/`.

## Resources

- Team-specific agents: `teams/backend/agents/`
- Team workflows: `teams/backend/workflows/`
- Shared resources: `shared/`
