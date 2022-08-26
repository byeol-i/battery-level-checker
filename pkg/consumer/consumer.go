package consumer

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/aglide100/battery-level-checker/pkg/config"
)

func GetTopics() {
	time.Sleep(20 * time.Second)
	saramaConfig := config.GetKafkaSarama()
	
	saramaConfig.Consumer.Return.Errors = true
	brokers := config.GetBrokerList()
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