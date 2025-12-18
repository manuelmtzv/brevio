package errors

import "net/http"

func Internal(err error) *AppError {
	return &AppError{
		Code:       "internal_error",
		MessageID:  "error.internal",
		HTTPStatus: http.StatusInternalServerError,
		Err:        err,
	}
}

func NotFound(err error) *AppError {
	return &AppError{
		Code:       "not_found",
		MessageID:  "error.not_found",
		HTTPStatus: http.StatusNotFound,
		Err:        err,
	}
}

func BadRequest(err error) *AppError {
	return &AppError{
		Code:       "bad_request",
		MessageID:  "error.bad_request",
		HTTPStatus: http.StatusBadRequest,
		Err:        err,
	}
}
