package main

import (
	"context"
	"flag"
	"log"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/consumer"
)

var (
	group = flag.String("group", "my-consumer-group", "using for consumer group")
	brokerList = flag.String("brokerList", "kafka-1:9092", "List of brokers to connect")
	topic = flag.String("topic", "device_event", "Topic name")
	partition = flag.Int("partition", 0, "Partition number")
	offsetType = flag.Int("offsetType", -1, "Offset Type (OffsetNewest | OffsetOldest)")
	messageCountStart = flag.Int("messageCountStart", 0, "Message counter start from:")
)

func main() {
	flag.Parse()
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	brokers := []string{*brokerList}
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := master.Close(); err != nil {
			log.Panic(err)
		}
	}()
	
	newClient, err := sarama.NewClient(brokers, config)
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
			if err := client.Consume(ctx, []string{*topic}, handler); err != nil {
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


// func createTopic()