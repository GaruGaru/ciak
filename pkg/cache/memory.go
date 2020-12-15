package cache

import "sync"

type InMemoryCache struct {
	cache *sync.Map
}

func Memory() *InMemoryCache {
	return &InMemoryCache{cache: &sync.Map{}}
}

func (i InMemoryCache) Get(key interface{}) (interface{}, bool, error) {
	value, found := i.cache.Load(key)
	return value, found, nil
}

func (i InMemoryCache) Set(key interface{}, value interface{}) error {
	i.cache.Store(key, value)
	return nil
}
