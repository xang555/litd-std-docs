# Backend Team Best Practices

## Overview
This document outlines the best practices for backend development at our company.

## Code Organization

### Package Structure

```
project/
├── cmd/                    # Main applications
├── internal/               # Private application code
│   ├── handlers/          # HTTP handlers
│   ├── models/            # Data models
│   ├── repository/        # Database access
│   ├── services/          # Business logic
│   └── middleware/        # HTTP middleware
└── pkg/                   # Public library code
```

### Import Organization

```go
import (
    // Standard library
    "context"
    "net/http"

    // External dependencies
    "github.com/gorilla/mux"

    // Internal packages
    "myapp/internal/models"
)
```

## Error Handling

### Error Wrapping

Always wrap errors with context:

```go
// Good
if err != nil {
    return fmt.Errorf("failed to create user: %w", err)
}

// Bad
if err != nil {
    return err
}
```

### Custom Error Types

Define custom errors for expected cases:

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
```

## Database Operations

### Connection Pooling

Configure appropriate pool sizes:

```go
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(5)
db.SetConnMaxLifetime(5 * time.Minute)
```

### Transaction Management

Always handle transactions properly:

```go
tx, err := db.BeginTx(ctx, nil)
if err != nil {
    return err
}
defer tx.Rollback() // Will rollback if not committed

// ... operations ...

if err = tx.Commit(); err != nil {
    return err
}
```

### N+1 Query Prevention

Use JOIN or preloading:

```go
// Bad: N+1 queries
users, _ := db.GetUsers(ctx)
for _, user := range users {
    orders, _ := db.GetOrdersByUser(ctx, user.ID) // N queries!
}

// Good: Single query with JOIN
usersWithOrders, _ := db.GetUsersWithOrders(ctx)
```

## API Design

### RESTful Conventions

```
GET    /users          # List users
GET    /users/{id}     # Get specific user
POST   /users          # Create user
PUT    /users/{id}     # Update user (replace)
PATCH  /users/{id}     # Update user (partial)
DELETE /users/{id}     # Delete user
```

### Response Format

Use consistent response format:

```go
type Response struct {
    Data    interface{} `json:"data,omitempty"`
    Error   *ErrorInfo  `json:"error,omitempty"`
    Meta    *MetaInfo   `json:"meta,omitempty"`
}

type ErrorInfo struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}
```

### Pagination

Use cursor-based pagination for large datasets:

```go
type PaginationParams struct {
    Limit  int    `json:"limit"`
    Cursor string `json:"cursor"`
}
```

## Context Usage

### Context Propagation

Always pass context through the call chain:

```go
func (s *Service) GetUser(ctx context.Context, id string) (*User, error) {
    return s.repo.GetByID(ctx, id)
}
```

### Context Timeouts

Set appropriate timeouts:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
```

## Concurrency

### Goroutine Management

Use wait groups for goroutine synchronization:

```go
var wg sync.WaitGroup
for _, item := range items {
    wg.Add(1)
    go func(i Item) {
        defer wg.Done()
        process(i)
    }(item)
}
wg.Wait()
```

### Channel Buffering

Buffer channels when appropriate:

```go
// Good: Buffered channel
ch := make(chan Result, 10)

// For unbuffered, ensure receiver is ready
ch := make(chan Result)
go func() { ch <- result }()
```

## Testing

### Table-Driven Tests

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        want     int
    }{
        {"positive", 1, 2, 3},
        {"negative", -1, -2, -3},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Add(tt.a, tt.b); got != tt.want {
                t.Errorf("Add() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Mock Interfaces

Use interfaces for testability:

```go
type Repository interface {
    GetByID(ctx context.Context, id string) (*User, error)
}

// In tests, use mock
type mockRepo struct {
    mock.Mock
}

func (m *mockRepo) GetByID(ctx context.Context, id string) (*User, error) {
    args := m.Called(ctx, id)
    return args.Get(0).(*User), args.Error(1)
}
```

## Security

### Input Validation

Always validate input:

```go
func CreateUser(name, email string) error {
    if name == "" {
        return errors.New("name is required")
    }
    if !isValidEmail(email) {
        return errors.New("invalid email format")
    }
    // ...
}
```

### SQL Injection Prevention

Use parameterized queries:

```go
// Good
db.Query("SELECT * FROM users WHERE id = ?", userID)

// Bad
db.Query("SELECT * FROM users WHERE id = " + userID)
```

### Secret Management

Never hardcode secrets:

```go
// Good
apiKey := os.Getenv("API_KEY")

// Bad
apiKey := "sk-1234567890"
```

## Performance

### Connection Reuse

Reuse HTTP clients:

```go
var httpClient = &http.Client{
    Timeout: 10 * time.Second,
}
```

### Memory Management

Be careful with goroutines:

```go
// Good: Limited concurrency
sem := make(chan struct{}, 10) // Max 10 concurrent
for _, item := range items {
    sem <- struct{}{}
    go func(i Item) {
        defer func() { <-sem }()
        process(i)
    }(item)
}
```

## Resources

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Uber Go Style Guide](https://github.com/uber-go/guide)
