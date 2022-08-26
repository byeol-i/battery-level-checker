package producer

import (
	"log"

	"github.com/Shopify/sarama"
	"github.com/aglide100/battery-level-checker/pkg/config"
)

func Write() {
	saramaConfig := config.GetKafkaSarama()
	
	saramaConfig.Consumer.Return.Errors = true
	brokers := config.GetBrokerList()
	
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Retry.Max = config.GetMaxRetry()
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
		Topic: config.GetTopic(),
		Value: sarama.StringEncoder("Something Cool"),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", config.GetTopic(), partition, offset)
}
