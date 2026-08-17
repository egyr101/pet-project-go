package config

import "time"

type PostgresConfig struct {
	DSN             string `envconfig:"DATABASE_URL" required:"true"`
	MaxConns        int    `envconfig:"DB_MAX_CONNS" default:"20"`
	MinConns        int    `envconfig:"DB_MIN_CONNS" default:"2"`
	MaxConnLifetime time.Duration `envconfig:"DB_MAX_CONN_LIFETIME" default:"1h"`
}