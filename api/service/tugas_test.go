package service

import (
	respons "Todo/api/service/Respons"
	"Todo/mocks"
	"Todo/models"
	"errors"
	"net/http/httptest"
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
		// Add other fields as needed
	}

	// Create a dummy gin.Context with a parameter
	mockContext, _ := gin.CreateTestContext(nil)

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
		// Add other fields as needed
	}
	// Create dummy gin.Context with parameters
	mockContext, _ := gin.CreateTestContext(nil)
	mockContext.Params = append(mockContext.Params, gin.Param{Key: "id", Value: mockID})

	// Case 1: Error while converting ID
	mockContext.Params[0].Value = "not_an_integer"
	result, err := tugasService.Update(mockContext)
	// assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "strconv.Atoi: parsing \"not_an_integer\": invalid syntax")

	// Case 2: Error while retrieving data from repository
	expectedError := errors.New("some repository error")
	mockRepo.On("GetById", mock.AnythingOfType("uint")).Return(models.Tugas{}, expectedError)
	result, err = tugasService.Update(mockContext)
	// assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, expectedError.Error())

	// Case 3: Success updating the data
	mockRepo.On("GetById", mock.AnythingOfType("uint")).Return(mockResult, nil)
	mockContext.Request = httptest.NewRequest("PUT", "tugas/", nil)
	mockContext.ShouldBindJSON(&mockResult)
	mockRepo.On("Update", mockResult).Return(mockResult, nil)

	result, err = tugasService.Update(mockContext)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Verify that the repository methods were called with the expected input
	mockRepo.AssertExpectations(t)

	mockRepo.AssertExpectations(t)
	assert.IsType(t, respons.GetIdTugasRespon{}, result)
	assert.EqualValues(t, mockResult.ID, result.(respons.GetIdTugasRespon).ID)
	assert.EqualValues(t, mockResult.Task, result.(respons.GetIdTugasRespon).Task)
	assert.EqualValues(t, mockResult.Task, result.(respons.GetIdTugasRespon).Task)
	assert.EqualValues(t, mockResult.Task, result.(respons.GetIdTugasRespon).Task)
	assert.EqualValues(t, mockResult.Task, result.(respons.GetIdTugasRespon).Task)

}

// func TestTugas_Delete(t *testing.T) {

// }
// func TestTugas_GetByStatus(t *testing.T) {

// }
// func TestTugas_GetByLevel(t *testing.T) {

// }
// func TestTugas_GetByDeadline(t *testing.T) {

// }
