package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `json:"port" mapstructure:"port"`
	} `json:"server" mapstructure:"server"`
	Otel struct {
		ServiceName string  `json:"service_name" mapstructure:"service_name"`
		SampleRatio float64 `json:"sample_ratio" mapstructure:"sample_ratio"`
		MetricsPort int     `json:"metrics_port" mapstructure:"metrics_port"`
	} `json:"otel" mapstructure:"otel"`
	Secrets struct {
		DBPassword string `json:"db_password" mapstructure:"db_password"`
		APIKey     string `json:"api_key" mapstructure:"api_key"`
	} `json:"secrets" mapstructure:"secrets"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	var config Config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
		return nil, err
	}

	return &config, nil
}
