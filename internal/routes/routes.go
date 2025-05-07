package routes

import (
	"dyegocruz.com.br/api-my-resume/internal/config"
	"dyegocruz.com.br/api-my-resume/internal/modules/resume"
	"dyegocruz.com.br/api-my-resume/internal/repositories"
	"dyegocruz.com.br/api-my-resume/internal/rest"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, cfg *config.Config) {

	db, err := config.ConnectMongo(cfg)
	if err != nil {
		panic("failed to connect database")
	}

	config.EnsureIndexes(db, cfg)

	// Initialize layers
	myResumeRepo := repositories.NewResumeRepository(db, cfg)
	myResumeService := resume.NewMyResumeService(myResumeRepo)
	myResumeHandler := rest.NewMyResumeHandler(myResumeService)

	router.GET("/resume", myResumeHandler.GetByUsername)
}
