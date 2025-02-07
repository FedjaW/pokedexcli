package pokecache

import (
	"fmt"
	"time"
    "sync"
)

type cacheEntry struct {
    createdAt time.Time
    val []byte
}

type Cache struct {
    mu sync.Mutex
    Cache map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
    fmt.Println("Chach me if you can!")
    cache := Cache{
        Cache: make(map[string]cacheEntry),
    }
    cache.reapLoop(interval)
    return &cache
}

func (cache *Cache) Add(key string, val []byte) {
    cacheEntry := cacheEntry{
        createdAt: time.Now(),
        val: val,
    }

    cache.Cache[key] = cacheEntry
}

func (cache *Cache) Get(key string) ([]byte, bool) {
    cacheEntry, exists := cache.Cache[key]
    fmt.Println(exists)
    if !exists {
        return []byte{}, false
    }

    return cacheEntry.val, true
}


func (cache *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    // defer ticker.Stop()
    go func() {
        for {
            select {
            case t := <- ticker.C:
                // Syncronize access to the cache
                cache.mu.Lock()
                for key, val := range cache.Cache {
                    if t.Sub(val.createdAt) > interval  {
                        // remove old cache entries
                        delete(cache.Cache, key)
                    }
                }
                cache.mu.Unlock()
            }
        }
    }()
}
