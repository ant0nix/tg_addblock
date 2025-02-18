package cacheclient

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

func Init() *Cache {
	return &Cache{Items: make(map[int64]CacheItem)}
}

func (c *Cache) Add(key int64, value any) {
	if _, ok := c.Items[key]; !ok {
		values := make([]any, 5)
		values[0] = value
		c.Items[key] = CacheItem{
			value: values,
			removeFunc: func(ticker *time.Ticker) {
				select {
				case <-ticker.C:
					delete(c.Items, key)
				}
			},
			ticker: time.NewTicker(1 * time.Minute),
		}
		c.Items[key].removeFunc(c.Items[key].ticker)
		return
	}
	ci := c.Items[key]
	ci.ticker.Reset(1 * time.Minute)
	ci.value = append(ci.value, value)
	ci.index++
	c.Items[key] = ci
}

func (c *Cache) Get(key int64) ([]any, error) {
	if ci, ok := c.Items[key]; ok {
		ci.ticker.Reset(1 * time.Minute)
		return ci.value, nil
	}
	return nil, errors.New("item not found")
}

func (c *Cache) GetKeys() string {
	var keys string
	for k := range c.Items {
		keys += fmt.Sprintf("%d,", k)
	}
	return keys
}
