package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type MongoEnv struct {
	MongoURI string `json:"mongo_uri"`
	Database string `json:"db_name"`
}

func getEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

func IsProduction() bool {
	return getEnv("GO_ENV") == "production"
}

func GetMongoEnv() MongoEnv {
	return MongoEnv{
		MongoURI: getEnv("MONGO_URI"),
		Database: getEnv("DB_NAME"),
	}
}
