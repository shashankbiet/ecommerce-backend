package kafka

import (
	"inventory-service/pkg/kafka"
	"inventory-service/pkg/logger"
	"sync"
)

var (
	onceProductProducer  sync.Once
	kafkaProductProducer *KafkaProductProducer
)

type KafkaProductProducer struct{}

func GetKafkaProductProducer() *KafkaProductProducer {
	onceProductProducer.Do(func() {
		initKafkaProductProducer()
	})
	return kafkaProductProducer
}

func initKafkaProductProducer() {
	kafkaProductProducer = &KafkaProductProducer{}
}

func (*KafkaProductProducer) Publish(value []byte) {
	topic := "product_update"
	producer := kafka.GetKafkaProducer()
	kafka.DeliveryReportHandler(producer)
	message := kafka.GetMessage(topic, value)
	err := producer.Produce(message, nil)
	if err != nil {
		logger.Log.Error("error in publishing product message", "error", err)
	} else {
		logger.Log.Info("successfully published product message")
	}
	producer.Flush(15 * 1000)
}
