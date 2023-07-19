package config

import (
	"flag"
	"strings"

	"github.com/Shopify/sarama"
)

var (
	// brokerList        = flag.String("brokerList", "localhost:9094", "List of brokers to connect")
	
	brokerList        = flag.String("brokerList", "kafka-1:9094", "List of brokers to connect")
	// brokerList        = flag.String("brokerList", "kafka-1:9094,kafka-2:9094,kafka-3:9094", "List of brokers to connect")
	// topic             = flag.String("topic", "device_event", "Topic name")
	// partition         = flag.String("partition", "2", "Partition number")
	// offsetType        = flag.Int("offsetType", -1, "Offset Type (OffsetNewest | OffsetOldest)")
	// messageCountStart = flag.Int("messageCountStart", 0, "Message counter start from:")
	maxRetry   		  = flag.Int("maxRetry", 5, "Retry limit")
)

type KafkaConfigImpl interface {
	GetKafkaSarama() *sarama.Config
	GetBrokerList() []string
	GetMaxRetry() int
	GetTopic() string
}

type KafkaConfig struct {
	KafkaConfigImpl
}

func (c KafkaConfig) GetKafkaSarama() *sarama.Config {
	config := sarama.NewConfig()

	return config
}

func (c KafkaConfig) GetBrokerList() []string {
	return strings.Split(*brokerList, ",")
}

func (c KafkaConfig) GetMaxRetry() int {
	return *maxRetry
}

// func (c KafkaConfig) GetTopic() string {
// 	return *topic
// }