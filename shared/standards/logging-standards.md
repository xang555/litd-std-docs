# Logging Standards

## Overview

This document outlines the standards for logging across all teams.

## Principles

1. **Structured Logging**: Use machine-readable formats (JSON)
2. **Log Levels**: Use appropriate severity levels
3. **Context**: Include relevant context (request ID, user ID)
4. **Privacy**: Never log sensitive data

## Log Levels

| Level | Usage | Example |
|-------|-------|---------|
| `DEBUG` | Detailed diagnostic information | Variable values, execution flow |
| `INFO` | General informational messages | Service started, request completed |
| `WARN` | Unexpected but recoverable issues | Retry attempt, deprecated usage |
| `ERROR` | Error events that might still allow continued operation | API failure, DB connection error |
| `FATAL` | Critical events requiring immediate attention | Service shutdown, panic |

## Log Format

### JSON Format

```json
{
  "timestamp": "2024-01-01T12:00:00Z",
  "level": "info",
  "message": "User logged in",
  "context": {
    "user_id": "123",
    "request_id": "req_abc123",
    "ip": "192.168.1.1"
  }
}
```

### Structured Fields

- `timestamp` - ISO 8601 format
- `level` - Lowercase log level
- `message` - Human-readable message
- `context` - Additional structured data
- `error` - Error object (if applicable)

## What to Log

### Always Include

- Request ID (for distributed tracing)
- User ID (when applicable)
- Action/operation being performed
- Timestamp

### Selectively Include

- Input parameters (sanitize sensitive data)
- Output/results
- Duration of operations
- External service calls

## Never Log

- Passwords
- API keys/secrets
- Credit card numbers
- Personal health information
- Full request bodies (unless necessary and sanitized)

## Best Practices

- Use structured logging libraries (slog, zerolog, pino)
- Log at appropriate levels
- Include correlation IDs for distributed tracing
- Use semantic messages
- Avoid logging in hot paths (performance critical)
