package config

import "errors"

var (
	ErrorEnvNotFound = errors.New(".env is not found")
	ErrorPortNotFound = errors.New("port is not found in .env")
	ErrorDatabaseUrlNotFound = errors.New("database url is not found in .env")
)