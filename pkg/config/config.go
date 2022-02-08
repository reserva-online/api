package config

import (
	"os"
)

type Config struct {
	DatabaseDriver           string
	DatabaseConnectionString string
	MigrationsPath           string
}

func NewConfig() Config {
	return Config{
		DatabaseDriver:           os.Getenv("DATABASE_DRIVER"),
		DatabaseConnectionString: os.Getenv("DATABASE_CONNECTION_STRING"),
		MigrationsPath:           os.Getenv("MIGRATION_PATH"),
	}
}
