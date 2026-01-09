package services

import (
	"context"

	"github.com/manuelmtzv/brevio/internal/models"
	"github.com/manuelmtzv/brevio/internal/shorturl/code"
	"github.com/manuelmtzv/brevio/internal/shorturl/ttl"
	"github.com/manuelmtzv/brevio/internal/store"
	"go.uber.org/zap"
)

type CreateShortURLInput struct {
	Target string `json:"target"`
}

type ShortURLServiceDeps struct {
	Store     store.Storage
	Gen       code.Generator
	TTLPolicy ttl.Policy
	Logger    *zap.SugaredLogger
}

type shortURLService struct {
	store     store.Storage
	gen       code.Generator
	ttlPolicy ttl.Policy
	logger    *zap.SugaredLogger
}

func NewShortURLService(
	deps ShortURLServiceDeps,
) ShortURLService {
	return &shortURLService{
		store:     deps.Store,
		gen:       deps.Gen,
		ttlPolicy: deps.TTLPolicy,
		logger:    deps.Logger,
	}
}

func (s *shortURLService) Create(ctx context.Context, input CreateShortURLInput) (*models.ShortURL, error) {
	ttl := s.ttlPolicy.TTL()
	return s.store.ShortURLs.Create(ctx, models.CreateShortURL{
		Code:   s.gen.Generate(),
		Target: input.Target,
		TTL:    &ttl,
	})
}
