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
}

// ============================================================================
// Constructor
// ============================================================================

// New creates a new handler instance
func New(logger *slog.Logger) *Handler {
	return &Handler{
		logger: logger,
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
	// Add validation logic
	return nil
}

// processRequest processes the business logic
func (h *Handler) processRequest(ctx context.Context, req *Request) (any, error) {
	// Add business logic
	return map[string]any{}, nil
}

// writeJSONResponse writes a JSON response
func (h *Handler) writeJSONResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// writeErrorResponse writes an error response
func (h *Handler) writeErrorResponse(w http.ResponseWriter, status int, message string) {
	h.writeJSONResponse(w, status, Response{
		Status:  "error",
		Message: message,
	})
}
