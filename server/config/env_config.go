package config

import (
	"log"

	"github.com/spf13/viper"
)

type envConfig struct {
	ApiKey string `mapstructure:"API_KEY"`
}

func LoadConfig(path string) (config envConfig) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}
