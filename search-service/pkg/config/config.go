package config

// Viper uses the mapstructure package under the hood for unmarshaling values, so we use the mapstructure tags to specify the name of each config field.
type Configuration struct {
	Environment        string              `mapstructure:"ENVIRONMENT"`
	AppName            string              `mapstructure:"APP_NAME"`
	ElasticSearch      ElasticSearchConfig `mapstructure:"ELASTIC_SEARCH"`
	SearchService      SearchServiceConfig `mapstructure:"SEARCH_SERVICE"`
	KafkaConfig        KafkaConfig         `mapstructure:"KAFKA_CONFIG"`
	ProductUpdateTopic KafkaTopicConfig    `mapstructure:"PRODUCT_UPDATE_TOPIC"`
}

type KafkaConfig struct {
	Server               string `mapstructure:"SERVER"`
	SearchServiceGroupId string `mapstructure:"SEARCH_SERVICE_GROUP_ID"`
}

type KafkaTopicConfig struct {
	Name            string `mapstructure:"NAME"`
	ClientId        string `mapstructure:"CLIENT_ID"`
	GroupId         string `mapstructure:"GROUP_ID"`
	AutoOffsetReset string `mapstructure:"AUTO_OFFSET_RESET"`
}

type SearchServiceConfig struct {
	GrpcPort int `mapstructure:"GRPC_PORT"`
}

type ElasticSearchConfig struct {
	Address   string `mapstructure:"ADDRESS"`
	Username  string `mapstructure:"USERNAME"`
	Password  string `mapstructure:"PASSWORD"`
	IndexName string `mapstructure:"INDEX_NAME"`
	TimeoutMs int    `mapstructure:"TIMEOUT_MS"`
}
