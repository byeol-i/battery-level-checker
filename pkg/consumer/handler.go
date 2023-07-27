package consumer

import (
	"fmt"
	"regexp"

	"github.com/Shopify/sarama"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"go.uber.org/zap"
)

type MessageHandler struct{}

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

	logger.Info("looking ", zap.Any("topic", topic))
	userId, deviceId, err := extractUUIDs(topic)
	if err != nil {
		logger.Error("Can't extract uuid", zap.Error(err))
	}

	for message := range claim.Messages() {
		logger.Info("UserId", zap.Any("userID",userId))
		logger.Info("DeviceId", zap.Any("deviceId",deviceId))
		fmt.Printf("Received message: %s, offset : %d\n", string(message.Value), message.Offset)
		session.MarkMessage(message, "") 
	}

	return nil
}

func extractUUIDs(input string) (uuid1, uuid2 string, err error) {
	re := regexp.MustCompile(`(?<=__)[^_]+`)
	matches := re.FindStringSubmatch(input)
	if len(matches) != 3 {
		return "", "", fmt.Errorf("invalid input format")
	}
	uuid1 = matches[1]
	uuid2 = matches[2]
	return uuid1, uuid2, nil
}
