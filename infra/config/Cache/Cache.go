package cache

import (
	"sync"
	"time"
)

type CacheItem struct {
	Data any
	Expiration int64
}

var (
	cache = make(map[string]CacheItem)
	mutex = sync.RWMutex{}
)

func GetCache() map[string]CacheItem {
	return cache
}

func SetCache(key string, value any) {
	mutex.Lock()
	cache[key] = CacheItem{
		Data: value,
		Expiration: time.Now().Add(time.Hour * 24).Unix(),
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

func IsExpired(item CacheItem) bool {
	return item.Expiration < time.Now().Unix()
}	