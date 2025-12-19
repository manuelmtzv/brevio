package services

import (
	"github.com/manuelmtzv/brevio/internal/models"
	"github.com/manuelmtzv/brevio/internal/store"
	"go.uber.org/zap"
)

type healthService struct {
	store  store.Storage
	logger *zap.SugaredLogger
}

func NewHealthService(store store.Storage, logger *zap.SugaredLogger) HealthService {
	return &healthService{
		store:  store,
		logger: logger,
	}
}

func (s *healthService) Check() (*models.HealthCheck, error) {
	return &models.HealthCheck{
		Status: models.HealthOK,
	}, nil
}
