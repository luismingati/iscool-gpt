package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	GeminiAPIKey      string
	Port              string
	RateLimitRequests int
	RateLimitWindow   time.Duration
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	rateLimitRequests := 10
	if rlr := os.Getenv("RATE_LIMIT_REQUESTS"); rlr != "" {
		if parsed, err := strconv.Atoi(rlr); err == nil {
			rateLimitRequests = parsed
		}
	}

	rateLimitWindow := 60 * time.Second
	if rlw := os.Getenv("RATE_LIMIT_WINDOW"); rlw != "" {
		if parsed, err := time.ParseDuration(rlw); err == nil {
			rateLimitWindow = parsed
		}
	}

	geminiAPIKey := os.Getenv("GEMINI_API_KEY")
	if geminiAPIKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable is required")
	}

	return &Config{
		GeminiAPIKey:      geminiAPIKey,
		Port:              port,
		RateLimitRequests: rateLimitRequests,
		RateLimitWindow:   rateLimitWindow,
	}
}
