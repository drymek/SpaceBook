package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config interface {
	GetNewRelicConfigAppName() string
	GetNewRelicConfigLicense() string
}

type config struct {
	NewRelicConfigAppName string `env:"NEW_RELIC_CONFIG_APP_NAME" envDefault:""`
	NewRelicConfigLicense string `env:"NEW_RELIC_CONFIG_LICENSE" envDefault:""`
}

func (c config) GetNewRelicConfigAppName() string {
	return c.NewRelicConfigAppName
}

func (c config) GetNewRelicConfigLicense() string {
	return c.NewRelicConfigLicense
}

func NewConfig(opts ...env.Options) (Config, error) {
	c := config{}
	if err := env.Parse(&c, opts...); err != nil {
		return c, fmt.Errorf("cannot parse main config: %w", err)
	}

	return c, nil
}
