package consumer

import (
	"fmt"

	"github.com/Shopify/sarama"
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
	for message := range claim.Messages() {
		
		fmt.Printf("Received message: %s, offset : %d\n", string(message.Value), message.Offset)
		session.MarkMessage(message, "") 
	}

	return nil
}