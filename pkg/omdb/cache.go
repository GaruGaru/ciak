package omdb

type Cache interface {
	Put(key string, movie Movie)

	Get(key string) (Movie, bool)
}

type MemoryCache struct {
	cacheMap map[string]Movie
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{cacheMap: make(map[string]Movie)}
}

func (m *MemoryCache) Put(key string, movie Movie) {
	m.cacheMap[key] = movie
}

func (m *MemoryCache) Get(key string) (Movie, bool) {
	movie, present := m.cacheMap[key]
	return movie, present
}
