package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/manuelmtzv/brevio/internal/api"
	"github.com/manuelmtzv/brevio/internal/api/handlers"
	"github.com/manuelmtzv/brevio/internal/api/presenters"
	"github.com/manuelmtzv/brevio/internal/api/services"
	"github.com/manuelmtzv/brevio/internal/config"
	appi18n "github.com/manuelmtzv/brevio/internal/i18n"
	"github.com/manuelmtzv/brevio/internal/shorturl/code"
	"github.com/manuelmtzv/brevio/internal/shorturl/ttl"
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
	bundle := appi18n.InitI18n()

	redisAddr, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		logger.Fatal("failed to parse redis url")
	}

	localizer := api.NewLocalizer(bundle)
	urlBuilder := presenters.NewHTTPURLBuilder(cfg.BaseURL)
	gen := code.Base62{Length: cfg.CodeLength}
	ttlPolicy := ttl.NewFixed(cfg.TTL)

	redisClient := redis.NewClient(redisAddr)
	storage := store.NewStorage(redisClient)

	services := services.NewServices(services.ServiceDeps{
		Store:     storage,
		Gen:       gen,
		TTLPolicy: ttlPolicy,
		Logger:    logger,
	})

	handlers := handlers.NewHandlers(handlers.HandlerDeps{
		Health:     services.Health,
		ShortURLs:  services.ShortURLs,
		URLBuilder: urlBuilder,
		Localizer:  localizer,
		Logger:     logger,
	})

	router := api.NewRouter(api.RouterDeps{
		Logger:       logger,
		Handlers:     handlers,
		ErrorHandler: api.NewErrorHandler(localizer, logger),
	})

	app := api.NewApplication(cfg, logger)
	app.SetRouter(
		router,
	)

	if err := app.Run(); err != nil {
		logger.Error(err.Error())
	}
}
