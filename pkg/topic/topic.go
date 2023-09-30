package topic

import (
	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/config"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"go.uber.org/zap"
)

type TopicManager struct {
	kafkaConf *sarama.Config
	brokerList []string
	numOfReplicationFactor int16
	numOfPartitions int32
}

func NewTopicManager() *TopicManager {
	manager := config.GetInstance()

	return &TopicManager{
		kafkaConf: manager.GetKafkaSarama(),
		brokerList: manager.GetBrokerList(),
		numOfReplicationFactor: manager.GetNumOfReplicationFactor(),
		numOfPartitions: manager.GetNumOfPartitions(),
	}
}

func (t *TopicManager) GetAdmin() (sarama.ClusterAdmin, error){
	admin, err := sarama.NewClusterAdmin(t.brokerList, t.kafkaConf)
	if err != nil {
		return nil, err
	}
	// defer admin.Close()

	return admin, nil
}

func (t *TopicManager) CreateTopic(admin sarama.ClusterAdmin, name string) (error) {
	logger.Info("t.numOfReplicationFactor", zap.Any("t.numOfReplicationFactor", t.numOfReplicationFactor))
	err := admin.CreateTopic(name, &sarama.TopicDetail{
		NumPartitions: t.numOfPartitions,
		ReplicationFactor: t.numOfReplicationFactor,
	}, false)
	 
	if err != nil {
		return err
	}

	defer admin.Close()

	return nil
}

func (t *TopicManager) DeleteTopic(admin sarama.ClusterAdmin, name string) error {
	err := admin.DeleteTopic(name)
	if err != nil {
		return err
	}

	defer admin.Close()

	return nil
}