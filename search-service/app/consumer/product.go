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
		panic(err)
	}
	defer c.Close()

	err = c.SubscribeTopics([]string{config.GetConfig().ProductUpdateTopic.Name}, nil)
	if err != nil {
		panic(err)
	}

	// Set up a signal handler to handle graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true

	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("Message on %s: %s\n", e.TopicPartition, string(e.Value))
				pc.updateProduct(e.Value)
				if e.Headers != nil {
					fmt.Printf("Headers: %v\n", e.Headers)
				}
			case kafka.Error:
				// Errors should generally be considered informational, the client will try to automatically recover.
				fmt.Fprintf(os.Stderr, "Error: %v\n", e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				// Ignored events, such as partitions assignment/revocation
			}
		}
	}

	fmt.Printf("Closing product consumer\n")
}

func (pc *ProductConsumer) updateProduct(data []byte) {
	var product model.Product
	err := json.Unmarshal(data, &product)
	if err != nil {
		fmt.Println("Error unmarshalling product data: ", err)
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

	pc.datastore.UpsertProduct(int(product.Id), productPb)
}
