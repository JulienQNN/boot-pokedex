package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entries: make(map[string]cacheEntry),
		mux:     &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.entries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.entries, k)
		}
	}
}
