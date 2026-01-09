package api

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	appErrors "github.com/manuelmtzv/brevio/internal/errors"
	appi18n "github.com/manuelmtzv/brevio/internal/i18n"
	"go.uber.org/zap"

	"github.com/manuelmtzv/brevio/internal/http/render"
	"github.com/manuelmtzv/brevio/internal/http/response"
)

type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error)

func NewErrorHandler(localizer appi18n.Localizer, logger *zap.SugaredLogger) ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		logger.Error("error handling request", zap.Error(err))

		var appErr *appErrors.AppError

		if !errors.As(err, &appErr) {
			appErr = appErrors.Internal(err)
		}

		if ve, ok := appErr.Err.(validator.ValidationErrors); ok {
			render.JSON(w, appErr.HTTPStatus,
				response.ValidationErrors(ve, localizer, r.Context()),
			)
			return
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
