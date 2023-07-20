package config

import (
	"flag"
	"sync"

	"github.com/byeol-i/battery-level-checker/pkg/logger"
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

			configManagerInstance = &configManager{}
		} 
	}

	return configManagerInstance
}