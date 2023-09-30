package consumer

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"sort"

	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/config"
)
var (
	patterns = []string{"^battery_device_(.*?)_(.*)", "^battery_user_(.*?)_(.*)"}
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

func FilterTopicsByPatterns(topics []string) []string {
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

func CompareTopics(a, b []string) bool {
	if (len(a) != len(b)) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	return reflect.DeepEqual(a, b)
}