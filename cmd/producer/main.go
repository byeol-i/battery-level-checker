package main

import (
	"log"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/config"
)

func main() {
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

	for i:=0; i<50; i++ {
		time.Sleep(1*time.Second)
		msg := &sarama.ProducerMessage{
			Topic: manager.GetTopic(),
			Headers: []sarama.RecordHeader{
				{
					Key: []byte("uid"), Value: []byte(strconv.Itoa(i)),
				},
			},
			Value: sarama.StringEncoder(time.Now().Format("01-02-2006 15:04:05")),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Panic(err)
		}
		log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", manager.GetTopic(), partition, offset)
	}
}


