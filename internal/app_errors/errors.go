package app_errors

import "fmt"

type ApiError struct {
	Message string
	err     error
}

func NewApiError(message string, err error) *ApiError {
	return &ApiError{Message: message, err: err}
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("message - %s, err - %s", e.Message, e.err.Error())
}

func (e *ApiError) Unwrap() error {
	return e.err
}
