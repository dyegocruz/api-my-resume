package resume_test

import (
	"errors"
	"testing"

	"dyegocruz.com.br/api-my-resume/internal/models"
	"dyegocruz.com.br/api-my-resume/internal/modules/resume"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock do repositório MyResumeRepository
type MockMyResumeRepository struct {
	mock.Mock
}

func (m *MockMyResumeRepository) FindByUsername(username string) (models.MyResume, error) {
	args := m.Called(username)
	return args.Get(0).(models.MyResume), args.Error(1)
}

func TestGetByUsername_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockMyResumeRepository)
	service := resume.NewMyResumeService(mockRepo)

	expectedResume := models.MyResume{
		Username: "testuser",
		Name:     "Test User",
	}

	mockRepo.On("FindByUsername", "testuser").Return(expectedResume, nil)

	// Act
	resume, err := service.GetByUsername("testuser")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resume)
	assert.Equal(t, "testuser", resume.Username)
	assert.Equal(t, "Test User", resume.Name)
	mockRepo.AssertExpectations(t)
}

func TestGetByUsername_Error(t *testing.T) {
	// Arrange
	mockRepo := new(MockMyResumeRepository)
	service := resume.NewMyResumeService(mockRepo)

	mockRepo.On("FindByUsername", "nonexistentuser").Return(models.MyResume{}, errors.New("user not found"))

	// Act
	resume, err := service.GetByUsername("nonexistentuser")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resume)
	assert.EqualError(t, err, "failed to get resume: user not found")
	mockRepo.AssertExpectations(t)
}
