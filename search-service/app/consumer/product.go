package consumer

import (
	"fmt"
	"os"
	"os/signal"
	"search-service/pkg/config"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type ProductConsumer struct {
}

func NewProductConsumer() *ProductConsumer {
	return &ProductConsumer{}
}

func (p *ProductConsumer) Consume() {
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
