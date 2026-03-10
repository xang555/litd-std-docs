# Logging Pattern Template

## Description

Standardized logging patterns across different languages with structured logging support.

## Log Levels Reference

| Level | Description | Usage |
|-------|-------------|-------|
| TRACE | finest granularity | Detailed execution tracing |
| DEBUG | diagnostic information | Development troubleshooting |
| INFO | general information | Normal operation events |
| WARN | warning situations | Potentially harmful situations |
| ERROR | error events | Error events that might still allow continued operation |
| FATAL | critical errors | Severe errors that will likely lead to application termination |

## Structured Log Format

```json
{
  "timestamp": "2024-01-15T10:30:00Z",
  "level": "INFO",
  "message": "User logged in successfully",
  "context": {
    "userId": "12345",
    "sessionId": "abc-def-ghi",
    "ipAddress": "192.168.1.1",
    "userAgent": "Mozilla/5.0..."
  },
  "service": "auth-service",
  "environment": "production",
  "requestId": "req_abc123xyz"
}
```

## Go Logging (slog)

```go
package logging

import (
	"context"
	"io"
	"log/slog"
	"os"
)

// Logger wraps slog.Logger with additional functionality
type Logger struct {
	*slog.Logger
}

// Config holds logger configuration
type Config struct {
	Level      string
	Format     string // "json" or "text"
	Output     string // "stdout" or "stderr" or file path
	Service    string
	Environment string
}

// New creates a new logger instance
func New(cfg Config) *Logger {
	// Set log level
	var level slog.Level
	switch cfg.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	// Set output
	var output io.Writer
	switch cfg.Output {
	case "stderr":
		output = os.Stderr
	default:
		output = os.Stdout
	}

	// Set handler options
	opts := &slog.HandlerOptions{
		Level: level,
		AddSource: true,
	}

	// Create handler based on format
	var handler slog.Handler
	if cfg.Format == "json" {
		handler = slog.NewJSONHandler(output, opts)
	} else {
		handler = slog.NewTextHandler(output, opts)
	}

	// Create logger with default attributes
	logger := slog.New(handler)
	logger = logger.With(
		"service", cfg.Service,
		"environment", cfg.Environment,
	)

	return &Logger{logger}
}

// WithRequestID adds a request ID to the logger
func (l *Logger) WithRequestID(requestID string) *Logger {
	return &Logger{l.Logger.With("requestId", requestID)}
}

// WithUserID adds a user ID to the logger
func (l *Logger) WithUserID(userID string) *Logger {
	return &Logger{l.Logger.With("userId", userID)}
}

// WithContext adds context values to the logger
func (l *Logger) WithContext(ctx context.Context) *Logger {
	// Extract values from context (request ID, user ID, etc.)
	if requestID := ctx.Value("requestId"); requestID != nil {
		l = l.WithRequestID(requestID.(string))
	}
	return l
}

// Example usage:
/*
logger := logging.New(logging.Config{
    Level:      "info",
    Format:     "json",
    Service:    "my-service",
    Environment: "production",
})

logger.Info("User logged in",
    "userId", "12345",
    "sessionId", "abc-def",
    "ipAddress", "192.168.1.1",
)

logger.Error("Failed to process request",
    "error", err,
    "path", "/api/users",
)
*/
```

## Node.js/TypeScript Logging (pino)

