package config

import (
	"time"

	"github.com/manuelmtzv/brevio/internal/env"
)

type Config struct {
	Port       string
	BaseURL    string
	RedisURL   string
	CodeLength int
	TTL        time.Duration
}

func LoadConfig() *Config {
	env.Load()

	return &Config{
		Port:       env.GetString("PORT", "8080"),
		BaseURL:    env.GetString("BASE_URL", "http://localhost:8080"),
		RedisURL:   env.GetRequired("REDIS_URL"),
		CodeLength: env.GetInt("CODE_LENGTH", 7),
		TTL:        time.Duration(env.GetInt("TTL", 86400)) * time.Second,
	}
}
