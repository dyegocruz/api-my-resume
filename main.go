package main

import (
	"net/http"
	"time"

	"dyegocruz.com.br/api-my-resume/config"
	"dyegocruz.com.br/api-my-resume/resume"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	if config.IsProduction() {
		gin.SetMode("release")
	}
}

func main() {

	config.EnsureIndexes()

	r := gin.Default()

	AllowOrigins := []string{"https://dyegocruz.com.br"}

	if !config.IsProduction() {
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

	r.GET("/resume", func(c *gin.Context) {
		username := c.Query("username")

		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Username is required",
			})
			return
		}

		resume, err := resume.GetByUsername(username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to get resume",
			})
			return
		}

		c.JSON(http.StatusOK, resume)
	})
	r.Run(":8080")
}
