package handlers

import (
	"github.com/manuelmtzv/brevio/internal/api/services"
	"go.uber.org/zap"
)

type ShortURLHandler struct {
	Service services.ShortURLService
	Logger  *zap.SugaredLogger
}

func NewShortURLHandler(service services.ShortURLService, logger *zap.SugaredLogger) *ShortURLHandler {
	return &ShortURLHandler{
		Service: service,
		Logger:  logger,
	}
}
