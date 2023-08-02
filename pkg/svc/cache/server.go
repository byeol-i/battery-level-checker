package cacheSvc

import (
	"context"

	pb_svc_cache "github.com/byeol-i/battery-level-checker/pb/svc/cache"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
)
type CacheSrv struct {
	pb_svc_cache.CacheServer
	deviceCache *cache.Cache
}

func NewCacheServiceServer(deviceCache *cache.Cache) *CacheSrv {
	return &CacheSrv{
		deviceCache: deviceCache,
	}
}

func (s CacheSrv) WriteMsg(ctx context.Context, in *pb_svc_cache.WriteMsgReq) (*pb_svc_cache.WriteMsgRes, error) {
	s.deviceCache.Set(in.DeviceId, in.Value, cache.DefaultExpiration)

	return &pb_svc_cache.WriteMsgRes{}, nil
}

func (s CacheSrv) GetMsg(ctx context.Context, in *pb_svc_cache.GetMsgReq) (*pb_svc_cache.GetMsgRes, error) {
	data, found := s.deviceCache.Get(in.DeviceId)
	logger.Info("data:", zap.Any("data", data))

	if found {
		if bytes, ok := data.([]byte); ok {
			return &pb_svc_cache.GetMsgRes{
				Result: string(bytes[:]),
			}, nil
		} else {
			return &pb_svc_cache.GetMsgRes{
				Result: "something is wrong",
			}, nil
		}
	}

	return &pb_svc_cache.GetMsgRes{
		Result: "nothing",
	}, nil
}