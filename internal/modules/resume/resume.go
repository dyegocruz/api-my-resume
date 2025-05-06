package resume

import (
	"fmt"
	"log"

	"dyegocruz.com.br/api-my-resume/internal/models"
	"dyegocruz.com.br/api-my-resume/internal/repositories"
)

type myResumeService struct {
	repo repositories.MyResumeRepository
}

func NewMyResumeService(repo repositories.MyResumeRepository) MyResumeService {
	return &myResumeService{repo: repo}
}

func (s *myResumeService) GetByUsername(username string) (*models.MyResume, error) {
	resume, err := s.repo.FindByUsername(username)
	if err != nil {
		log.Println("Error getting resume:", err)
		return nil, fmt.Errorf("failed to get resume: %w", err)
	}

	return &resume, nil
}
