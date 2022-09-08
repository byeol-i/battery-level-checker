package config

import (
	"github.com/Shopify/sarama"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	brokerList        = kingpin.Flag("brokerList", "List of brokers to connect").Default("kafka-1:9094", "kafka-2:9094").Strings()
	topic             = kingpin.Flag("topic", "Topic name").Default("important").String()
	partition         = kingpin.Flag("partition", "Partition number").Default("0").String()
	offsetType        = kingpin.Flag("offsetType", "Offset Type (OffsetNewest | OffsetOldest)").Default("-1").Int()
	messageCountStart = kingpin.Flag("messageCountStart", "Message counter start from:").Int()
	maxRetry   = kingpin.Flag("maxRetry", "Retry limit").Default("5").Int()
)


func GetKafkaSarama() *sarama.Config {
	kingpin.Parse()
	config := sarama.NewConfig()

	return config
}

func GetBrokerList() []string {
	kingpin.Parse()

	return *brokerList
}

func GetMaxRetry() int {
	kingpin.Parse()

	return *maxRetry
}

func GetTopic() string {
	kingpin.Parse()

	return *topic
}

