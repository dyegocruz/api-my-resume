package main

import (
	"time"

	"dyegocruz.com.br/api-my-resume/internal/config"
	"dyegocruz.com.br/api-my-resume/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.FromEnv()

	r := gin.Default()

	AllowOrigins := []string{"https://dyegocruz.com.br"}

	if !config.IsProduction(cfg.App.Environment) {
		AllowOrigins = []string{"http://localhost", "http://localhost:3000", "http://localhost:3003"}
	}

	// Configure o middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     AllowOrigins,
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.Setup(r, cfg)

	r.Run(":8080")
}
