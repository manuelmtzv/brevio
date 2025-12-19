package handlers

import (
	"github.com/manuelmtzv/brevio/internal/api/services"
	"go.uber.org/zap"

	appi18n "github.com/manuelmtzv/brevio/internal/i18n"
)

type HandlerDeps struct {
	Health    services.HealthService
	ShortURLs services.ShortURLService
	Localizer appi18n.Localizer
	Logger    *zap.SugaredLogger
}

type Handlers struct {
	Health   *HealthHandler
	ShortURL *ShortURLHandler
}

func NewHandlers(deps HandlerDeps) *Handlers {
	return &Handlers{
		Health:   NewHealthHandler(deps.Health, deps.Localizer, deps.Logger),
		ShortURL: NewShortURLHandler(deps.ShortURLs, deps.Localizer, deps.Logger),
	}
}
