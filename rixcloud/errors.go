package rixcloud

import "fmt"

// APIError the gerenral error
type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("rixcloud api error status: %d message: %v", e.Status, e.Message)
}

// firstError returns the first error among err and apiError which is non-nil
// (or non-Empty in the case of apiError) or nil if neither represent errors.
//
// A common use case is an API which prefers to return any a network error,
// if any, and return an API error in the absence of network errors.
func firstError(err error, apiError *APIError) error {
	if err != nil {
		return err
	}
	// Status == 0 means no error
	if apiError != nil && apiError.Status != 0 {
		return apiError
	}
	return nil
}
