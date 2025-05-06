package resume

import "dyegocruz.com.br/api-my-resume/internal/models"

type MyResumeService interface {
	GetByUsername(username string) (*models.MyResume, error)
}
