package services

import (
	"github.com/manuelmtzv/brevio/internal/shorturl/code"
	"github.com/manuelmtzv/brevio/internal/shorturl/ttl"
	"github.com/manuelmtzv/brevio/internal/store"
	"go.uber.org/zap"
)

type ServiceDeps struct {
	Store     store.Storage
	Gen       code.Generator
	TTLPolicy ttl.Policy
	Logger    *zap.SugaredLogger
}

type Services struct {
	Health    HealthService
	ShortURLs ShortURLService
}

func NewServices(
	deps ServiceDeps,
) *Services {
	return &Services{
		Health: NewHealthService(deps.Store, deps.Logger),
		ShortURLs: NewShortURLService(ShortURLServiceDeps{
			Store:     deps.Store,
			Gen:       deps.Gen,
			TTLPolicy: deps.TTLPolicy,
			Logger:    deps.Logger,
		}),
	}
}
