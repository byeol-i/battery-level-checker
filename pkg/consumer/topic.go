package consumer

import (
	"fmt"
	"log"
	"regexp"

	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/config"
)
var (
	patterns = []string{"battery_device____"}
)

type topicManager struct {
	client sarama.Client
	group sarama.ConsumerGroup
	wantedTopics chan string
}

func NewTopicManager() (*topicManager, error) {
	manager := config.GetInstance()

	saramaConfig := manager.GetKafkaSarama()
	brokers := manager.GetBrokerList()

	newClient, err := sarama.NewClient(brokers, saramaConfig)
	if err != nil {
		log.Panic(err)
	}

	// availableTopics, err := newClient.Topics()
	// if err != nil {
	// 	return nil, err
	// }

	client, err := sarama.NewConsumerGroupFromClient(manager.GetConsumerGroup(), newClient)
	if err != nil {
		log.Panic(err)
	}

	return &topicManager{
		client: newClient,
		group: client,
		wantedTopics: make(chan string),
	}, nil
}


func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func remove(slice []string, str string) []string {
	index := -1
	for i, s := range slice {
		if s == str {
			index = i
			break
		}
	}
	if index == -1 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
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