package rest_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"dyegocruz.com.br/api-my-resume/internal/models"
	"dyegocruz.com.br/api-my-resume/internal/rest"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock do serviço MyResumeService
type MockMyResumeService struct {
	mock.Mock
}

func (m *MockMyResumeService) GetByUsername(username string) (*models.MyResume, error) {
	args := m.Called(username)
	if args.Get(0) != nil {
		return args.Get(0).(*models.MyResume), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestGetByUsername_Success(t *testing.T) {
	// Arrange
	mockService := new(MockMyResumeService)
	handler := rest.NewMyResumeHandler(mockService)

	expectedResume := &models.MyResume{
		Username: "testuser",
		Name:     "Test User",
	}

	mockService.On("GetByUsername", "testuser").Return(expectedResume, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/resume", handler.GetByUsername)

	req, _ := http.NewRequest(http.MethodGet, "/resume?username=testuser", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
	assert.Contains(t, w.Body.String(), `"username":"testuser"`)
	assert.Contains(t, w.Body.String(), `"name":"Test User"`)
}

func TestGetByUsername_MissingUsername(t *testing.T) {
	// Arrange
	mockService := new(MockMyResumeService)
	handler := rest.NewMyResumeHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/resume", handler.GetByUsername)

	req, _ := http.NewRequest(http.MethodGet, "/resume", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"Username is required"`)
}

func TestGetByUsername_ServiceError(t *testing.T) {
	// Arrange
	mockService := new(MockMyResumeService)
	handler := rest.NewMyResumeHandler(mockService)

	mockService.On("GetByUsername", "testuser").Return(nil, errors.New("service error"))

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/resume", handler.GetByUsername)

	req, _ := http.NewRequest(http.MethodGet, "/resume?username=testuser", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"error":"Failed to get resume"`)
}
