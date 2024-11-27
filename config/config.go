package config

import (
	"log"

	"github.com/caarlos0/env"
)

type PostgresConfig struct {
	POSTGRES_HOST     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	POSTGRES_PORT     string `env:"POSTGRES_PORT" envDefault:"5433"`
	POSTGRES_USER     string `env:"POSTGRES_USER" envDefault:"root"`
	POSTGRES_PASSWORD string `env:"POSTGRES_PASSWORD" envDefault:"root"`
	POSTGRES_DB       string `env:"POSTGRES_DB" envDefault:"studentpostgres"`
}

func LoadPostgresConfig() (*PostgresConfig, error) {
	cfg := &PostgresConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Printf("Failed to load PostgresConfig: %v", err)
		return nil, err
	}
	return cfg, nil
}
