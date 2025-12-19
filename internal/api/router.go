package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/manuelmtzv/brevio/internal/api/handlers"
	"go.uber.org/zap"

	customMiddleware "github.com/manuelmtzv/brevio/internal/api/middleware"
)

type RouterDeps struct {
	Logger       *zap.SugaredLogger
	Handlers     *handlers.Handlers
	ErrorHandler ErrorHandler
}

func NewRouter(deps RouterDeps) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(customMiddleware.Language)

	r.Route("/api", func(r chi.Router) {
		r.Get("/health", Adapt(deps.Handlers.Health.Check, deps.ErrorHandler))
	})

	return r
}
