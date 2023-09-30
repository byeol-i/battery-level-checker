package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/Shopify/sarama"
	pb_svc_cache "github.com/byeol-i/battery-level-checker/pb/svc/cache"
	cacheManager "github.com/byeol-i/battery-level-checker/pkg/cache"
	"github.com/byeol-i/battery-level-checker/pkg/config"
	"github.com/byeol-i/battery-level-checker/pkg/consumer"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	server "github.com/byeol-i/battery-level-checker/pkg/svc/cache"
	"github.com/byeol-i/battery-level-checker/pkg/topic"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {
	if err := realMain(); err != nil {
		log.Printf("err :%s", err)
		os.Exit(1)
	}
}

func realMain() error {
	configManager := config.GetInstance()
	saramaConfig := configManager.GetKafkaSarama()
	brokers := configManager.GetBrokerList()
	
	gRPCL, err := net.Listen("tcp", configManager.GrpcConfig.GetCacheAddr())
	if err != nil {
		return err
	}
	defer gRPCL.Close()

	var opts []grpc.ServerOption

	deviceCache := cache.New(2*time.Hour, 1*time.Hour)
	userCache := cache.New(2*time.Hour, 1*time.Hour)

	cManager := cacheManager.NewCacheManager(deviceCache, userCache)
	grpcServer := grpc.NewServer(opts...)
	
	cacheSrv := server.NewCacheServiceServer(cManager)
	pb_svc_cache.RegisterCacheServer(grpcServer, cacheSrv)

	wg, _ := errgroup.WithContext(context.Background())
	
	wg.Go(func() error {
		logger.Info("Starting grpc server..." + configManager.GrpcConfig.GetCacheAddr())
		err := grpcServer.Serve(gRPCL)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
			return err
		}

		return nil
	})

	TopicManager := topic. NewTopicManager()
	

	wg.Go(func() error {
		duration, _ := time.ParseDuration("600s")
		ticker := time.NewTicker(duration)
		defer ticker.Stop()

		newClient, err := sarama.NewClient(brokers, saramaConfig)
		if err != nil {
			log.Panic(err)
		}
		availableTopics, err := newClient.Topics()
		if err != nil {
			log.Panic(err)
		}
		filteredTopics := consumer.FilterTopicsByPatterns(availableTopics)
		
		for _, val := range filteredTopics {
			userId, deviceId, err := consumer.ExtractUUIDs(val)
			if err != nil {
				logger.Error("Can't extract uuid", zap.Error(err))
			}
			_, found := deviceCache.Get(deviceId)
			if !found {
				admin, err := TopicManager.GetAdmin()
				if err != nil {
					logger.Error("Can't get admin", zap.Error(err))
				}
				log.Printf("delete %s", deviceId)
				err = TopicManager.DeleteTopic(admin, "battery_device_"+userId+"_"+deviceId)
				if err != nil {
					logger.Error("Can't delete topic", zap.Error(err))
				}
			}
		}
		return nil
	})

	return wg.Wait()
}
