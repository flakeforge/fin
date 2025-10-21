package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		App      App
		Logger   Logger
		Postgres Postgres
	}

	App struct {
		Name    string `env:"APP_NAME" envDefault:"fin"`
		Version string `env:"APP_VERSION" envDefault:"1.0.0"`
		Port    string `env:"SERVER_PORT" envDefault:"8080"`
	}
	Logger struct {
		Level   string `env:"LOG_LEVEL" envDefault:"info"`
		LogFile string `env:"LOG_FILE" envDefault:"logs/app.log"`
		Pretty  bool   `env:"LOG_PRETTY" envDefault:"true"`
		Console bool   `env:"LOG_CONSOLE" envDefault:"true"`
	}
	Postgres struct {
		URL     string `env:"POSTGRES_URL,required"`
		PoolMax int    `env:"PG_POOL_MAX" envDefault:"10"`
	}
)

func New() (*Config, error) {
  config := &Config{}
	if err := env.Parse(config); err != nil {
    return nil, fmt.Errorf("failed to parse config: %w", err)
  }

  return config, nil
}
