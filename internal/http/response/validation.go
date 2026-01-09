package response

import (
	"context"

	"github.com/go-playground/validator/v10"
	appi18n "github.com/manuelmtzv/brevio/internal/i18n"
)

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type ValidationErrorsResponse struct {
	Errors []ValidationError `json:"errors"`
}

func ValidationErrors(
	errs validator.ValidationErrors,
	localizer appi18n.Localizer,
	ctx context.Context,
) ValidationErrorsResponse {
	out := make([]ValidationError, 0, len(errs))

	for _, fe := range errs {
		out = append(out, ValidationError{
			Field: fe.Field(),
			Error: localizer.Message(
				ctx,
				"Validation."+fe.Tag(),
				validationFallback(fe),
				map[string]any{
					"Field": fe.Field(),
					"Value": fe.Value(),
				},
			),
		})
	}

	return ValidationErrorsResponse{Errors: out}
}

func validationFallback(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is required"
	case "url":
		return "must be a valid URL"
	default:
		return "is invalid"
	}
}
