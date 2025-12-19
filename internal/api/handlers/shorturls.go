package handlers

import (
	"github.com/manuelmtzv/brevio/internal/api/services"
	"github.com/manuelmtzv/brevio/internal/i18n"
	"go.uber.org/zap"
)

type ShortURLHandler struct {
	service   services.ShortURLService
	localizer i18n.Localizer
	logger    *zap.SugaredLogger
}

func NewShortURLHandler(service services.ShortURLService, localizer i18n.Localizer, logger *zap.SugaredLogger) *ShortURLHandler {
	return &ShortURLHandler{
		service:   service,
		localizer: localizer,
		logger:    logger,
	}
}
