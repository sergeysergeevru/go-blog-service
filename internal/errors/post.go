package errors

import "fmt"

var ErrPostNotFound = fmt.Errorf("post not found")

type PostValidationError struct {
	reason string
}

func (p PostValidationError) Error() string {
	return p.reason
}

func NewPostValidationError(reason string) *PostValidationError {
	return &PostValidationError{
		reason: reason,
	}
}
