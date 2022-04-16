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
	GetDatabase() Database
}

type Database interface {
	GetDatabaseHost() string
	GetDatabasePort() int
	GetDatabaseUser() string
	GetDatabasePassword() string
	GetDatabaseName() string
}

type config struct {
	NewRelicConfigAppName string        `env:"NEW_RELIC_CONFIG_APP_NAME" envDefault:""`
	NewRelicConfigLicense string        `env:"NEW_RELIC_CONFIG_LICENSE" envDefault:""`
	Timeout               time.Duration `env:"TIMEOUT" envDefault:"3s"`
	HttpAddr              string        `env:"HTTP_ADDR" envDefault:"0.0.0.0:8080"`
	Database              Database
}

func (c config) GetDatabase() Database {
	return c.Database
}

type database struct {
	DatabaseHost     string `env:"DATABASE_HOST" envDefault:"localhost"`
	DatabasePort     int    `env:"DATABASE_PORT" envDefault:"5432"`
	DatabaseUser     string `env:"DATABASE_USER" envDefault:"sb"`
	DatabasePassword string `env:"DATABASE_PASSWORD" envDefault:"password"`
	DatabaseName     string `env:"DATABASE_NAME" envDefault:"sb"`
}

func (d database) GetDatabaseHost() string {
	return d.DatabaseHost
}

func (d database) GetDatabasePort() int {
	return d.DatabasePort
}

func (d database) GetDatabaseUser() string {
	return d.DatabaseUser
}

func (d database) GetDatabasePassword() string {
	return d.DatabasePassword
}

func (d database) GetDatabaseName() string {
	return d.DatabaseName
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

	d := database{}
	if err := env.Parse(&d, opts...); err != nil {
		return c, fmt.Errorf("cannot parse database config: %w", err)
	}
	c.Database = d
	return c, nil
}
