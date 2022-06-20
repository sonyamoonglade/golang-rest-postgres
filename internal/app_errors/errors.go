package app_errors

import "fmt"

type EmptyResultError struct {
	Message string
	err     error
}

func NewEmptyResultError(message string, err error) *EmptyResultError {
	return &EmptyResultError{Message: message, err: err}
}

func (e *EmptyResultError) Error() string {
	return fmt.Sprintf("message - %s, err - %s", e.Message, e.err.Error())
}

func (e *EmptyResultError) Unwrap() error {
	return e.err
}
