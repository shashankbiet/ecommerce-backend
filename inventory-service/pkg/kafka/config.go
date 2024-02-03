package kafka

import (
	"inventory-service/pkg/config"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var kafkaConfigMap *kafka.ConfigMap
var onceKafkaConfig sync.Once

func getKafkaConfigMap() *kafka.ConfigMap {
	onceKafkaConfig.Do(func() {
		initKafkaConfigMap()
	})
	return kafkaConfigMap
}

func initKafkaConfigMap() {
	kafkaConfigMap = &kafka.ConfigMap{
		"bootstrap.servers": config.GetConfig().KafkaConfig.BootstrapServers,
	}
}