```typescript
import pino from 'pino';

interface LoggerConfig {
  level: string;
  service: string;
  environment: string;
  prettyPrint?: boolean;
}

interface LogContext {
  requestId?: string;
  userId?: string;
  sessionId?: string;
  [key: string]: any;
}

export class Logger {
  private logger: pino.Logger;

  constructor(config: LoggerConfig) {
    this.logger = pino({
      level: config.level || 'info',
      formatters: {
        level: (label) => {
          return { level: label };
        },
      },
      base: {
        service: config.service,
        environment: config.environment,
      },
      ...(config.prettyPrint && {
        transport: {
          target: 'pino-pretty',
          options: {
            colorize: true,
            translateTime: 'HH:MM:ss Z',
            ignore: 'pid,hostname',
          },
        },
      }),
    });
  }

  withContext(context: LogContext): Logger {
    const child = new Logger({
      level: this.logger.level,
      service: '',
      environment: '',
    });
    child.logger = this.logger.child(context);
    return child;
  }

  info(message: string, obj?: any): void {
    this.logger.info(obj, message);
  }

  warn(message: string, obj?: any): void {
    this.logger.warn(obj, message);
  }

  error(message: string, error?: Error | any, obj?: any): void {
    if (error instanceof Error) {
      this.logger.error({ err: error, ...obj }, message);
    } else {
      this.logger.error(error || obj, message);
    }
  }

  debug(message: string, obj?: any): void {
    this.logger.debug(obj, message);
  }
}

// Example usage:
/*
const logger = new Logger({
  level: 'info',
  service: 'my-service',
  environment: 'production',
  prettyPrint: process.env.NODE_ENV !== 'production',
});

logger.info('User logged in', {
  userId: '12345',
  sessionId: 'abc-def',
  ipAddress: '192.168.1.1',
});

logger.withContext({ requestId: 'req_abc' }).info('Processing request');
*/
```

## Python Logging (structlog)

```python
import logging
import structlog
from typing import Any, Dict

def configure_logging(
    level: str = "INFO",
    service: str = "my-service",
    environment: str = "production",
    json_output: bool = True
) -> None:
    """Configure structured logging for the application."""

    # Configure standard logging
    logging.basicConfig(
        format="%(message)s",
        level=getattr(logging, level.upper()),
    )

    # Configure structlog
    processors = [
        structlog.contextvars.merge_contextvars,
        structlog.stdlib.add_log_level,
        structlog.stdlib.add_logger_name,
        structlog.processors.TimeStamper(fmt="iso"),
        structlog.processors.StackInfoRenderer(),
        structlog.processors.format_exc_info,
    ]

    if json_output:
        processors.append(structlog.processors.JSONRenderer())
    else:
        processors.append(structlog.dev.ConsoleRenderer())

    structlog.configure(
        processors=processors,
        wrapper_class=structlog.stdlib.BoundLogger,
        context_class=dict,
        logger_factory=structlog.stdlib.LoggerFactory(),
        cache_logger_on_first_use=True,
    )

    # Add default context
    structlog.configure(
        processors=processors + [
            structlog.processors.CallsiteParameterAdder(
                [
                    structlog.processors.CallsiteParameter.FILENAME,
                    structlog.processors.CallsiteParameter.LINENO,
                ]
            ),
        ],
    )

def get_logger(**kwargs) -> structlog.stdlib.BoundLogger:
    """Get a logger instance with optional context."""
    return structlog.get_logger(**kwargs)

# Example usage:
"""
configure_logging(
    level="INFO",
    service="auth-service",
    environment="production"
)

logger = get_logger()

logger.info("User logged in", user_id="12345", session_id="abc-def")

# With context binding
logger = logger.bind(request_id="req_abc123", user_id="12345")
logger.info("Processing request", path="/api/users")

logger.error("Failed to connect to database", error=str(e), host="db.example.com")
"""
```

## Common Logging Patterns

### Request Logging
```json
{
  "level": "INFO",
  "message": "Incoming request",
  "method": "GET",
  "path": "/api/users/123",
  "statusCode": 200,
  "duration": 125,
  "requestId": "req_abc123",
  "userId": "12345",
  "ipAddress": "192.168.1.1"
}
```

### Error Logging
```json
{
  "level": "ERROR",
  "message": "Database connection failed",
  "error": "connection timeout",
  "stackTrace": "...",
  "host": "db.example.com",
  "port": 5432,
  "retryAttempt": 3,
  "requestId": "req_abc123"
}
```

### Business Event Logging
```json
{
  "level": "INFO",
  "message": "Payment processed",
  "eventType": "payment.completed",
  "paymentId": "pay_abc123",
  "amount": 99.99,
  "currency": "USD",
  "userId": "12345",
  "timestamp": "2024-01-15T10:30:00Z"
}
```

## Tags

`logging`, `monitoring`, `debugging`, `observability`, `template`, `universal`
