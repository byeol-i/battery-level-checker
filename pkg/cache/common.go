package cache

import (
	"github.com/patrickmn/go-cache"
)

type CacheManager struct {
	deviceCache *cache.Cache
	userCache *cache.Cache
}

func NewCacheManager(deviceCache, userCache *cache.Cache) *CacheManager {
	return &CacheManager{
		deviceCache: deviceCache,
		userCache: userCache,
	}
}