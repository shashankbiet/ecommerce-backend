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
					logger.Error("Kafka message delivery failed", "topic", ev.TopicPartition, "error", ev.TopicPartition.Error)
				}
			}
		}
	}()
}
