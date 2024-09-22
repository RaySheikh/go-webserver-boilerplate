package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config holds configuration variables
type Config struct {
	Server struct {
		Port int `json:"port"`
	} `json:"server"`

	OTel struct {
		ServiceName  string  `json:"service_name"`
		CollectorURL string  `json:"collector_url"`
		SampleRatio  float64 `json:"sample_ratio"`
	} `json:"otel"`

	Secrets struct {
		DBPassword string `json:"db_password"`
		APIKey     string `json:"api_key"`
	} `json:"secrets"`
}

// LoadConfig loads configuration from YAML and environment variables
func LoadConfig() (*Config, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	// Bind environment variables and check for errors
	if err := viper.BindEnv("secrets.db_password", "DB_PASSWORD"); err != nil {
		log.Fatalf("Error binding environment variable DB_PASSWORD: %v", err)
	}
	if err := viper.BindEnv("secrets.api_key", "API_KEY"); err != nil {
		log.Fatalf("Error binding environment variable API_KEY: %v", err)
	}

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
