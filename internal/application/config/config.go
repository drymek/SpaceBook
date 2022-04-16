package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
)

type Config interface {
	GetNewRelicConfigAppName() string
	GetNewRelicConfigLicense() string
	GetTimeout() time.Duration
	GetHttpAddr() string
}

type config struct {
	NewRelicConfigAppName string        `env:"NEW_RELIC_CONFIG_APP_NAME" envDefault:""`
	NewRelicConfigLicense string        `env:"NEW_RELIC_CONFIG_LICENSE" envDefault:""`
	Timeout               time.Duration `env:"TIMEOUT" envDefault:"3s"`
	HttpAddr              string        `env:"HTTP_ADDR" envDefault:"0.0.0.0:8080"`
}

func (c config) GetHttpAddr() string {
	return c.HttpAddr
}

func (c config) GetTimeout() time.Duration {
	return c.Timeout
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
