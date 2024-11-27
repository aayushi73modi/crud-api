package main

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)
type PostgresConfig struct {
	POSTGRES_HOST     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	POSTGRES_PORT     string `env:"POSTGRES_PORT" envDefault:"5433"`
	POSTGRES_USER     string `env:"POSTGRES_USER" envDefault:"root"`
	POSTGRES_PASSWORD string `env:"POSTGRES_PASSWORD" envDefault:"root"`
	POSTGRES_DB       string `env:"POSTGRES_DB" envDefault:"studentpostgres"`
}
func LoadPostgresConfig() (*PostgresConfig, error) {
	var config PostgresConfig
	if err := env.Parse(&config); err != nil {
		return nil, fmt.Errorf("failed to load PostgresConfig: %w", err)
	}
	return &config, nil
}