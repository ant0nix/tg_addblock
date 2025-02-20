package repositories

import "context"

type CacheInterface interface {
	Add(key int64, value any)
	GetKeys() string // debug func
}

type RedisInterface interface {
	Get(ctx context.Context, key string) (int, error)
	Set(ctx context.Context, key string, value int) error
}
