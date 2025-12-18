package api

import (
	"errors"
	"net/http"

	appErrors "github.com/manuelmtzv/brevio/internal/errors"

	"github.com/manuelmtzv/brevio/internal/http/render"
	"github.com/manuelmtzv/brevio/internal/http/response"
)

type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error)

func NewErrorHandler(localizer Localizer) ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		var appErr *appErrors.AppError

		if !errors.As(err, &appErr) {
			appErr = appErrors.Internal(err)
		}

		message := localizer.Message(
			r.Context(),
			appErr.MessageID,
			"internal server error",
			nil,
		)

		render.JSON(w, appErr.HTTPStatus, response.ErrorResponse{
			Error: message,
		})
	}
}
