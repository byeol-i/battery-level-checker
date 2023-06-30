package producer

import (
	"log"

	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/config"
)

func Write() {
	manager := config.NewKafkaConfigManager()

	saramaConfig := manager.GetKafkaSarama()
	
	saramaConfig.Consumer.Return.Errors = true
	brokers := manager.GetBrokerList()
	
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Retry.Max = manager.GetMaxRetry()
	saramaConfig.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, saramaConfig)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Panic(err)
		}
	}()
	msg := &sarama.ProducerMessage{
		Topic: manager.GetTopic(),
		Value: sarama.StringEncoder("Something Cool"),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", manager.GetTopic(), partition, offset)
}
