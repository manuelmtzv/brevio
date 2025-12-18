package handlers

import (
	"github.com/manuelmtzv/brevio/internal/api/services"
	"go.uber.org/zap"
)

type HandlerDeps struct {
	Health    services.HealthService
	ShortURLs services.ShortURLService
	Logger    *zap.SugaredLogger
}

type Handlers struct {
	Health   *HealthHandler
	ShortURL *ShortURLHandler
}

func NewHandlers(deps HandlerDeps) *Handlers {
	return &Handlers{
		Health:   NewHealthHandler(deps.Health, deps.Logger),
		ShortURL: NewShortURLHandler(deps.ShortURLs, deps.Logger),
	}
}
