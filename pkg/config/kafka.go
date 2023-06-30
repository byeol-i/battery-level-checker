package config

import (
	"flag"
	"strings"

	"github.com/Shopify/sarama"
)

var (
	brokerList        = flag.String("brokerList", "kafka-1:9094", "List of brokers to connect")
	// brokerList        = flag.String("brokerList", "kafka-1:9094,kafka-2:9094,kafka-3:9094", "List of brokers to connect")
	topic             = flag.String("topic", "device_event", "Topic name")
	partition         = flag.String("partition", "2", "Partition number")
	offsetType        = flag.Int("offsetType", -1, "Offset Type (OffsetNewest | OffsetOldest)")
	messageCountStart = flag.Int("messageCountStart", 0, "Message counter start from:")
	maxRetry   		  = flag.Int("maxRetry", 5, "Retry limit")
)

func init() {
	flag.Parse()
}

func (c *KafkaConfigManager) GetKafkaSarama() *sarama.Config {
	config := sarama.NewConfig()

	return config
}

func (c *KafkaConfigManager) GetBrokerList() []string {
	return strings.Split(*brokerList, ",")
}

func (c *KafkaConfigManager) GetMaxRetry() int {
	return *maxRetry
}

func (c *KafkaConfigManager) GetTopic() string {
	return *topic
}