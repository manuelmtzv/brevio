package services

import (
	"github.com/manuelmtzv/brevio/internal/store"
	"go.uber.org/zap"
)

type Services struct {
	Health    HealthService
	ShortURLs ShortURLService
}

func NewServices(store store.Storage, logger *zap.SugaredLogger) *Services {
	return &Services{
		Health: NewHealthService(store, logger),
	}
}
