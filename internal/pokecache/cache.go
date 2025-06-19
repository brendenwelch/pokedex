package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, exists := c.entry[key]
	c.mu.Unlock()
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c *Cache) reapLoop(interval time.Duration) {
	nextReap := time.Now().Add(interval)
	for {
		if until := time.Until(nextReap); until > 0 {
			time.Sleep(until)
		}

		c.mu.Lock()
		for key, entry := range c.entry {
			if time.Since(entry.createdAt) >= interval {
				delete(c.entry, key)
			}
		}
		c.mu.Unlock()

		nextReap = time.Now().Add(interval)
	}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entry: map[string]cacheEntry{},
		mu:    sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}
