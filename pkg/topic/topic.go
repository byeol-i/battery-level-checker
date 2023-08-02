package topic

import (
	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/config"
)

type Admin struct {
	kafkaConf *sarama.Config
	brokerList []string
}

func NewAdmin() *Admin {
	manager := config.GetInstance()

	return &Admin{
		kafkaConf: manager.GetKafkaSarama(),
		brokerList: manager.GetBrokerList(),
	}
}

func (c *Admin) GetAdmin() (sarama.ClusterAdmin, error){
	admin, err := sarama.NewClusterAdmin(c.brokerList, c.kafkaConf)
	if err != nil {
		return nil, err
	}
	// defer admin.Close()

	return admin, nil
}

func (c *Admin) CreateTopic(admin sarama.ClusterAdmin, name string) (error) {
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

func (c *Admin) DeleteTopic(admin sarama.ClusterAdmin, name string) error {
	err := admin.DeleteTopic(name)
	if err != nil {
		return err
	}

	defer admin.Close()

	return nil
}