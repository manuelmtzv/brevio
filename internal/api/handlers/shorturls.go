package handlers

import "go.uber.org/zap"

type ShortURLHandler struct {
	Logger *zap.SugaredLogger
}

func NewShortURLHandler(logger *zap.SugaredLogger) *ShortURLHandler {
	return &ShortURLHandler{
		Logger: logger,
	}
}
