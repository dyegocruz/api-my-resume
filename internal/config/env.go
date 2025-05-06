package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func IsProduction(environment string) bool {
	return environment == "production"
}

type Config struct {
	App struct {
		Environment string `env:"GO_ENV" envDefault:"development"`
		GinMode     string `env:"GIN_MODE" envDefault:"debug"`
	}
	MongoDB struct {
		MongoURI string `env:"MONGO_URI"     envDefault:"mongodb://localhost:27017/"`
		Database string `env:"MONGO_DB_NAME" envDefault:"api-my-resume-dev"`
	}
}

func FromEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Couldn't load .env file: %v", err)
	}

	var c Config

	err = env.Parse(&c)
	if err != nil {
		log.Println("Error parsing environment variables:", err)
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}

	gin.SetMode(c.App.GinMode)

	return &c, nil
}
