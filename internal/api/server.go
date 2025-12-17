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

func (app *Application) ServeHTTP() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config.Port),
		Handler: app.Router,
	}

	shutdownErrorStream := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		app.Logger.Infof("received signal: %s, shutting down server...", s)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			shutdownErrorStream <- err
			return
		}

		shutdownErrorStream <- nil
	}()

	app.Logger.Infow("starting server", "port", app.Config.Port)

	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownErrorStream
	if err != nil {
		return err
	}

	app.Logger.Infow("stopped server", "port", app.Config.Port)

	return nil
}
