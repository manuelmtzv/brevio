package handlers

import (
	"net/http"

	"github.com/manuelmtzv/brevio/internal/api/presenters"
	"github.com/manuelmtzv/brevio/internal/api/services"
	"github.com/manuelmtzv/brevio/internal/http/render"
	"github.com/manuelmtzv/brevio/internal/i18n"
	"go.uber.org/zap"
)

type HealthHandler struct {
	service   services.HealthService
	localizer i18n.Localizer
	logger    *zap.SugaredLogger
}

func NewHealthHandler(
	service services.HealthService,
	localizer i18n.Localizer,
	logger *zap.SugaredLogger,
) *HealthHandler {
	return &HealthHandler{
		service:   service,
		localizer: localizer,
		logger:    logger,
	}
}

func (h *HealthHandler) HandleCheck(w http.ResponseWriter, r *http.Request) error {
	result, err := h.service.Check()
	if err != nil {
		return err
	}

	resp := presenters.Health(r.Context(), result, h.localizer)

	return render.JSON(w, http.StatusOK, resp)
}
