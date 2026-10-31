package db

import (
	"search-service/pkg/config"
	"sync"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"search-service/pkg/logger"
)

var (
	singletonElasticsearchClient *elasticsearch.Client
	once                         sync.Once
)

func GetElasticSearchConnection() *elasticsearch.Client {
	once.Do(func() {
		initElasticSearchConnection()
	})
	return singletonElasticsearchClient
}

func initElasticSearchConnection() {
	cfg := elasticsearch.Config{
		Addresses: []string{config.GetConfig().ElasticSearch.Address},
		// Username:  config.GetConfig().ElasticSearch.Username,
		// Password:  config.GetConfig().ElasticSearch.Password,
	}
	elasticsearchClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		logger.Error("failed to initialize Elasticsearch client", "error", err)
		panic(err)
	}
	singletonElasticsearchClient = elasticsearchClient
}
