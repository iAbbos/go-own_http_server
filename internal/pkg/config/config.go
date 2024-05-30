package config

import (
	"flag"
	"os"
)

type Config struct {
	APP         string
	Environment string
	Server      struct {
		Host string
		Port string
	}
	FilesDir string
}

func NewConfig() (*Config, error) {
	var config Config

	config.APP = getEnv("APP", "app")
	config.Environment = getEnv("ENVIRONMENT", "develop")

	config.Server.Host = getEnv("SERVER_HOST", "localhost")
	config.Server.Port = getEnv("SERVER_PORT", ":4221")

	dir := flag.String("dir", "", "The path to the directory where the files are stored.")
	flag.Parse()

	config.FilesDir = *dir

	return &config, nil
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
