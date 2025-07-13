package config

import (
	"fmt"
	"os"
)

type Config struct {
	pgConn string
}

func (c Config) PgConn() string {
	return c.pgConn
}

func NewConfig() *Config {
	return &Config{
		pgConn: fmt.Sprintf(
			"user=%s password=%s host=%s dbname=%s sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
		),
	}
}
