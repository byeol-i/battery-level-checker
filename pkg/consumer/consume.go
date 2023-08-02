package consumer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/config"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	cacheSvc "github.com/byeol-i/battery-level-checker/pkg/svc/cache"
	"go.uber.org/zap"
)

func KeepConsume(ctx context.Context, topics []string, client sarama.ConsumerGroup, cacheSvcAddr, dbSvcAddr string) {
	handler := &MessageHandler{
		cacheClient: cacheSvc.NewCacheSvcClient(cacheSvcAddr),
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("Context has something!", ctx.Err())
			return
		default:
			if err := client.Consume(ctx, topics, handler); err != nil {
				log.Printf("Error from consumer: %v", err)
			}
		}
	}
}

func GetTopics() {
	time.Sleep(20 * time.Second)
	manager := config.GetInstance()
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

func ConsumeLatestMessage(topic string) error {
	manager := config.GetInstance()
	saramaConfig := manager.GetKafkaSarama()

	client, err := sarama.NewClient(manager.GetBrokerList(), saramaConfig)
	if err != nil {
		logger.Error("Can't create client", zap.Error(err))
		return err
	}

	defer client.Close()

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		logger.Error("Can't get Consumer from client", zap.Error(err))
		return err
	}

	defer consumer.Close()

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		logger.Error("Can't get partitions", zap.Error(err))
		return err
	}

	for _, partition := range partitions {
		offset, err := client.GetOffset(topic, partition, sarama.OffsetNewest)
		if err != nil {
			logger.Error("Can't get offset", zap.Error(err))
			return err
		}

		pc, err := consumer.ConsumePartition(topic, partition, offset)
		if err != nil {
			logger.Error("Can't create ConsumePartition", zap.Error(err))
			return err
		}

		go func(pc sarama.PartitionConsumer) {
			defer pc.Close()

			for message := range pc.Messages() {
				logger.Info("Consume msg", zap.Any("topic", message.Topic), zap.Any("offset", message.Offset), zap.Any("value", message.Value))
			}
		}(pc)
	}
	
	return nil
}