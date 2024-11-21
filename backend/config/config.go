package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	LLMApiKey   string
	Port        string
}

func Load() *Config {
	// Load .env file if present
	godotenv.Load()

	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgresql://localhost:5432/gnome?sslmode=disable"),
		LLMApiKey:   getEnv("LLM_API_KEY", ""),
		Port:        getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
