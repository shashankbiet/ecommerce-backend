package consumer

import (
	"encoding/json"
	"os"
	"os/signal"
	"search-service/app/dao"
	model "search-service/app/models"
	"search-service/pkg/config"
	"search-service/pkg/db"
	"search-service/pkg/logger"
	searchpb "search-service/proto/search"
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

func (pc *ProductConsumer) Consume() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.GetConfig().KafkaConfig.Servers,
		"group.id":          config.GetConfig().KafkaConfig.GroupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		logger.Error("failed to create Kafka product consumer", "error", err)
		panic(err)
	}
	defer c.Close()

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
				pc.updateProduct(e.Value)
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

func (pc *ProductConsumer) updateProduct(data []byte) {
	var product model.Product
	err := json.Unmarshal(data, &product)
	if err != nil {
		logger.Error("failed to decode product update", "error", err)
		return
	}

	productPb := &searchpb.Product{
		Id:          int32(product.Id),
		Name:        product.Name,
		Description: product.Description,
		Brand:       product.Brand,
		Category:    product.Category,
		SubCategory: product.SubCategory,
	}

	if err := pc.datastore.UpsertProduct(int(product.Id), productPb); err != nil {
		logger.Error("failed to upsert product", "product_id", product.Id, "error", err)
	}
}
