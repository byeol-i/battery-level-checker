package consumer

import (
	"github.com/Shopify/sarama"
)

func (c *Consumer) GetAdmin() (sarama.ClusterAdmin, error){
	admin, err := sarama.NewClusterAdmin(c.brokerList, c.kafkaConf)
	if err != nil {
		return nil, err
	}
	// defer admin.Close()

	return admin, nil
}

func (c *Consumer) CreateTopic(admin sarama.ClusterAdmin, name string) (error) {
	err := admin.CreateTopic(name, &sarama.TopicDetail{
		NumPartitions: 1,
		ReplicationFactor: 1,
	}, false)
	 
	if err != nil {
		return err
	}

	defer admin.Close()

	return nil
}

func (c *Consumer) DeleteTopic(admin sarama.ClusterAdmin, name string) error {
	err := admin.DeleteTopic(name)
	if err != nil {
		return err
	}

	defer admin.Close()

	return nil
}