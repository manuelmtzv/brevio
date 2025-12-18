package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/manuelmtzv/brevio/internal/config"
	"go.uber.org/zap"
)

type Application struct {
	Config *config.Config
	Router http.Handler
	Logger *zap.SugaredLogger
}

func NewApplication(cfg *config.Config, logger *zap.SugaredLogger) *Application {
	return &Application{
		Config: cfg,
		Logger: logger,
	}
}

func (app *Application) SetRouter(router http.Handler) {
	app.Router = router
}

func (app *Application) Run() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config.Port),
		Handler: app.Router,
	}

	shutdownErr := make(chan error, 1)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(quit)

		s := <-quit

		app.Logger.Infof("received signal: %s, shutting down server...", s)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		shutdownErr <- server.Shutdown(ctx)
	}()

	app.Logger.Infow("starting server", "port", app.Config.Port)

	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	if err := <-shutdownErr; err != nil {
		return err
	}

	app.Logger.Infow("stopped server", "port", app.Config.Port)

	return nil
}
