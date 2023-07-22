package main

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/config"
	"github.com/byeol-i/battery-level-checker/pkg/consumer"
)

var (
	patterns = []string{"battery_device____", "battery_user__"}

)

// For testing
func main() {
	manager := config.GetInstance()
	saramaConfig := manager.GetKafkaSarama()
	brokers := manager.GetBrokerList()
	
	// batteryUserTopic := "battery_user__"
	// batteryDeviceTopic := "battery_device____.*"

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

	client, err := sarama.NewConsumerGroupFromClient(manager.GetConsumerGroup(), newClient)
	if err != nil {
		log.Panic(err)
	}

	availableTopics, err := newClient.Topics()
	if err != nil {
		log.Panic(err)
	}

	
	var mu sync.Mutex
	data := make([]string, 0)

	handler := &consumer.MessageHandler{}

	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for {
			mu.Lock() 
			filteredTopics := filterTopicsByPatterns(patterns, availableTopics)
			data = nil
			data = filteredTopics
			mu.Unlock()
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			mu.Lock()
			if (len(data) == 0) {
				continue
			}

			if err := client.Consume(ctx, data, handler); err != nil {
				log.Printf("Error from consumer: %v", err)
				cancel()
				return
			}
						
			if ctx.Err() != nil {
				log.Println(ctx.Err())
				return
			}
			mu.Unlock() 
			time.Sleep(2 * time.Second)
		}
	}()
	log.Printf("Keep running!...")
	select {}
}

func filterTopicsByPatterns(patterns []string, topics []string) []string {
	filtered := make([]string, 0)

	for _, str := range topics {
		for _, pattern := range patterns {
			regExp, err := regexp.Compile(pattern)
			if err != nil {
				fmt.Printf("Invalid pattern: %s - %v\n", pattern, err)
				continue
			}

			if regExp.MatchString(str) {
				filtered = append(filtered, str)
				break 
			}
		}
	}

	return filtered
}