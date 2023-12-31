package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func GetMessage(topic string, value []byte) *kafka.Message {
	return &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: value,
	}
}
