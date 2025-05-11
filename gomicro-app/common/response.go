package common

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// SuccessResponse represents a standard success response
type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewErrorResponse creates a new error response
func NewErrorResponse(status int, message string, err error) ErrorResponse {
	return ErrorResponse{
		Status:  status,
		Message: message,
		Error:   err.Error(),
	}
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(status int, message string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`