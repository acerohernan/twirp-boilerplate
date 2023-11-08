package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	Logger *LoggerConfig
}

type LoggerConfig struct {
	Level string
}

func NewConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		Port: os.Getenv("PORT"),
		Logger: &LoggerConfig{
			Level: os.Getenv("LOG_LEVEL"),
		},
	}
}
