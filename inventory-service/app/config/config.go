package config

// Viper uses the mapstructure package under the hood for unmarshaling values, so we use the mapstructure tags to specify the name of each config field.
type Configuration struct {
	Environment          string        `mapstructure:"ENVIRONMENT"`
	AppName              string        `mapstructure:"APP_NAME"`
	SqlConfig            SqlConfig     `mapstructure:"SQL_CONFIG"`
	HttpServer           HttpServer    `mapstructure:"HTTP_SERVER"`
	KafkaProductConsumer KafkaConsumer `mapstructure:"KAFKA_PRODUCT_CONSUMER"`
	ProductUpdateTopic   KafkaTopic    `mapstructure:"PRODUCT_UPDATE_TOPIC"`
}

// Config struct holds the mysql database configuration
type SqlConfig struct {
	Username     string `mapstructure:"USERNAME"`
	Password     string `mapstructure:"PASSWORD"`
	Host         string `mapstructure:"HOST"`
	Port         string `mapstructure:"PORT"`
	Database     string `mapstructure:"DATABASE"`
	MaxLifetime  int    `mapstructure:"MAX_LIFE_TIME"`
	MaxOpenConns int    `mapstructure:"MAX_OPEN_CONNS"`
	MaxIdleConns int    `mapstructure:"MAX_IDLE_CONNS"`
}

type HttpServer struct {
	Port               int `mapstructure:"PORT"`
	ReadTimeoutMs      int `mapstructure:"READ_TIMEOUT_MS"`
	WriteTimeoutMs     int `mapstructure:"WRITE_TIMEOUT_MS"`
	IdleTimeoutMs      int `mapstructure:"IDLE_TIMEOUT_MS"`
	KeepAliveTimeoutMs int `mapstructure:"KEEP_ALIVE_TIMEOUT_MS"`
}

type KafkaConsumer struct {
	Id               string `mapstructure:"ID"`
	Host             string `mapstructure:"HOST"`
	Username         string `mapstructure:"PRIMARY_API_KEY"`
	Password         string `mapstructure:"PRIMARY_API_SECRET"`
	Auth             string `mapstructure:"AUTH"`
	Retries          int    `mapstructure:"RETRIES"`
	RetryInterval    int    `mapstructure:"RETRY_INTERVAL"`
	DeadLettering    bool   `mapstructure:"DEAD_LETTERING"`
	DelayInMs        int    `mapstructure:"DELAY_IN_MS"`
	CollectMetrics   bool   `mapstructure:"COLLECT_METRICS"`
	RetryConcurrency int    `mapstructure:"RETRY_CONCURRENCY"`
}

type KafkaTopic struct {
	TopicName       string `mapstructure:"NAME"`
	ClientId        string `mapstructure:"CLIENT_ID"`
	GroupId         string `mapstructure:"GROUP_ID"`
	AutoOffsetReset string `mapstructure:"AUTO_OFFSET_RESET"`
}
