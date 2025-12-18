package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/manuelmtzv/brevio/internal/api/handlers"
	"go.uber.org/zap"
)

type RouterDeps struct {
	Logger   *zap.SugaredLogger
	Handlers *handlers.Handlers
}

func NewRouter(deps RouterDeps) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Get("/health", deps.Handlers.Health.Check)
	})

	return r
}
