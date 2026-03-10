# API Response Format Template

## Description

Standardized API response format for RESTful APIs with success, error, and pagination handling.

## Success Response Format

```json
{
  "success": true,
  "status": "success",
  "message": "Operation completed successfully",
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "Example Resource",
    "createdAt": "2024-01-15T10:30:00Z",
    "updatedAt": "2024-01-15T10:30:00Z"
  },
  "metadata": {
    "timestamp": "2024-01-15T10:30:00Z",
    "requestId": "req_abc123xyz"
  }
}
```

## Error Response Format

```json
{
  "success": false,
  "status": "error",
  "message": "Validation failed",
  "error": {
    "code": "VALIDATION_ERROR",
    "type": "ValidationError",
    "details": [
      {
        "field": "email",
        "message": "Email is required",
        "code": "REQUIRED_FIELD"
      },
      {
        "field": "password",
        "message": "Password must be at least 8 characters",
        "code": "MIN_LENGTH"
      }
    ]
  },
  "metadata": {
    "timestamp": "2024-01-15T10:30:00Z",
    "requestId": "req_abc123xyz",
    "path": "/api/v1/users"
  }
}
```

## Paginated Response Format

```json
{
  "success": true,
  "status": "success",
  "data": [
    {
      "id": "1",
      "name": "Item 1"
    },
    {
      "id": "2",
      "name": "Item 2"
    }
  ],
  "pagination": {
    "page": 1,
    "pageSize": 20,
    "totalItems": 100,
    "totalPages": 5,
    "hasNext": true,
    "hasPrevious": false
  },
  "metadata": {
    "timestamp": "2024-01-15T10:30:00Z"
  }
}
```

## HTTP Status Code Reference

### Success Codes
| Code | Description | Usage |
|------|-------------|-------|
| 200 | OK | Successful GET, PUT, PATCH |
| 201 | Created | Successful POST |
| 204 | No Content | Successful DELETE |

### Client Error Codes
| Code | Description | Usage |
|------|-------------|-------|
| 400 | Bad Request | Invalid request format |
| 401 | Unauthorized | Missing or invalid authentication |
| 403 | Forbidden | Valid auth but insufficient permissions |
| 404 | Not Found | Resource doesn't exist |
| 409 | Conflict | Resource already exists |
| 422 | Unprocessable Entity | Validation errors |
| 429 | Too Many Requests | Rate limit exceeded |

### Server Error Codes
| Code | Description | Usage |
|------|-------------|-------|
| 500 | Internal Server Error | Unexpected server error |
| 502 | Bad Gateway | Upstream service error |
| 503 | Service Unavailable | Service temporarily down |

## Implementation: Go

```go
package response

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Success  bool        `json:"success"`
	Status   string      `json:"status"`
	Message  string      `json:"message,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Error    *ErrorInfo  `json:"error,omitempty"`
	Metadata *Metadata   `json:"metadata,omitempty"`
}

type ErrorInfo struct {
	Code    string            `json:"code"`
	Type    string            `json:"type"`
	Details []ErrorDetail     `json:"details,omitempty"`
}

type ErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}

type Metadata struct {
	Timestamp  string `json:"timestamp"`
	RequestID  string `json:"requestId,omitempty"`
	Path       string `json:"path,omitempty"`
}

type Pagination struct {
	Page       int  `json:"page"`
	PageSize   int  `json:"pageSize"`
	TotalItems int  `json:"totalItems"`
	TotalPages int  `json:"totalPages"`
	HasNext    bool `json:"hasNext"`
	HasPrev    bool `json:"hasPrevious"`
}

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Success(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusOK, Response{
		Success: true,
		Status:  "success",
		Data:    data,
		Metadata: &Metadata{
			Timestamp: time.Now().Format(time.RFC3339),
		},
	})
}

func Created(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusCreated, Response{
		Success: true,
		Status:  "success",
		Message: "Resource created successfully",
		Data:    data,
		Metadata: &Metadata{
			Timestamp: time.Now().Format(time.RFC3339),
		},
	})
}

func Error(w http.ResponseWriter, status int, code string, message string) {
	JSON(w, status, Response{
		Success: false,
		Status:  "error",
		Message: message,
		Error: &ErrorInfo{
			Code: code,
			Type: http.StatusText(status),
		},
		Metadata: &Metadata{
			Timestamp: time.Now().Format(time.RFC3339),
		},
	})
}

func ValidationError(w http.ResponseWriter, details []ErrorDetail) {
	JSON(w, http.StatusUnprocessableEntity, Response{
		Success: false,
		Status:  "error",
		Message: "Validation failed",
		Error: &ErrorInfo{
			Code:    "VALIDATION_ERROR",
			Type:    "ValidationError",
			Details: details,
		},
		Metadata: &Metadata{
			Timestamp: time.Now().Format(time.RFC3339),
		},
	})
}

func Paginated(w http.ResponseWriter, data interface{}, pagination Pagination) {
	JSON(w, http.StatusOK, Response{
		Success:    true,
		Status:     "success",
		Data:       data,
		Pagination: &pagination,
		Metadata: &Metadata{
			Timestamp: time.Now().Format(time.RFC3339),
		},
	})
}
```

## Implementation: Node.js/TypeScript

```typescript
interface ErrorDetail {
  field?: string;
  message: string;
  code?: string;
}

interface ErrorInfo {
  code: string;
  type: string;
  details?: ErrorDetail[];
}

