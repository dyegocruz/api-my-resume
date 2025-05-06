package config

import (
	"github.com/caarlos0/env/v6"
)

// type MongoEnv struct {
// 	MongoURI string `json:"mongo_uri"`
// 	Database string `json:"db_name"`
// }

// func getEnv(key string) string {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	return os.Getenv(key)
// }

func IsProduction(env string) bool {
	return env == "production"
}

// func GetMongoEnv() MongoEnv {
// 	return MongoEnv{
// 		MongoURI: getEnv("MONGO_URI"),
// 		Database: getEnv("MONGO_DB_NAME"),
// 	}
// }

type Config struct {
	App struct {
		Environment string `env:"GO_ENV" envDefault:"development"`
	}
	MongoDB struct {
		MongoURI string `env:"MONGO_URI"     envDefault:"mongodb://localhost:27017/"`
		Database string `env:"MONGO_DB_NAME" envDefault:"api-my-resume-dev"`
	}
}

func FromEnv() *Config {
	var c Config

	if err := env.Parse(&c.App); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.MongoDB); err != nil {
		panic(err)
	}

	return &c
}
