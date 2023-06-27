package consumer

import (
	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/config"
)

func GetAdmin() (sarama.ClusterAdmin, error){
	brokerList := config.GetBrokerList()
	kafkaConf := config.GetKafkaSarama()
	admin, err := sarama.NewClusterAdmin(brokerList, kafkaConf)
	if err != nil {
		return nil, err
	}
	defer admin.Close()

	return admin, nil
}

func CreateTopic(admin sarama.ClusterAdmin, name string) (error) {
	err := admin.CreateTopic(name, &sarama.TopicDetail{
		NumPartitions: 1,
		ReplicationFactor: 1,
	}, false)
	 
	if err != nil {
		return err
	}

	return nil
}