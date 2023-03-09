package config

import (
	"flag"
	"strings"

	"github.com/Shopify/sarama"
)

var (
	brokerList        = flag.String("brokerList", "kafka-1:9094,kafka-2:9094,kafka-3:9094", "List of brokers to connect")
	topic             = flag.String("topic", "important", "Topic name")
	partition         = flag.String("partition", "0", "Partition number")
	offsetType        = flag.Int("offsetType", -1, "Offset Type (OffsetNewest | OffsetOldest)")
	messageCountStart = flag.Int("messageCountStart", 0, "Message counter start from:")
	maxRetry   		  = flag.Int("maxRetry", 5, "Retry limit")
)

func init() {
	flag.Parse()
}

func GetKafkaSarama() *sarama.Config {
	config := sarama.NewConfig()

	return config
}

func GetBrokerList() []string {
	return strings.Split(*brokerList, ",")
}

func GetMaxRetry() int {
	return *maxRetry
}

func GetTopic() string {
	return *topic
}