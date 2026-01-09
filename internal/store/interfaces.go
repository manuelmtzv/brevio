package store

import (
	"context"

	"github.com/manuelmtzv/brevio/internal/models"
)

type ShortURLStorage interface {
	Create(ctx context.Context, data models.CreateShortURL) (*models.ShortURL, error)
	FindByCode(ctx context.Context, code string) (*models.ShortURL, error)
}
