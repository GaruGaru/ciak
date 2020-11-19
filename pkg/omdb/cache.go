package omdb

import "sync"

type Cache interface {
	Put(key string, movie Movie)
	Get(key string) (Movie, bool)
}

type MemoryCache struct {
	cacheMap map[string]Movie
	mutex *sync.RWMutex
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		cacheMap: make(map[string]Movie),
		mutex: &sync.RWMutex{},
	}
}

func (m *MemoryCache) Put(key string, movie Movie) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.cacheMap[key] = movie
}

func (m *MemoryCache) Get(key string) (Movie, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	movie, present := m.cacheMap[key]
	return movie, present
}
