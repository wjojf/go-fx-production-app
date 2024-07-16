package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

func FromEnv() (*Config, error) {

	godotenv.Load()

	if cfgInstance != nil {
		return cfgInstance, nil
	}

	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	cfgInstance = &cfg
	return cfgInstance, nil
}
