package db

import (
	"search-service/pkg/config"
	"sync"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
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
		panic("error in initializing elasticsearch")
	}
	singletonElasticsearchClient = elasticsearchClient
}
