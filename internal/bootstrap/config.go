package bootstrap

import (
	"errors"
	"os"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type Config struct {
	Port        int    `env:"PORT"`
	PostgresURL string `env:"POSTGRES_URL"`
}

func NewConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	var c Config
	err = env.Parse(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
