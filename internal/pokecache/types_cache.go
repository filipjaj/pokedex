package pokecache

import "time"

type Cache struct {
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}


func (c *Cache) Add(key string, val []byte) {
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func NewCache(reapInterval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
	}
	go c.reapLoop(reapInterval)
	return c
}

func (c *Cache) reapLoop(reapInterval time.Duration) {
	ticker := time.NewTicker(reapInterval)
	for {
		<-ticker.C
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > reapInterval {
				delete(c.entries, key)
			}
		}
	}
}
