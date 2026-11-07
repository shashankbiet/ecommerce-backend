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

// loadSecretFiles sets NAME from the file named by NAME_FILE when NAME is empty,
// so secrets can be mounted as files (Compose secrets) instead of plain env vars.
func loadSecretFiles() error {
	for _, kv := range os.Environ() {
		fileVar, path, found := strings.Cut(kv, "=")
		if !found || !strings.HasSuffix(fileVar, "_FILE") {
			continue
		}
		name := strings.TrimSuffix(fileVar, "_FILE")
		if os.Getenv(name) != "" {
			continue
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading %s: %w", fileVar, err)
		}
		if err := os.Setenv(name, strings.TrimSpace(string(content))); err != nil {
			return fmt.Errorf("setting %s: %w", name, err)
		}
	}
	return nil
}

func initConfig() {
	if err := loadSecretFiles(); err != nil {
		panic(fmt.Errorf("error in loading secret files: %w", err))
	}

	viper.SetConfigName("default")
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./conf") // path to look for the config file in
	viper.AutomaticEnv()          // Viper will check for an environment variable any time a viper.Get request is made

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
