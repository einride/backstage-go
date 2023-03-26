package catalog

import (
	"net/http"
)

// StatusError represents an HTTP status error.
type StatusError struct {
	// Status of the error.
	Status string
	// StatusCode of the error.
	StatusCode int
}

func newStatusError(httpResponse *http.Response) error {
	return &StatusError{
		Status:     httpResponse.Status,
		StatusCode: httpResponse.StatusCode,
	}
}

// Error implements error.
func (s *StatusError) Error() string {
	return s.Status
}
