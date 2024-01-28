package kafka

import (
	"inventory-service/pkg/logger"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// Delivery report handler for produced messages
func DeliveryReportHandler(producer *kafka.Producer) {
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					logger.Info("Delivery failed", "topic", ev.TopicPartition, "error", ev.TopicPartition.Error)
				} else {
					logger.Info("Message delivered", "topic", ev.TopicPartition)
				}
			}
		}
	}()
}
