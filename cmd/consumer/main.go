package main

import (
	"context"
	"flag"
	"log"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/config"
	"github.com/byeol-i/battery-level-checker/pkg/consumer"
)

var (
	group = flag.String("group", "my-consumer-group", "using for consumer group")
)

// For testing
func main() {
	manager := config.GetInstance()
	saramaConfig := manager.GetKafkaSarama()
	brokers := manager.GetBrokerList()
	topic := manager.GetTopic()

	saramaConfig.Consumer.Return.Errors = true

	master, err := sarama.NewConsumer(brokers, saramaConfig)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := master.Close(); err != nil {
			log.Panic(err)
		}
	}()
	
	newClient, err := sarama.NewClient(brokers, saramaConfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := sarama.NewConsumerGroupFromClient(*group, newClient)
	if err != nil {
		log.Panic(err)
	}

	handler := &consumer.MessageHandler{}

	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, []string{topic}, handler); err != nil {
				log.Printf("Error from consumer: %v", err)
				cancel()
				return
			}
			
			if ctx.Err() != nil {
				log.Println(ctx.Err())
				return
			}
		}
	}()
	log.Printf("Keep running!...")
	select {}
}
