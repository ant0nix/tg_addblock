package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"os"
	"time"
)

type Config struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`
	TgToken  string `env:"TG_TOKEN" envDefault:""`
	ChatID   int64  `env:"CHAT_ID" envDefault:""`
	//MongoDB
	MongoDBURI   string        `env:"MONGODB_URI" envDefault:""`
	Timeout      time.Duration `env:"TIMEOUT" envDefault:"60s"`
	PingInterval time.Duration `env:"PING_INTERVAL" envDefault:"5s"`
	//Redis
	RedisHost string `env:"REDIS_HOST" envDefault:""`
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
