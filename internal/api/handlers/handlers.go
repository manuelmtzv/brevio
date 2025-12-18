package handlers

import "go.uber.org/zap"

type Handlers struct {
	Health   *HealthHandler
	ShortURL *ShortURLHandler
}

func NewHandlers(logger *zap.SugaredLogger) *Handlers {
	return &Handlers{
		Health:   NewHealthHandler(logger),
		ShortURL: NewShortURLHandler(logger),
	}
}
