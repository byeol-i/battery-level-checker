package cacheSvc

import (
	"context"
	"time"

	pb_svc_cache "github.com/byeol-i/battery-level-checker/pb/svc/cache"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
const contextTime = time.Second * 5

type CacheSvcClient struct {
	addr string
}

func NewCacheSvcClient(addr string) (*CacheSvcClient) {
	return &CacheSvcClient{addr: addr}
}

func (c *CacheSvcClient) CallWriteMsg(userId, deviceId string, value []byte) error {	
	ctx, cancel := context.WithTimeout(context.Background(), contextTime)
	defer cancel()
	
	conn, err := grpc.DialContext(ctx, c.addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb_svc_cache.NewCacheClient(conn)

	in := &pb_svc_cache.WriteMsgReq{
		DeviceId: deviceId,
		Value: value,
	}

	_, err = client.WriteMsg(ctx, in)
	if err != nil {
		logger.Error("Can't call grpc call", zap.Error(err) )
		return err
	}

	return nil
}
