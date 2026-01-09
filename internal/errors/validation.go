package errors

import "net/http"

func Validation(err error) *AppError {
	return &AppError{
		Code:       "validation_error",
		MessageID:  "Error.Validation",
		HTTPStatus: http.StatusBadRequest,
		Err:        err,
	}
}
