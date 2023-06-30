package config

import (
	"flag"

	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/db"
)

type KafkaConfigManagerImpl interface {
	GetKafkaSarama() *sarama.Config
	GetBrokerList() []string
	GetMaxRetry() int
	GetTopic() string
}

type KafkaConfigManager struct {
	KafkaConfigManagerImpl
}

func NewKafkaConfigManager() *KafkaConfigManager {
	flag.Parse()
	return &KafkaConfigManager{}
}

type DBConfigManagerImpl interface {
	GetDBConfig() *db.DBConfig
}

type DBConfigManager struct {
	DBConfigManagerImpl
}

func NewDBConfigManager() *DBConfigManager{
	return &DBConfigManager{}
}