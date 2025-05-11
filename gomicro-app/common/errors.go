package common

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrNotFound      = errors.New("resource not found")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrInvalidInput  = errors.New("invalid input")
	ErrInternalError = errors.New("internal server error")
)

func ErrorResponse(err error) APIResponse {
	return APIResponse{
		Success: false,
		Error:   err.Error(),
	}
}

// HandleError handles errors and returns a JSON error response
func HandleError(w http.ResponseWriter, status int, message string, err error) {
	response := NewErrorResponse(status, message, err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
} 