package handlers

import (
	"net/http"

	"github.com/manuelmtzv/brevio/internal/http/render"
	"github.com/manuelmtzv/brevio/internal/models"
	"go.uber.org/zap"
)

type HealthHandler struct {
	Logger *zap.SugaredLogger
}

func NewHealthHandler(logger *zap.SugaredLogger) *HealthHandler {
	return &HealthHandler{
		Logger: logger,
	}
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, http.StatusOK, models.HealthCheckResponse{
		Status: "ok",
	})
}
