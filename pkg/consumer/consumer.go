package consumer

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/config"
)

func GetTopics() {
	time.Sleep(20 * time.Second)
	manager := config.NewKafkaConfigManager()
	saramaConfig := manager.GetKafkaSarama()
	
	saramaConfig.Consumer.Return.Errors = true
	brokers := manager.GetBrokerList()
	master, err := sarama.NewConsumer(brokers, saramaConfig)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := master.Close(); err != nil {
			log.Panic(err)
		}
	}()

	topics, _ := master.Topics()
	for index := range topics {
		fmt.Println(topics[index])
	}
}