package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server    ServerConfig
	Petfinder PetfinderConfig
}

type ServerConfig struct {
	Port string
	Host string
}

type PetfinderConfig struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
}

var Cfg Config

func Load() {
	Cfg.Server = ServerConfig{
		Port: getEnv("SERVER_PORT", "8080"),
		Host: getEnv("SERVER_HOST", "localhost"),
	}
	Cfg.Petfinder = PetfinderConfig{
		ClientID:     getEnv("PETFINDER_CLIENT_ID", ""),
		ClientSecret: getEnv("PETFINDER_CLIENT_SECRET", ""),
		AccessToken:  getEnv("PETFINDER_ACCESS_TOKEN", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
