package cacheSvc

import (
	pb_svc_cache "github.com/byeol-i/battery-level-checker/pb/svc/cache"
)
type CacheSrv struct {
	pb_svc_cache.CacheServer
}

func NewCacheServiceServer() *CacheSrv {
	return &CacheSrv{}
}

