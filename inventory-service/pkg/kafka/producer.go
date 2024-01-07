package kafka

import (
	"inventory-service/pkg/logger"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var kafkaProducer *kafka.Producer
var onceProducer sync.Once

func GetKafkaProducer() *kafka.Producer {
	onceProducer.Do(func() {
		initKafkaProducer()
	})
	return kafkaProducer
}

func initKafkaProducer() {
	kafkaConfigMap := getKafkaConfigMap()
	producer, err := kafka.NewProducer(kafkaConfigMap)
	if err != nil {
		logger.Log.Error("failed to create producer", "error", err)
		panic(err)
	}
	kafkaProducer = producer
}
