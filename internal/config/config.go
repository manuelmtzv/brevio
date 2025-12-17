package config

import "github.com/manuelmtzv/brevio/internal/env"

type Config struct {
	Port     string
	RedisURL string
}

func LoadConfig() *Config {
	env.Load()

	return &Config{
		Port:     env.GetString("PORT", "8080"),
		RedisURL: env.GetRequired("REDIS_URL"),
	}
}
