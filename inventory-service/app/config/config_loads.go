package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var singletonConfig Configuration
var once sync.Once

func GetConfig() *Configuration {
	once.Do(initConfig)
	return &singletonConfig
}

func initConfig() {
	viper.SetConfigName("default")                   // name of config file (without extension)
	viper.SetConfigType("yaml")                      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("../inventory-service/conf") // path to look for the config file in
	viper.AutomaticEnv()                             // Viper will check for an environment variable any time a viper.Get request is made

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error in reading config file: %w", err))
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&singletonConfig); err != nil {
		panic(fmt.Errorf("error in config unmarshal: %w", err))
	}
}
