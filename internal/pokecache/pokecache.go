package pokecache
import (
	"time"
	"sync"
)


type Cache struct {
	entries map[string]cacheEntry
	interval time.Duration
	mu sync.Mutex
}


type cacheEntry struct {
	createdAt time.Time
	val []byte 
}



func NewCache(interval time.Duration) *Cache {
	c := &Cache {
		entries: make(map[string]cacheEntry),
		interval: interval,
	}

	go c.reapLoop()
	return c
}


func (s *Cache) Add(key string, val []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (s *Cache) Get(key string) ([]byte, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	val, ok := s.entries[key]
	if ok {
		return val.val, true
	} else {
		return nil, false
	}
}

func (s *Cache) reapLoop() {
	ticker := time.NewTicker(s.interval)
	for range ticker.C {
		s.mu.Lock()

		for key, entry := range s.entries {
			age := time.Since(entry.createdAt)
			if age > s.interval {
				delete(s.entries, key)
			}
		}

		s.mu.Unlock()
	}
}