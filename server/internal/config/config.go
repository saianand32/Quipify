package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/saianand32/Quipify/internal/env"
)

type Config struct {
	Port string
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port: env.GetString("PORT", "8080"),
	}
}
