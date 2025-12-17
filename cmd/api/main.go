package main

import (
	"net/http"

	"github.com/manuelmtzv/brevio/internal/api"
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

	app.SetRouter(
		http.DefaultServeMux, // TODO: implement real router
	)

	if err := app.ServeHTTP(); err != nil {
		logger.Error(err.Error())
	}
}