interface Metadata {
  timestamp: string;
  requestId?: string;
  path?: string;
}

interface Pagination {
  page: number;
  pageSize: number;
  totalItems: number;
  totalPages: number;
  hasNext: boolean;
  hasPrevious: boolean;
}

interface ApiResponse<T = any> {
  success: boolean;
  status: string;
  message?: string;
  data?: T;
  error?: ErrorInfo;
  pagination?: Pagination;
  metadata?: Metadata;
}

export class ResponseBuilder {
  static success<T>(res: any, data: T, message?: string): void {
    const response: ApiResponse<T> = {
      success: true,
      status: 'success',
      data,
      metadata: {
        timestamp: new Date().toISOString(),
      },
    };

    if (message) {
      response.message = message;
    }

    res.status(200).json(response);
  }

  static created<T>(res: any, data: T): void {
    const response: ApiResponse<T> = {
      success: true,
      status: 'success',
      message: 'Resource created successfully',
      data,
      metadata: {
        timestamp: new Date().toISOString(),
      },
    };

    res.status(201).json(response);
  }

  static error(
    res: any,
    status: number,
    code: string,
    message: string
  ): void {
    const response: ApiResponse = {
      success: false,
      status: 'error',
      message,
      error: {
        code,
        type: this.getStatusText(status),
      },
      metadata: {
        timestamp: new Date().toISOString(),
      },
    };

    res.status(status).json(response);
  }

  static validationError(
    res: any,
    details: ErrorDetail[]
  ): void {
    const response: ApiResponse = {
      success: false,
      status: 'error',
      message: 'Validation failed',
      error: {
        code: 'VALIDATION_ERROR',
        type: 'ValidationError',
        details,
      },
      metadata: {
        timestamp: new Date().toISOString(),
      },
    };

    res.status(422).json(response);
  }

  static paginated<T>(
    res: any,
    data: T[],
    pagination: Pagination
  ): void {
    const response: ApiResponse<T[]> = {
      success: true,
      status: 'success',
      data,
      pagination,
      metadata: {
        timestamp: new Date().toISOString(),
      },
    };

    res.status(200).json(response);
  }

  private static getStatusText(status: number): string {
    const statusTexts: Record<number, string> = {
      400: 'Bad Request',
      401: 'Unauthorized',
      403: 'Forbidden',
      404: 'Not Found',
      409: 'Conflict',
      422: 'Unprocessable Entity',
      429: 'Too Many Requests',
      500: 'Internal Server Error',
      502: 'Bad Gateway',
      503: 'Service Unavailable',
    };
    return statusTexts[status] || 'Error';
  }
}
```

## Implementation: Python

```python
from dataclasses import dataclass, field
from datetime import datetime
from typing import Any, List, Optional
from flask import jsonify

@dataclass
class ErrorDetail:
    field: Optional[str] = None
    message: str = ""
    code: Optional[str] = None

@dataclass
class ErrorInfo:
    code: str
    type: str
    details: Optional[List[ErrorDetail]] = None

@dataclass
class Metadata:
    timestamp: str
    request_id: Optional[str] = None
    path: Optional[str] = None

@dataclass
class Pagination:
    page: int
    page_size: int
    total_items: int
    total_pages: int
    has_next: bool
    has_previous: bool

@dataclass
class ApiResponse:
    success: bool
    status: str
    message: Optional[str] = None
    data: Optional[Any] = None
    error: Optional[ErrorInfo] = None
    pagination: Optional[Pagination] = None
    metadata: Optional[Metadata] = None

    def to_dict(self):
        return {k: v for k, v in self.__dict__.items() if v is not None}

def success(data: Any, message: Optional[str] = None) -> tuple:
    """Return a success response."""
    response = ApiResponse(
        success=True,
        status="success",
        data=data,
        message=message,
        metadata=Metadata(timestamp=datetime.utcnow().isoformat())
    )
    return jsonify(response.to_dict()), 200

def created(data: Any) -> tuple:
    """Return a created response."""
    response = ApiResponse(
        success=True,
        status="success",
        message="Resource created successfully",
        data=data,
        metadata=Metadata(timestamp=datetime.utcnow().isoformat())
    )
    return jsonify(response.to_dict()), 201

def error(status: int, code: str, message: str) -> tuple:
    """Return an error response."""
    response = ApiResponse(
        success=False,
        status="error",
        message=message,
        error=ErrorInfo(code=code, type=f"HTTP_{status}"),
        metadata=Metadata(timestamp=datetime.utcnow().isoformat())
    )
    return jsonify(response.to_dict()), status

def validation_error(details: List[ErrorDetail]) -> tuple:
    """Return a validation error response."""
    response = ApiResponse(
        success=False,
        status="error",
        message="Validation failed",
        error=ErrorInfo(
            code="VALIDATION_ERROR",
            type="ValidationError",
            details=details
        ),
        metadata=Metadata(timestamp=datetime.utcnow().isoformat())
    )
    return jsonify(response.to_dict()), 422

def paginated(data: List[Any], pagination: Pagination) -> tuple:
    """Return a paginated response."""
    response = ApiResponse(
        success=True,
        status="success",
        data=data,
        pagination=pagination,
        metadata=Metadata(timestamp=datetime.utcnow().isoformat())
    )
    return jsonify(response.to_dict()), 200
```

## Tags

`api`, `rest`, `response-format`, `template`, `universal`, `backend`
