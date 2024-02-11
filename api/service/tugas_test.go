package service

import (
	respons "Todo/api/service/Respons"
	"Todo/mocks"
	"Todo/models"
	"errors"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestTugas_GetAll(t *testing.T) {
	mockRepo := new(mocks.TugasRepository)
	tugasService := NewTugasService(mockRepo)
	assert.NotNil(t, tugasService)
	// assert.Nil(t, tugasService)

	mockData := []models.Tugas{
		{
			Model: gorm.Model{
				ID:        2,
				CreatedAt: time.Now(),
			},
			Task:        "Sample Task",
			Level:       "medium",
			Deadline:    "tomorrow",
			Description: "Sample Description",
			Status:      true,
		},
	}

	mockRepo.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(mockData)

	result := tugasService.GetAll(nil, 1, 10)

	assert.NotNil(t, result)
	assert.Len(t, result, len(mockData))

	mockRepo.AssertExpectations(t)

	mockRepo.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(nil)
}

// func TestTugasServiceImpl_Create(t *testing.T) {
// 	// Create a mock repository
// 	mockRepo := new(mocks.TugasRepository) // Replace with your actual repository type
// 	tugasService := NewTugasService(mockRepo)

// 	mockData := respons.CreateTugasRespon{
// 		ID:          0,
// 		Task:        "",
// 		Level:       "",
// 		Deadline:    "",
// 		Description: "",
// 		Status:      false,
// 		// Created_at:  time.Now(),
// 	}

// 	// Create a dummy gin.Context
// 	mockContext := new(gin.Context)

// 	// Mock the ShouldBindJSON call on gin.Context and set expectations
// 	mockRepo.On("ShouldBindJSON", mock.Anything).Return(nil).Once()
// 	// Additional assertions or modifications to the gin.Context as needed

// 	// Mock the repository call and set expectations
// 	mockRepo.On("Create", mock.Anything).Return(mockData, nil).Once()

// 	// Call the Create method on your service with the mock context
// 	result, err := tugasService.Create(mockContext)

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.NotNil(t, result)
// 	assert.EqualValues(t, mockData, result)

// 	// Verify that the repository method was called with the expected input
// 	mockRepo.AssertExpectations(t)

// 	// Additional assertions based on your requirements
// 	assert.NoError(t, err)

// 	// Mock the repository call to simulate an error
// 	mockRepo.On("Create", mock.Anything).Return(respons.CreateTugasRespon{}, errors.New("mocked error"))

// 	// Call the Create method on your service
// 	resultWithError, err := tugasService.Create(mockContext)

// 	// Additional assertions based on your requirements for the error case
// 	assert.Error(t, err)
// 	assert.EqualValues(t, respons.CreateTugasRespon{}, resultWithError)
// }

func TestTugasServiceImpl_GetById(t *testing.T) {
	// Create a mock repository
	mockRepo := new(mocks.TugasRepository) // Replace with your actual repository type
	tugasService := NewTugasService(mockRepo)

	// Mock data for request and response
	mockID := "1"
	mockResult := models.Tugas{
		Model: gorm.Model{
			ID: 1,
		},
		Task:        "Sample Task",
		Level:       "medium",
		Deadline:    "tomorrow",
		Description: "Sample Description",
		Status:      true,
	}

	// Create a dummy gin.Context with a parameter
	mockContext, _ := gin.CreateTestContext(nil)
	// assert.Nil(t,errc)

	mockContext.Params = append(mockContext.Params, gin.Param{Key: "id", Value: mockID})

	// Mock the repository call and set expectations
	mockRepo.On("GetById", mock.AnythingOfType("uint")).Return(mockResult, nil)

	// Call the GetById method on your service with the mock context
	result, err := tugasService.GetById(mockContext)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Verify that the repository method was called with the expected input
	mockRepo.AssertExpectations(t)

	// Additional assertions based on your requirements
	assert.IsType(t, respons.GetIdTugasRespon{}, result)
	assert.EqualValues(t, mockResult.ID, result.(respons.GetIdTugasRespon).ID)
	assert.EqualValues(t, mockResult.Task, result.(respons.GetIdTugasRespon).Task)
	assert.EqualValues(t, mockResult.Task, result.(respons.GetIdTugasRespon).Task)
	assert.EqualValues(t, mockResult.Task, result.(respons.GetIdTugasRespon).Task)
	assert.EqualValues(t, mockResult.Task, result.(respons.GetIdTugasRespon).Task)
}

func TestTugas_Update(t *testing.T) {
	// Create a mock repository
	mockRepo := new(mocks.TugasRepository)
	tugasService := NewTugasService(mockRepo)

	// Mock context with ID parameter
	mockID := "1"
	mockContext, _ := gin.CreateTestContext(nil)
	mockContext.Params = append(mockContext.Params, gin.Param{Key: "id", Value: mockID})

	// Mock data for GetById
	mockTugas := models.Tugas{
		Model: gorm.Model{
			ID: 1,
		},
		Task:        "Sample Task",
		Level:       "medium",
		Deadline:    "tomorrow",
		Description: "Sample Description",
		Status:      true,
	}
	mockRepo.On("GetById", mock.AnythingOfTypeArgument("uint")).Return(mockTugas, nil)

	// Mock data for Update
	mockUpdatedTugas := models.Tugas{
		Model: gorm.Model{
			ID: 1,
		},
		Task:        "Updated Task",
		Level:       "high",
		Deadline:    "next week",
		Description: "Updated Description",
		Status:      false,
	}
	mockRepo.On("Update", mock.AnythingOfType("models.Tugas")).Return(mockUpdatedTugas, nil)

	// Call the Update method
	result, err := tugasService.Update(mockContext)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)

	// Verify mock calls
	mockRepo.AssertExpectations(t)
}

func TestUpdate_InvalidID(t *testing.T) {
	// Mock repository
	mockRepo := new(mocks.TugasRepository) // Replace with your actual repository type

	tugasService := NewTugasService(mockRepo)

	// Mock context
	mockContext, _ := gin.CreateTestContext(nil)
	mockContext.Params = append(mockContext.Params, gin.Param{Key: "id", Value: "invalid"})

	// Create TugasServiceImpl instance

	// Call Update method
	result, _ := tugasService.Update(mockContext)

	// Assertions
	// assert.Error(t, err)
	assert.Nil(t, result)

	// Verify mock calls
	mockRepo.AssertExpectations(t)
}

func TestUpdate_GetByIDError(t *testing.T) {
	// Mock repository
	mockRepo := new(mocks.TugasRepository) // Replace with your actual repository type

	// Mock context
	mockContext, _ := gin.CreateTestContext(nil)
	mockContext.Params = append(mockContext.Params, gin.Param{Key: "id", Value: "1"})

	// Mock error for GetById
	getIDError := errors.New("error getting tugas")
	mockRepo.On("GetById", uint(1)).Return(models.Tugas{}, getIDError)

	// Create TugasServiceImpl instance
	tugasService := NewTugasService(mockRepo)

	// Call Update method
	result, err := tugasService.Update(mockContext)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)

	// Verify mock calls
	mockRepo.AssertExpectations(t)
}

// func TestTugas_Delete(t *testing.T) {

// }
// func TestTugas_GetByStatus(t *testing.T) {

// }
// func TestTugas_GetByLevel(t *testing.T) {

// }
// func TestTugas_GetByDeadline(t *testing.T) {

// }
