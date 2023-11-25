package crdbclient

import (
	"fmt"
	"time"

	"example/configure"
)

type Config struct {
	Host                  string        `env:"DB_HOST"`
	Port                  int           `env:"DB_PORT"`
	Database              string        `env:"DB_DATABASE"`
	User                  string        `env:"DB_USER"`
	Password              string        `env:"DB_PASSWORD"`
	MaxConnections        int           `env:"DB_MAX_CONNECTIONS,default=15"`
	MaxConnectionLifeTime time.Duration `env:"DB_MAX_CONNECTION_LT,default=5m"`
	MaxConnectionIdleTime time.Duration `env:"DB_MAX_CONNECTION_IT,default=1m"`
	QueryTimeout          time.Duration `env:"DB_QUERY_TIMEOUT,default=3s"`
}

func (c *Config) ToConnString() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		c.User, c.Password, c.Host, c.Port, c.Database)
}

func LoadConfig() *Config {
	cfg := &Config{}

	err := configure.LoadFromEnv(cfg)
	if err != nil {
		panic(fmt.Errorf("failed to configure.LoadFromEnv: %w", err))
	}

	return cfg
}
