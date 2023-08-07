package config

import (
	"flag"
	"strings"

	"github.com/Shopify/sarama"
)

var (
	group = flag.String("group", "battery_consumer", "using for consumer group")
	brokerList        = flag.String("brokerList", "kafka-1:9094,kafka-2:9094", "List of brokers to connect")
	// brokerList        = flag.String("brokerList", "kafka-1:9094,kafka-2:9094,kafka-3:9094", "List of brokers to connect")
	// topic             = flag.String("topic", "device_event", "Topic name")
	// partition         = flag.String("partition", "2", "Partition number")
	// offsetType        = flag.Int("offsetType", -1, "Offset Type (OffsetNewest | OffsetOldest)")
	// messageCountStart = flag.Int("messageCountStart", 0, "Message counter start from:")
	
	numOfReplicationFactor = flag.Int("numOfReplicationFactor", 2, "numOfReplicationFactor")
	numOfPartitions = flag.Int("numOfPartitions", 2, "numOfPartitions")
	maxRetry   		  = flag.Int("maxRetry", 5, "Retry limit")
)

type KafkaConfigImpl interface {
	GetConsumerGroup() string
	GetKafkaSarama() *sarama.Config
	GetBrokerList() []string
	GetMaxRetry() int
	GetNumOfPartitions() int
	GetNumOfReplicationFactor() int
	GetTopic() string
}

type KafkaConfig struct {
	KafkaConfigImpl
}

func (c KafkaConfig) GetNumOfPartitions() int32 {
	return int32(*numOfPartitions)
}

func (c KafkaConfig) GetNumOfReplicationFactor() int16 {
	return int16(*numOfReplicationFactor)
}

func (c KafkaConfig) GetConsumerGroup() string {
	return *group
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