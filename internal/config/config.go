package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

const (
	envFileName = ".env"
	envPrefix   = ""
)

type Config struct {
	Server    ServerConfig
	Petfinder PetfinderConfig
	Gyeonggi  GyeonggiConfig
}

type ServerConfig struct {
	Port string `envconfig:"SERVER_PORT" default:"8080"`
	Host string `envconfig:"SERVER_HOST" default:"localhost"`
}

type PetfinderConfig struct {
	ClientID     string `envconfig:"PETFINDER_CLIENT_ID" required:"true"`
	ClientSecret string `envconfig:"PETFINDER_CLIENT_SECRET" required:"true"`
	AccessToken  string `envconfig:"PETFINDER_ACCESS_TOKEN" required:"true"`
}

type GyeonggiConfig struct {
	GyeonggiApiKey string `envconfig:"GYEONGGI_API_KEY" required:"true"`
}

var Cfg Config

func Load() {
	_ = godotenv.Load(envFileName)
	if err := envconfig.Process(envPrefix, &Cfg); err != nil {
		os.Exit(1)
	}
	log.Info().Msg("successfully loaded configs")
}
