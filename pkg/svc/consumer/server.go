package consumerSvc

import (
	"context"

	pb_svc_consumer "github.com/byeol-i/battery-level-checker/pb/svc/consumer"
	"github.com/byeol-i/battery-level-checker/pb/unit/common"
	"github.com/byeol-i/battery-level-checker/pkg/topic"
)

type ConsumerSrv struct {
	pb_svc_consumer.ConsumerServer
	TopicManager *topic.TopicManager
}

func NewConsumerServiceServer() *ConsumerSrv {
	return &ConsumerSrv{
		TopicManager : topic.NewTopicManager(),
	}
}

// func CreateNewMsg

func (s ConsumerSrv) CreateNewTopic(ctx context.Context, in *pb_svc_consumer.CreateNewTopicReq) (*pb_svc_consumer.CreateNewTopicRes, error) {
	admin, err := s.TopicManager.GetAdmin()
	if err != nil {
		return &pb_svc_consumer.CreateNewTopicRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err 
	}

	err = s.TopicManager.CreateTopic(admin, in.Topic)
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

// func (s ConsumerSrv) GetUserDevices(ctx context.Context, in *pb_svc_consumer.GetUserDevicesReq) (*pb_svc_consumer.GetUserDevicesRes, error) {
// 	result, err := s.Admin.GetUserDevice(in.Uid)
// 	if err != nil {
// 		return &pb_svc_consumer.GetUserDevicesRes{
// 			Result: &common.ReturnMsg{
// 				Error: err.Error(),
// 			},
// 		}, err 
// 	}

// 	return &pb_svc_consumer.GetUserDevicesRes{
// 		Result: &common.ReturnMsg{
// 			Result: strings.Join(result, ","),
// 		},
// 	}, nil
// }