package repositories

import "context"

type CacheInterface interface {
	Add(key int64, value any)
	Get(key int64) ([]any, error)
	GetKeys() string // debug func
}

type RepositoryInterface interface {
	Insert(ctx context.Context, data any) error
}
