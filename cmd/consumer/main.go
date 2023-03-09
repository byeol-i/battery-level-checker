package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

var (
	brokerList = flag.String("brokerList", "kafka-1:9092", "List of brokers to connect")
	topic = flag.String("topic", "important", "Topic name")
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
	consumer, err := master.ConsumePartition(*topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Panic(err)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Println(err)
			case msg := <-consumer.Messages():
				*messageCountStart++
				log.Println("Received messages", string(msg.Key), string(msg.Value))
			case <-signals:
				log.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
	log.Println("Processed", *messageCountStart, "messages")
}
