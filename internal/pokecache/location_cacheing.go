package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	mapper map[string]cacheEntry
	mtx    *sync.Mutex
}

func NewCache(alive time.Duration) *Cache {
	newCache := &Cache{
		mapper: make(map[string]cacheEntry),
		mtx:    &sync.Mutex{},
	}
	go newCache.reapLoop(alive)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.mapper[key] = cacheEntry{
		createdAt: time.Now(),
		value:     val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if res, exist := c.mapper[key]; exist == true {
		return res.value, true
	}
	return nil, false
}

func (c *Cache) reapLoop(alive time.Duration) {
	ticker := time.NewTicker(alive)
	for range ticker.C {
		c.mtx.Lock()
		for key := range c.mapper {
			if time.Since(c.mapper[key].createdAt) > alive {
				delete(c.mapper, key)
			}
		}
		c.mtx.Unlock()
	}
}
