package cacheclient

import (
	"context"
	redisclient "github.com/ant0nix/tg_addblock/internal/bootstrap/redis"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type RedisRepository struct {
	redis *redis.Client
}

func NewRedisRepository(redis *redisclient.Redis) *RedisRepository {
	return &RedisRepository{redis: redis.Redis}
}

func (r *RedisRepository) Get(ctx context.Context, key string) (int, error) {
	resp := r.redis.Get(ctx, key)
	if resp.Err() != nil {
		return 0, resp.Err()
	}
	amount := resp.Val()
	return strconv.Atoi(amount)
}

func (r *RedisRepository) Set(ctx context.Context, key string, value int) error {
	return r.redis.Set(ctx, key, strconv.Itoa(value), 0).Err()
}
