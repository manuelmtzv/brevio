package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/manuelmtzv/brevio/internal/api"
	"github.com/manuelmtzv/brevio/internal/api/handlers"
	"github.com/manuelmtzv/brevio/internal/api/services"
	"github.com/manuelmtzv/brevio/internal/config"
	"github.com/manuelmtzv/brevio/internal/store"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer func() {
		_ = logger.Sync()
	}()

	logger.Info("Welcome to Brevio!")

	cfg := config.LoadConfig()

	redisAddr, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		logger.Fatal("failed to parse redis url")
	}

	redisClient := redis.NewClient(redisAddr)
	storage := store.NewStorage(redisClient)
	services := services.NewServices(storage, logger)

	handlers := handlers.NewHandlers(handlers.HandlerDeps{
		Health:    services.Health,
		ShortURLs: services.ShortURLs,
		Logger:    logger,
	})

	router := api.NewRouter(api.RouterDeps{
		Logger:   logger,
		Handlers: handlers,
	})

	app := api.NewApplication(cfg, logger)
	app.SetRouter(
		router,
	)

	if err := app.Run(); err != nil {
		logger.Error(err.Error())
	}
}
