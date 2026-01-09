package presenters

import (
	"context"
	"strings"

	appi18n "github.com/manuelmtzv/brevio/internal/i18n"
	"github.com/manuelmtzv/brevio/internal/models"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func Health(ctx context.Context, h *models.HealthCheck, l appi18n.Localizer) HealthResponse {
	return HealthResponse{
		Status: l.Message(
			ctx,
			"Health."+strings.ToUpper(string(h.Status)),
			string(h.Status),
			nil,
		),
	}
}
