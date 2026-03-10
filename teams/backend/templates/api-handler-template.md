# API Handler Template

## Description

Standard Go HTTP handler template with proper error handling, validation, and response formatting.

## File Structure

```
handlers/
├── handler.go        # Main handler implementation
├── middleware.go     # Middleware functions
└── response.go       # Response utilities
```

## Handler Template

```go
// Package handlers provides HTTP request handlers for the API.
package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"log/slog"
)

// ============================================================================
// Types
// ============================================================================

// Request represents the incoming request structure
type Request struct {
	Field1 string `json:"field1" validate:"required"`
	Field2 int    `json:"field2" validate:"min=0,max=100"`
}

// Response represents the API response structure
type Response struct {
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// Handler dependencies
type Handler struct {
	logger *slog.Logger
	// Add other dependencies (database, services, etc.)
	db    Database
	cache Cache
}

// ============================================================================
// Constructor
// ============================================================================

// New creates a new handler instance
func New(logger *slog.Logger, db Database, cache Cache) *Handler {
	return &Handler{
		logger: logger,
		db:     db,
		cache:  cache,
	}
}

// ============================================================================
// HTTP Handler
// ============================================================================

// HandleRequest handles the API endpoint
// Method: POST
// Path: /api/resource
func (h *Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Validate request method
	if r.Method != http.MethodPost {
		h.writeErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Parse request body
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.ErrorContext(ctx, "failed to decode request", "error", err)
		h.writeErrorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Validate request
	if err := h.validateRequest(&req); err != nil {
		h.logger.WarnContext(ctx, "validation failed", "error", err)
		h.writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// Process request
	result, err := h.processRequest(ctx, &req)
	if err != nil {
		h.logger.ErrorContext(ctx, "processing failed", "error", err)
		h.writeErrorResponse(w, http.StatusInternalServerError, "processing failed")
		return
	}

	// Write response
	h.writeJSONResponse(w, http.StatusOK, Response{
		Status: "success",
		Data:   result,
	})
}

// ============================================================================
// Private Methods
// ============================================================================

// validateRequest validates the incoming request
func (h *Handler) validateRequest(req *Request) error {
	if req.Field1 == "" {
		return fmt.Errorf("field1 is required")
	}
	if req.Field2 < 0 || req.Field2 > 100 {
		return fmt.Errorf("field2 must be between 0 and 100")
	}
	return nil
}

// processRequest processes the business logic
func (h *Handler) processRequest(ctx context.Context, req *Request) (any, error) {
	// Check cache first
	if cached, found := h.cache.Get(ctx, req.Field1); found {
		return cached, nil
	}

	// Process business logic
	result, err := h.db.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	// Cache the result
	h.cache.Set(ctx, req.Field1, result, 5*time.Minute)

	return result, nil
}

// writeJSONResponse writes a JSON response
func (h *Handler) writeJSONResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.logger.Error("failed to encode response", "error", err)
	}
}

// writeErrorResponse writes an error response
func (h *Handler) writeErrorResponse(w http.ResponseWriter, status int, message string) {
	h.writeJSONResponse(w, status, Response{
		Status:  "error",
		Message: message,
	})
}
```

## Middleware Template

```go
// Package middleware provides HTTP middleware functions.
package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

// Logger logs HTTP requests
func Logger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Create response writer wrapper to capture status
			rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}

			// Call next handler
			next.ServeHTTP(rw, r)

			// Log request
			duration := time.Since(start)
			logger.InfoContext(r.Context(),
				"HTTP request",
				"method", r.Method,
				"path", r.URL.Path,
				"status", rw.status,
				"duration", duration,
				"ip", r.RemoteAddr,
			)
		})
	}
}

// CORS handles Cross-Origin Resource Sharing
func CORS(allowedOrigins []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			// Check if origin is allowed
			allowed := false
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin || allowedOrigin == "*" {
					allowed = true
					w.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}

			if allowed {
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				w.Header().Set("Access-Control-Max-Age", "86400")
			}

			// Handle preflight requests
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// Recovery recovers from panics
func Recovery(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rvr := recover(); rvr != nil {
					logger.ErrorContext(r.Context(),
						"panic recovered",
						"error", rvr,
						"stack", debug.Stack(),
					)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
```

## Response Utilities

```go
// Package response provides standardized HTTP response helpers.
package response

import (
	"encoding/json"
	"net/http"
)

// JSON writes a JSON response with status code
func JSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// Success writes a success response
func Success(w http.ResponseWriter, data any) error {
	return JSON(w, http.StatusOK, map[string]any{
		"status": "success",
		"data":   data,
	})
}

// Error writes an error response
func Error(w http.ResponseWriter, status int, message string) error {
	return JSON(w, status, map[string]any{
		"status":  "error",
		"message": message,
	})
}

// ValidationError writes a validation error response
func ValidationError(w http.ResponseWriter, errors map[string]string) error {
	return JSON(w, http.StatusBadRequest, map[string]any{
		"status":  "error",
		"message": "validation failed",
		"errors":  errors,
	})
}
```

## Tags

`backend`, `go`, `api`, `handler`, `http`, `template`
