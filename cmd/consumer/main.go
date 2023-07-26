package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/config"
	"github.com/byeol-i/battery-level-checker/pkg/consumer"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	manager := config.GetInstance()
	saramaConfig := manager.GetKafkaSarama()
	brokers := manager.GetBrokerList()
	
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

	topics := make([]string, 0)
	
    rwMutex := new(sync.RWMutex)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer wg.Done()
		for {

			newClient, err := sarama.NewClient(brokers, saramaConfig)
			if err != nil {
				log.Panic(err)
			}
			availableTopics, err := newClient.Topics()
			if err != nil {
				log.Panic(err)
			}
			filteredTopics := consumer.FilterTopicsByPatterns(availableTopics)
			
			if (!consumer.CompareTopics(topics, filteredTopics)) {
				rwMutex.Lock()

				topics = nil
				topics = make([]string, len(filteredTopics))
				copy(topics, filteredTopics)
				
				rwMutex.Unlock()
			}
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		duration, _ := time.ParseDuration("1s")
		ticker := time.NewTicker(duration)
		defer ticker.Stop()
		defer wg.Done()


		rwMutex.RLock()
		subscribeTopics := make([]string, len(topics))
		copy(subscribeTopics, topics)
		rwMutex.RUnlock()
		
		if (len(subscribeTopics) != 0) {
			go consumer.KeepConsume(ctx, topics, client)
		} 

		for {
			<-ticker.C

			if (!consumer.CompareTopics(subscribeTopics, topics)) {
				cancel()

				rwMutex.RLock()
				subscribeTopics = nil
				subscribeTopics = make([]string, len(topics))
				copy(subscribeTopics, topics)
				rwMutex.RUnlock()
				if (len(subscribeTopics) != 0) {
					logger.Info("Find new Topics, restart consume func", zap.Any("topic's count", len(subscribeTopics)))
					ctx, cancel = context.WithCancel(context.Background())

                	go consumer.KeepConsume(ctx, topics, client)
				} 
			}
		}
		
	}() 
	wg.Wait()
	log.Printf("Keep running!...")
	select {}
}

