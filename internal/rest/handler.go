package rest

import (
	"net/http"

	"dyegocruz.com.br/api-my-resume/internal/modules/resume"
	"github.com/gin-gonic/gin"
)

type MyResumeHandler struct {
	service resume.MyResumeService
}

func NewMyResumeHandler(service resume.MyResumeService) *MyResumeHandler {
	return &MyResumeHandler{service: service}
}

func (h *MyResumeHandler) GetByUsername(c *gin.Context) {

	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username is required",
		})
		return
	}

	resume, err := h.service.GetByUsername(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get resume",
		})
		return
	}

	c.JSON(http.StatusOK, resume)
}
