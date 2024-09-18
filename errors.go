package emailtemplates

import (
	"errors"
	"fmt"
)

var (
	// ErrMissingTemplate is returned when an email template is missing from the template directory
	ErrMissingTemplate = errors.New("missing email template")
)

// MissingRequiredFieldError is returned when a required field was not provided in a request
type MissingRequiredFieldError struct {
	// RequiredField that is missing
	RequiredField string
}

// Error returns the MissingRequiredFieldError in string format
func (e *MissingRequiredFieldError) Error() string {
	return fmt.Sprintf("%s is required", e.RequiredField)
}

// newMissingRequiredField returns an error for a missing required field
func newMissingRequiredFieldError(field string) *MissingRequiredFieldError {
	return &MissingRequiredFieldError{
		RequiredField: field,
	}
}
