package models

import "time"

type ErrorResponse struct {
	Error     string                 `json:"error"`
	Code      string                 `json:"code,omitempty"`
	Details   string                 `json:"details,omitempty"`
	Timestamp time.Time              `json:"timestamp"`
	Path      string                 `json:"path"`
	Method    string                 `json:"method"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

type SuccessResponse struct {
	Message   string                 `json:"message,omitempty"`
	Data      interface{}            `json:"data,omitempty"`
	Total     int                    `json:"total,omitempty"`
	Timestamp time.Time              `json:"timestamp"`
	Meta      map[string]interface{} `json:"meta,omitempty"`
}

const (
	ErrCodeValidation        = "VALIDATION_ERROR"
	ErrCodeNotFound          = "NOT_FOUND"
	ErrCodeConflict          = "CONFLICT"
	ErrCodeInternalServer    = "INTERNAL_SERVER_ERROR"
	ErrCodeBadRequest        = "BAD_REQUEST"
	ErrCodeUnauthorized      = "UNAUTHORIZED"
	ErrCodeForbidden         = "FORBIDDEN"
	ErrCodeInvalidFormat     = "INVALID_FORMAT"
	ErrCodeDuplicateResource = "DUPLICATE_RESOURCE"
)
