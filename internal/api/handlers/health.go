package handlers

import (
	"net/http"

	"github.com/manuelmtzv/brevio/internal/api/services"
	"github.com/manuelmtzv/brevio/internal/http/render"
	"go.uber.org/zap"
)

type HealthHandler struct {
	Service services.HealthService
	Logger  *zap.SugaredLogger
}

func NewHealthHandler(service services.HealthService, logger *zap.SugaredLogger) *HealthHandler {
	return &HealthHandler{
		Service: service,
		Logger:  logger,
	}
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) error {
	result, err := h.Service.Check()
	if err != nil {
		return err
	}

	render.JSON(w, http.StatusOK, result)
	return nil
}
