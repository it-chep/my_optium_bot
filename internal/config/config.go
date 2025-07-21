package config

import (
	"fmt"
	"os"
)

type Config struct {
	pgConn string
	BotConfig
}

type BotConfig struct {
	token      string `yaml:"token"`
	webhookURL string `yaml:"webhook"`
}

func (c Config) PgConn() string {
	return c.pgConn
}

func (c Config) Token() string {
	return c.token
}

func (c Config) WebhookURL() string {
	return c.webhookURL
}

func NewConfig() *Config {
	return &Config{
		BotConfig: BotConfig{
			webhookURL: os.Getenv("WEBHOOK_URL"),
			token:      os.Getenv("BOT_TOKEN"),
		},
		pgConn: fmt.Sprintf(
			"user=%s password=%s host=%s dbname=%s sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
		),
	}
}
