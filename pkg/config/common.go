package config

import (
	"flag"
	"sync"

	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"go.uber.org/zap"
)

var lock = &sync.Mutex{}

type configManager struct {
	ApidConfig
	FirebaseConfig
	GrpcConfig
	KafkaConfig
	DBConfig
}

var configManagerInstance *configManager

func GetInstance() *configManager {
	if configManagerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		
		if configManagerInstance == nil {
			logger.Info("Creating configManager instance")
			flag.Parse()
			logger.Info("Config -brokerList", zap.Any("brokerList", brokerList))
			logger.Info("Config -numOfReplicationFactor", zap.Any("numOfReplicationFactor", numOfReplicationFactor))
			logger.Info("Config -numOfPartitions", zap.Any("numOfPartitions", numOfReplicationFactor))
	

			configManagerInstance = &configManager{}
		} 
	}

	return configManagerInstance
}