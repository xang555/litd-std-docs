# Error Handling

## Overview

This document outlines the standards for error handling across all teams.

## Principles

1. **Fail Fast**: Detect and handle errors early
2. **Be Specific**: Provide clear, actionable error messages
3. **Log Context**: Include relevant information for debugging
4. **Don't Swallow**: Never silently ignore errors

## Error Types

### Client Errors (4xx)

- `400 Bad Request` - Invalid input
- `401 Unauthorized` - Missing or invalid credentials
- `403 Forbidden` - Insufficient permissions
- `404 Not Found` - Resource doesn't exist
- `409 Conflict` - Resource state conflict
- `422 Unprocessable Entity` - Validation errors
- `429 Too Many Requests` - Rate limit exceeded

### Server Errors (5xx)

- `500 Internal Server Error` - Unexpected error
- `502 Bad Gateway` - Upstream service error
- `503 Service Unavailable` - Service overloaded
- `504 Gateway Timeout` - Upstream timeout

## Response Format

```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "User-friendly message",
    "details": {
      "field": "email",
      "issue": "Invalid format"
    },
    "request_id": "req_abc123"
  }
}
```

## Error Codes

| Code | Description |
|------|-------------|
| `VALIDATION_ERROR` | Request validation failed |
| `AUTHENTICATION_ERROR` | Authentication failed |
| `AUTHORIZATION_ERROR` | Insufficient permissions |
| `NOT_FOUND` | Resource not found |
| `CONFLICT_ERROR` | Resource conflict |
| `RATE_LIMIT_ERROR` | Rate limit exceeded |
| `INTERNAL_ERROR` | Unexpected server error |

## Logging

- Log all errors with context
- Include request ID
- Sanitize sensitive data
- Use appropriate log levels
