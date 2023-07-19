package consumer

import (
	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/config"
)

type Consumer struct {
	kafkaConf *sarama.Config
	brokerList []string
}

// type Connector interface {
// 	NewConsumer() *Consumer
// }

// type ConsumerImpl interface {
// 	GetAdmin() (sarama.ClusterAdmin, error)
// 	CreateTopic(admin sarama.ClusterAdmin, name string) (error)
// }	

func NewConsumer() *Consumer {
	manager := config.GetInstance()

	return &Consumer{
		kafkaConf: manager.GetKafkaSarama(),
		brokerList: manager.GetBrokerList(),
	}
}