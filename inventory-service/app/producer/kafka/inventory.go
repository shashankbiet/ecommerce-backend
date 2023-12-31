package kafka

import (
	"inventory-service/pkg/kafka"
	"inventory-service/pkg/logger"
	"sync"
)

var onceInventoryProducer sync.Once
var kafkaInventoryProducer *KafkaInventoryProducer

type KafkaInventoryProducer struct{}

func GetKafkaInventoryProducer() *KafkaInventoryProducer {
	onceInventoryProducer.Do(func() {
		initKafkaInventoryProducer()
	})
	return kafkaInventoryProducer
}

func initKafkaInventoryProducer() {
	kafkaInventoryProducer = &KafkaInventoryProducer{}
}

func (*KafkaInventoryProducer) Publish(value []byte) {
	topic := "inventory_update"
	producer := kafka.GetKafkaProducer()
	kafka.DeliveryReportHandler(producer)
	message := kafka.GetMessage(topic, value)
	err := producer.Produce(message, nil)
	if err != nil {
		logger.Log.Error("error in publishing inventory message", "error", err)
	} else {
		logger.Log.Info("successfully published inventory message")
	}
	producer.Flush(15 * 1000)
}
