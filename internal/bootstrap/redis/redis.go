package redisclient

import (
	"context"
	"github.com/ant0nix/tg_addblock/internal/config"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	Redis *redis.Client
}

func New(cfg *config.Config) *Redis {
	return &Redis{
		Redis: redis.NewClient(&redis.Options{
			Addr: cfg.RedisHost,
		}),
	}
}

const redisPint = time.Minute

func (r *Redis) Alive(ctx context.Context) error {
	ticker := time.NewTicker(redisPint)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			err := r.Redis.Ping(ctx).Err()
			if err != nil {
				return err
			}
		}
	}
}
