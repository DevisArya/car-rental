package helper

import (
	"fmt"
)

// Custom error struct untuk validasi
type ValidationErrors struct {
	StatusCode int      `json:"statusCode"`
	Messages   []string `json:"errors"`
}

func (e *ValidationErrors) Error() string {
	return fmt.Sprintf("validation error: %v", e.Messages)
}

// Fungsi untuk membuat error validasi
func NewValidationError(statusCode int, messages []string) *ValidationErrors {
	return &ValidationErrors{
		StatusCode: statusCode,
		Messages:   messages,
	}
}
