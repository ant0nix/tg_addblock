package cacheclient

import "time"

type Cache struct {
	Items map[int64]CacheItem
}
type CacheItem struct {
	value      []any
	index      int
	removeFunc func(ticker *time.Ticker)
	ticker     *time.Ticker
}
