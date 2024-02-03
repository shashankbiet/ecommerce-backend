package config

// Viper uses the mapstructure package under the hood for unmarshaling values, so we use the mapstructure tags to specify the name of each config field.
type Configuration struct {
	Environment          string           `mapstructure:"ENVIRONMENT"`
	AppName              string           `mapstructure:"APP_NAME"`
	SqlConfig            SqlConfig        `mapstructure:"SQL_CONFIG"`
	HttpServer           HttpServerConfig `mapstructure:"HTTP_SERVER"`
	KafkaConfig          KafkaConfig      `mapstructure:"KAFKA_CONFIG"`
	ProductUpdateTopic   KafkaTopicConfig `mapstructure:"PRODUCT_UPDATE_TOPIC"`
	InventoryUpdateTopic KafkaTopicConfig `mapstructure:"INVENTORY_UPDATE_TOPIC"`
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

type HttpServerConfig struct {
	Port               int `mapstructure:"PORT"`
	ReadTimeoutMs      int `mapstructure:"READ_TIMEOUT_MS"`
	WriteTimeoutMs     int `mapstructure:"WRITE_TIMEOUT_MS"`
	IdleTimeoutMs      int `mapstructure:"IDLE_TIMEOUT_MS"`
	KeepAliveTimeoutMs int `mapstructure:"KEEP_ALIVE_TIMEOUT_MS"`
}

type KafkaConfig struct {
	BootstrapServers string `mapstructure:"BOOTSTRAP_SERVERS"`
}

type KafkaTopicConfig struct {
	Name            string `mapstructure:"NAME"`
	ClientId        string `mapstructure:"CLIENT_ID"`
	GroupId         string `mapstructure:"GROUP_ID"`
	AutoOffsetReset string `mapstructure:"AUTO_OFFSET_RESET"`
}
