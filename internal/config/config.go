package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Cfg struct {
	DbHost string `env:"DB_HOST,required"`
}

func Load() (Cfg, error) {
	_ = godotenv.Load()

	var cfg Cfg

	if err := env.Parse(&cfg); err != nil {
		return Cfg{}, err
	}

	return cfg, nil
}
