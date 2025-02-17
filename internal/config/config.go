package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"os"
)

type Config struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`
	TgToken  string `env:"TG_TOKEN" envDefault:""`
}

func Read() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, errors.Wrap(err, "failed to load .env file")
		}
	}
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, errors.Wrap(err, "failed to parse config")
	}
	return cfg, nil
}
