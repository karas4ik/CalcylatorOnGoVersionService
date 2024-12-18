package cache

import (
	"time"
)

type CacheItem struct {
	Value      string
	Expiration time.Time
}

type Cache struct {
	items map[string]CacheItem
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]CacheItem),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	item, found := c.items[key]
	if !found || time.Now().After(item.Expiration) {
		return "", false
	}
	return item.Value, true
}

func (c *Cache) Set(key string, value string, duration time.Duration) {
	c.items[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(duration),
	}
}
