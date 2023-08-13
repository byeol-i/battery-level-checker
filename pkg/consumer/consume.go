package consumer

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
	cacheSvc "github.com/byeol-i/battery-level-checker/pkg/svc/cache"
	dbSvc "github.com/byeol-i/battery-level-checker/pkg/svc/db"
)

func KeepConsume(ctx context.Context, topics []string, client sarama.ConsumerGroup, cacheSvcAddr, dbSvcAddr string) {
	handler := &MessageHandler{
		cacheClient: cacheSvc.NewCacheSvcClient(cacheSvcAddr),
		dbClient: dbSvc.NewDBSvcClient(dbSvcAddr),
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("Context has err!", ctx.Err())
			return
		default:
			if err := client.Consume(ctx, topics, handler); err != nil {
				log.Printf("Error from consumer: %v", err)
			}
		}
	}
}