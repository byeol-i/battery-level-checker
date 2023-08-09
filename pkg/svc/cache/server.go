package cacheSvc

import (
	"context"

	pb_svc_cache "github.com/byeol-i/battery-level-checker/pb/svc/cache"
	cacheManager "github.com/byeol-i/battery-level-checker/pkg/cache"
)
type CacheSrv struct {
	pb_svc_cache.CacheServer
	cacheManager *cacheManager.CacheManager
}

func NewCacheServiceServer(cacheManager *cacheManager.CacheManager) *CacheSrv {
	return &CacheSrv{
		cacheManager: cacheManager,
	}
}

func (s CacheSrv) WriteMsg(ctx context.Context, in *pb_svc_cache.WriteMsgReq) (*pb_svc_cache.WriteMsgRes, error) {
	s.cacheManager.Write(in.DeviceId, in.UserId, in.Value)
	
	return &pb_svc_cache.WriteMsgRes{}, nil
}

func (s CacheSrv) GetCurrentMsg(ctx context.Context, in *pb_svc_cache.GetCurrentMsgReq) (*pb_svc_cache.GetCurrentMsgRes, error) {
	result := s.cacheManager.GetDataFromUserId(in.UserId)

	return &pb_svc_cache.GetCurrentMsgRes{
		Result: result,
	}, nil
}