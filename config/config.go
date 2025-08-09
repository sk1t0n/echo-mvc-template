package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Http Http
	}

	Http struct {
		Port string `env:"SERVER_PORT" env-default:"8080"`
	}
)

var config Config

func NewConfig() (*Config, error) {
	err := cleanenv.ReadConfig(".env", &config)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return &config, nil
}
