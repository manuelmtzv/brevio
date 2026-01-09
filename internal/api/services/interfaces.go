package services

import (
	"context"

	"github.com/manuelmtzv/brevio/internal/models"
)

type HealthService interface {
	Check() (*models.HealthCheck, error)
}

type ShortURLService interface {
	Create(ctx context.Context, input CreateShortURLInput) (*models.ShortURL, error)
}
