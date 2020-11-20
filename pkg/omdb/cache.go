package omdb

import "sync"

type Cache interface {
	Put(key string, movie Movie)
	Get(key string) (Movie, bool)
}

type MemoryCache struct {
	sync.Map
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{}
}

func (m *MemoryCache) Put(key string, movie Movie) {
	m.Store(key, movie)
}

func (m *MemoryCache) Get(key string) (Movie, bool) {
	res, ok := m.Load(key)
	if ok {
		return res.(Movie), ok
	}
	return Movie{}, ok
}

func (m *MemoryCache) Del(key string) {
	m.Delete(key)
}
