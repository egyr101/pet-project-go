package config

import (
	"time"
)

type ServerConfig struct {
	Port         string        `envconfig:"HTTP_PORT" reqired:"true"`
	ReadTimeout  time.Duration `envconfig:"HTTP_READ_TIMEOUT" default:"5s"`
	WriteTimeout time.Duration `envconfig:"HTTP_WRITE_TIMEOUT" default:"5s"`
}
