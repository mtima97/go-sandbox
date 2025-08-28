package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Cfg struct {
	AppPort        string   `env:"APP_PORT,required"`
	AllowedOrigins []string `env:"ALLOWED_ORIGINS"`
	DbUser         string   `env:"DB_USER,required"`
	DbPass         string   `env:"DB_PASS,required"`
	DbHost         string   `env:"DB_HOST,required"`
	DbPort         string   `env:"DB_PORT,required"`
	DbName         string   `env:"DB_NAME,required"`
}

func Load() (Cfg, error) {
	_ = godotenv.Load()

	var cfg Cfg

	if err := env.Parse(&cfg); err != nil {
		return Cfg{}, err
	}

	return cfg, nil
}
