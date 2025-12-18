package main

import (
	"github.com/manuelmtzv/brevio/internal/api"
	"github.com/manuelmtzv/brevio/internal/api/handlers"
	"github.com/manuelmtzv/brevio/internal/config"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer func() {
		_ = logger.Sync()
	}()

	logger.Info("Welcome to Brevio!")

	cfg := config.LoadConfig()

	app := api.NewApplication(cfg, logger)
	handlers := handlers.NewHandlers(logger)
	router := api.NewRouter(api.RouterDeps{
		Logger: logger,
		Handlers: handlers,
	})  

	app.SetRouter(
		router,
	)

	if err := app.Run(); err != nil {
		logger.Error(err.Error())
	}
}
