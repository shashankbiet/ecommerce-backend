package consumer

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"search-service/app/dao"
	model "search-service/app/models"
	"search-service/pkg/config"
	"search-service/pkg/db"
	"search-service/pkg/logger"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type ProductConsumer struct {
	datastore dao.IProductDataStore
}

func NewProductConsumer() *ProductConsumer {
	elasticSearchClient := db.GetElasticSearchConnection()
	indexName := config.GetConfig().ElasticSearch.IndexName
	if indexName == "" {
		panic("ElasticSearch index name is not set")
	}
	datastore := dao.NewProductDataStore(elasticSearchClient, indexName)
	return &ProductConsumer{
		datastore: datastore,
	}
}

func NewProductConsumerWithDataStore(datastore dao.IProductDataStore) *ProductConsumer {
	return &ProductConsumer{datastore: datastore}
}

func (pc *ProductConsumer) Consume() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.GetConfig().KafkaConfig.Server,
		"group.id":          config.GetConfig().KafkaConfig.SearchServiceGroupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		logger.Error("failed to create Kafka product consumer", "error", err)
		panic(err)
	}
	defer func() { _ = c.Close() }()

	err = c.SubscribeTopics([]string{config.GetConfig().ProductUpdateTopic.Name}, nil)
	if err != nil {
		logger.Error("failed to subscribe to product updates", "error", err)
		panic(err)
	}

	// Set up a signal handler to handle graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true

	for run {
		select {
		case sig := <-sigchan:
			logger.Info("stopping product consumer", "signal", sig)
			run = false
		default:
			ev := c.Poll(100)
			switch e := ev.(type) {
			case *kafka.Message:
				if err := pc.updateProduct(e.Value); err != nil {
					logger.Error("failed to process product update", "error", err)
				}
			case kafka.Error:
				if e.IsFatal() || e.Code() == kafka.ErrAllBrokersDown {
					logger.Error("fatal Kafka consumer error", "error", e)
					run = false
				}
			default:
				// Ignored events, such as partitions assignment/revocation
			}
		}
	}

}

func (pc *ProductConsumer) updateProduct(data []byte) error {
	normalized, err := model.NormalizeJSON(data)
	if err != nil {
		return fmt.Errorf("normalize product update: %w", err)
	}
	var product model.Product
	if err := json.Unmarshal(normalized, &product); err != nil {
		return fmt.Errorf("decode product update: %w", err)
	}
	if product.Id <= 0 {
		return fmt.Errorf("invalid product id %d", product.Id)
	}
	if !product.Status.IsValid() {
		return fmt.Errorf("invalid product status %d", product.Status)
	}
	if err := pc.datastore.UpsertProduct(product.Id, &product); err != nil {
		return fmt.Errorf("upsert product %d: %w", product.Id, err)
	}
	return nil
}
