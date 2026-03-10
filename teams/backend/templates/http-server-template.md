# HTTP Server Template

## Description

Production-ready HTTP server template with graceful shutdown, middleware, and proper configuration.

## Main Server Template

```go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := Config{
		Port:         getEnv("PORT", "8080"),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Initialize logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// Create router
	router := mux.NewRouter()

	// Register middleware
	router.Use(middleware.Logger(logger))
	router.Use(middleware.Recovery(logger))
	router.Use(middleware.CORS([]string{"*"}))

	// Register handlers
	userHandler := handlers.NewUserHandler(logger, db)
	router.HandleFunc("/users", userHandler.List).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", userHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/users", userHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", userHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", userHandler.Delete).Methods(http.MethodDelete)

	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods(http.MethodGet)

	// Create server
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	// Start server in goroutine
	go func() {
		logger.Info("Starting server", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", "error", err)
		os.Exit(1)
	}

	logger.Info("Server exited gracefully")
}

// Config holds server configuration
type Config struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
```

## Route Group Template

```go
// SetupRoutes configures all application routes
func SetupRoutes(router *mux.Router, handlers *Handlers) {
	// API v1 routes
	v1 := router.PathPrefix("/api/v1").Subrouter()

	// User routes
	users := v1.PathPrefix("/users").Subrouter()
	users.HandleFunc("", handlers.User.List).Methods(http.MethodGet)
	users.HandleFunc("/{id}", handlers.User.Get).Methods(http.MethodGet)
	users.HandleFunc("", handlers.User.Create).Methods(http.MethodPost)
	users.HandleFunc("/{id}", handlers.User.Update).Methods(http.MethodPut)
	users.HandleFunc("/{id}", handlers.User.Delete).Methods(http.MethodDelete)

	// Auth routes (no authentication required)
	auth := v1.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", handlers.Auth.Login).Methods(http.MethodPost)
	auth.HandleFunc("/register", handlers.Auth.Register).Methods(http.MethodPost)
	auth.HandleFunc("/refresh", handlers.Auth.Refresh).Methods(http.MethodPost)

	// Protected routes (authentication required)
	protected := v1.PathPrefix("/protected").Subrouter()
	protected.Use(middleware.Auth())
	protected.HandleFunc("/profile", handlers.User.Profile).Methods(http.MethodGet)
}
```

## Context Middleware Template

```go
// Package middleware provides HTTP middleware for the application.
package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

type contextKey string

const (
	UserIDKey   contextKey = "userId"
	RequestIDKey contextKey = "requestId"
)

// RequestID adds a unique request ID to the context
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = generateID()
		}

		ctx := context.WithValue(r.Context(), RequestIDKey, requestID)
		w.Header().Set("X-Request-ID", requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Auth validates JWT tokens and adds user context
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Remove "Bearer " prefix
		token = strings.TrimPrefix(token, "Bearer ")

		// Validate token
		claims, err := validateToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add user ID to context
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequestTimeout adds a timeout to the request context
func RequestTimeout(timeout time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()

			done := make(chan struct{})
			go func() {
				defer close(done)
				next.ServeHTTP(w, r.WithContext(ctx))
			}()

			select {
			case <-done:
				return
			case <-ctx.Done():
				http.Error(w, "Request timeout", http.StatusGatewayTimeout)
			}
		})
	}
}
```

## Tags

`backend`, `go`, `server`, `http`, `template`, `middleware`
