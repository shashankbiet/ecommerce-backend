package config

import (
	"fmt"
	"os"
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
	viper.SetConfigName(getEnvironment()) // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./conf")         // path to look for the config file in
	viper.AutomaticEnv()                  // Viper will check for an environment variable any time a viper.Get request is made

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error in reading config file: %w", err))
	}

	// Viper unmarshal the loaded env variables into the struct
	if err := viper.Unmarshal(&singletonConfig); err != nil {
		panic(fmt.Errorf("error in config unmarshal: %w", err))
	}
}

func getEnvironment() string {
	env := os.Getenv("ENVIRONMENT")
	if len(env) == 0 {
		return "default"
	}
	return env
}
