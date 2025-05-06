package repositories

import "dyegocruz.com.br/api-my-resume/internal/models"

type MyResumeRepository interface {
	FindByUsername(username string) (models.MyResume, error)
}
