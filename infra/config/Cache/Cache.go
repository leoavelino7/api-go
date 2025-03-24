package cache

import (
	"sync"
	"time"
)

type CacheItem struct {
	Data any
	Expiration int64
	Expired bool
}

var (
	cache = make(map[string]CacheItem)
	mutex = sync.RWMutex{}
)

func init() {
	go cleanupLoop()
}

func cleanupLoop() {
	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		cleanup()
	}
}

func cleanup() {
	now := time.Now().Unix()
	mutex.Lock()

	for _, item := range cache {
		if item.Expiration <= now {
			setIsExpired(&item)
		}
	}

	mutex.Unlock()
}

func GetCache() map[string]CacheItem {
	return cache
}

func SetCache(key string, value any) {
	mutex.Lock()
	cache[key] = CacheItem{
		Data: value,
		Expiration: time.Now().Add(time.Minute * 1).Unix(),
		Expired: false,
	}
	mutex.Unlock()
}

func UpdateCache(key string, value any) {
	mutex.Lock()
	cache[key] = CacheItem{
		Data: value,
		Expiration: time.Now().Add(time.Minute * 1).Unix(),
		Expired: false,
	}
	mutex.Unlock()
}	

func DeleteCache(key string) {
	mutex.Lock()
	delete(cache, key)
	mutex.Unlock()
}

func ClearCache() {	
	mutex.Lock()
	cache = make(map[string]CacheItem)
	mutex.Unlock()
}

func setIsExpired(item *CacheItem) {
	item.Expired = true
}