package config

import (
	"os"
)

type Config struct {
	APP         string
	Environment string
	Server      struct {
		Host string
		Port string
	}
}

func NewConfig() (*Config, error) {
	var config Config

	config.APP = getEnv("APP", "app")
	config.Environment = getEnv("ENVIRONMENT", "develop")

	config.Server.Host = getEnv("SERVER_HOST", "localhost")
	config.Server.Port = getEnv("SERVER_PORT", ":4221")

	return &config, nil
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
