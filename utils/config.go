package utils

import "github.com/spf13/viper"

// Config stores all the configuration of the application
type Config struct {
	ProjectId      string `mapstructure:"PROJECT_ID"`
	CollectionName string `mapstructure:"COLLECTION_NAME"`
	ServerPort     string `mapstructure:"SERVER_PORT"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}