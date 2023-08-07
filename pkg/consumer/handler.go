package consumer

import (
	"fmt"
	"regexp"

	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/device"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	cacheSvc "github.com/byeol-i/battery-level-checker/pkg/svc/cache"

	dbSvc "github.com/byeol-i/battery-level-checker/pkg/svc/db"
	"go.uber.org/zap"
)

type MessageHandler struct{
	cacheClient *cacheSvc.CacheSvcClient
	dbClient  *dbSvc.DBSvcClient
}

func NewMessageHandler(cacheClient *cacheSvc.CacheSvcClient, dbClient *dbSvc.DBSvcClient) *MessageHandler {
	return &MessageHandler{
		cacheClient: cacheClient,
		dbClient: dbClient,
	}
}

// Impl Setup
func (h *MessageHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Impl Cleanup
func (h *MessageHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// Impl ConsumeClaim
func (h *MessageHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {	
	topic := claim.Topic()

	userId, deviceId, err := extractUUIDs(topic)
	if err != nil {
		logger.Error("Can't extract uuid", zap.Error(err))
	}

	for message := range claim.Messages() {
		err := h.cacheClient.CallWriteMsg(userId, deviceId, message.Value)
		if err != nil {
			logger.Error("Can't write to cache srv", zap.Error(err))
		}

		bl, err := device.ParseFromJson(string(message.Value[:]))
		if err != nil {
			logger.Error("Can't parse from json", zap.Error(err))
		}

		err = h.dbClient.CallUpdateBatteryLevel(deviceId, userId, bl)
		if err != nil {
			logger.Error("Can't write to db srv", zap.Error(err))
		}
		
		session.MarkMessage(message, "") 
	}

	return nil
}

func extractUUIDs(input string) (uuid1, uuid2 string, err error) {
	re := regexp.MustCompile(`__([^_]+)`)
	matches := re.FindAllStringSubmatch(input, -1)
	if len(matches) <  2 {
		return "", "", fmt.Errorf("invalid input format")
	}

	uuid1 = matches[0][1]
	uuid2 = matches[1][1]
	
	return uuid1, uuid2, nil
}
