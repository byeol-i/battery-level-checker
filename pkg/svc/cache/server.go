package cacheSvc

import (
	"context"

	pb_svc_cache "github.com/byeol-i/battery-level-checker/pb/svc/cache"
	"github.com/byeol-i/battery-level-checker/pkg/cache"
)
type CacheSrv struct {
	pb_svc_cache.CacheServer
	cManager *cache.CacheManager
}

func NewCacheServiceServer(cManager *cache.CacheManager) *CacheSrv {
	return &CacheSrv{
		cManager: cManager,
	}
}

func (s CacheSrv) WriteMsg(ctx context.Context, in *pb_svc_cache.WriteMsgReq) (*pb_svc_cache.WriteMsgRes, error) {
	s.cManager.Write(in.DeviceId, in.UserId, in.Value)
	
	return &pb_svc_cache.WriteMsgRes{}, nil
}

func (s CacheSrv) GetCurrentMsg(ctx context.Context, in *pb_svc_cache.GetCurrentMsgReq) (*pb_svc_cache.GetCurrentMsgRes, error) {
	result := s.cManager.GetDataFromUserId(in.UserId)

	return &pb_svc_cache.GetCurrentMsgRes{
		Result: result,
	}, nil
}