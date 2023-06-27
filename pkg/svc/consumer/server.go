package consumerSvc

import (
	"context"

	pb_svc_consumer "github.com/byeol-i/battery-level-checker/pb/svc/consumer"
	"github.com/byeol-i/battery-level-checker/pb/unit/common"
	"github.com/byeol-i/battery-level-checker/pkg/consumer"
)

type ConsumerSrv struct {
	pb_svc_consumer.ConsumerServer
}

func NewConsumerServiceServer() *ConsumerSrv {
	return &ConsumerSrv{}
}

// func CreateNewMsg

func CreateNewTopic(ctx context.Context, in *pb_svc_consumer.CreateNewTopicReq) (*pb_svc_consumer.CreateNewTopicRes, error) {

	admin, err := consumer.GetAdmin()
	if err != nil {
		return &pb_svc_consumer.CreateNewTopicRes{
			Result: &common.ReturnMsg{
				Error: err.Error(),
			},
		}, err 
	}

	err = consumer.CreateTopic(admin, in.Topic)
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