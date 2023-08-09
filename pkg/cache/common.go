package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)
type CacheManager struct {
	deviceCache *cache.Cache
	userCache *cache.Cache
}

func NewCacheManager() *CacheManager {
	return &CacheManager{
		deviceCache: cache.New(2*time.Hour, 1*time.Hour),
		userCache: cache.New(2*time.Hour, 1*time.Hour),
	}
}