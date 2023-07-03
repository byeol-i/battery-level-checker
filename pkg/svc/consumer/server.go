package consumerSvc

import (
	"context"
	"strings"

	pb_svc_consumer "github.com/byeol-i/battery-level-checker/pb/svc/consumer"
	"github.com/byeol-i/battery-level-checker/pb/unit/common"
	"github.com/byeol-i/battery-level-checker/pkg/consumer"
)

type ConsumerSrv struct {
	pb_svc_consumer.ConsumerServer
	Consumer *consumer.Consumer
}

func NewConsumerServiceServer() *ConsumerSrv {
	return &ConsumerSrv{
		Consumer : consumer.NewConsumer(),
	}
}

// func CreateNewMsg

func (s ConsumerSrv) CreateNewTopic(ctx context.Context, in *pb_svc_consumer.CreateNewTopicReq) (*pb_svc_consumer.CreateNewTopicRes, error) {
	admin, err := s.Consumer.GetAdmin()
	if err != nil {
		return &pb_svc_consumer.CreateNewTopicRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err 
	}

	err = s.Consumer.CreateTopic(admin, in.Topic)
	if err != nil {
		return &pb_svc_consumer.CreateNewTopicRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err 
	}

	return &pb_svc_consumer.CreateNewTopicRes{
		Result: &common.ReturnMsg{
			Result: string("Done"),
		},
	}, nil
}

func (s ConsumerSrv) GetUserDevices(ctx context.Context, in *pb_svc_consumer.GetUserDevicesReq) (*pb_svc_consumer.GetUserDevicesRes, error) {
	result, err := s.Consumer.GetUserDevice(in.Uid)
	if err != nil {
		return &pb_svc_consumer.GetUserDevicesRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err 
	}

	return &pb_svc_consumer.GetUserDevicesRes{
		Result: &common.ReturnMsg{
			Result: strings.Join(result, ","),
		},
	}, nil
}