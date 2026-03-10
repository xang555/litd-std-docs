# Go Conventions Standard

## Scope
This standard applies to all Go code in backend services and applications.

## Rules

### 1. Package Naming

**Rule:** Package names should be short, lowercase, single words.

**Rationale:** Short package names are easier to type and read. Avoid unnecessary repetition.

**Example:**
```go
// Good
package user
package auth
package httpx

// Bad
package users
package authentication
package httpUtilities
```

### 2. File Naming

**Rule:** Use lowercase, snake_case for file names.

**Rationale:** Consistent with Go community conventions and works across all file systems.

**Example:**
```
// Good
user_service.go
auth_handler.go
database.go

// Bad
userService.go
auth-handler.go
Database.go
```

### 3. Error Handling

**Rule:** Always handle errors explicitly. Never ignore returned errors.

**Rationale:** Explicit error handling prevents silent failures and makes debugging easier.

**Example:**
```go
// Good
result, err := someFunction()
if err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}

// Bad
result, _ := someFunction()
```

### 4. Interface Design

**Rule:** Accept interfaces, return structs.

**Rationale:** This makes your code flexible for consumers while strict in what you provide.

**Example:**
```go
// Good
func StoreData(w io.Writer, data []byte) error {
    // writes to any io.Writer
}

func NewRepository() *Repository {
    // returns concrete type
}

// Bad
func StoreData(f *os.File, data []byte) error {
    // only accepts files
}
```

### 5. Context Usage

**Rule:** Pass context.Context as the first parameter to functions that need it.

**Rationale:** Consistent with standard library and enables cancellation and timeout.

**Example:**
```go
// Good
func FetchUser(ctx context.Context, id string) (*User, error) {
    // ...
}

// Bad
func FetchUser(id string) (*User, error) {
    // ...
}
```

### 6. Struct Initialization

**Rule:** Use named fields when initializing structs.

**Rationale:** Makes code more readable and resilient to field reordering.

**Example:**
```go
// Good
user := User{
    Name:  "John",
    Email: "john@example.com",
}

// Bad
user := User{"John", "john@example.com"}
```

### 7. Variable Naming

**Rule:** Use camelCase for variables and exported names, camelCase for unexported.

**Rationale:** Follows Go conventions. Exported names start with capital letter.

**Example:**
```go
// Good
type UserService struct {}
var userCount int
func NewUserService() *UserService {}

// Bad
type userService struct {}
var user_count int
```

### 8. Comments

**Rule:** Exported functions must have godoc comments.

**Rationale:** Documentation is essential for maintainable code.

**Example:**
```go
// Good
// CreateUser creates a new user in the system.
// It returns the created user or an error if creation fails.
func CreateUser(name, email string) (*User, error) {
    // ...
}

// Bad
func CreateUser(name, email string) (*User, error) {
    // ...
}
```

### 9. Goroutine Leaks

**Rule:** Always ensure goroutines can exit. Use context cancellation for cleanup.

**Rationale:** Prevents goroutine leaks that exhaust resources.

**Example:**
```go
// Good
func process(ctx context.Context) {
    go func() {
        for {
            select {
            case <-ctx.Done():
                return // Exit on cancellation
            case <-time.After(time.Second):
                // Do work
            }
        }
    }()
}

// Bad
func process() {
    go func() {
        for {
            time.Sleep(time.Second)
            // Do work - never exits!
        }
    }()
}
```

### 10. Testing

**Rule:** Write table-driven tests for multiple test cases.

**Rationale:** Table-driven tests are concise and easy to extend.

**Example:**
```go
// Good
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 1, 2, 3},
        {"negative", -1, -2, -3},
        {"zero", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Add(tt.a, tt.b); got != tt.expected {
                t.Errorf("Add() = %v, want %v", got, tt.expected)
            }
        })
    }
}
```

## Enforcement

### Linting
- `golangci-lint` configuration: `.golangci.yml`
- Run: `make lint`

### Review
All pull requests must:
- Pass `go vet`
- Pass `golangci-lint`
- Have tests with >80% coverage
- Be reviewed by at least one team member

### Pre-commit Hooks
- `go fmt` runs automatically
- `go vet` runs before push
- Tests run before push

## Exceptions

Exceptions require:
- Team lead approval
- Documentation of rationale
- Code comment explaining deviation

## Related Standards
- Database Standards
- API Design Standards
- Security Standards

## Tags
`backend`, `go`, `standards`, `conventions`, `best-practices`
