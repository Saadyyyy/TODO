package service

import (
	mocks "Todo/api/service/mocks"
	"Todo/models"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestTugas_GetAll(t *testing.T) {
	// import mock tugas repository
	mockRepo := new(mocks.TugasRepository)
	// mengasign mock tugas repository menjadi service
	tugasService := NewTugasService(mockRepo)

	//membuat mocking data
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
	// membuat testcase
	t.Run("Success", func(t *testing.T) {
		mockRepo.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(mockData)

		//memanggil function get all
		result := tugasService.GetAll(nil, 1, 10)

		assert.NotNil(t, result)
		assert.Len(t, result, len(mockData))

		//cek apakah repository get all benar
		mockRepo.AssertExpectations(t)
	})
}

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
	mockContext, _ := gin.CreateTestContext(nil)

	t.Run("Succes", func(t *testing.T) {
		mockContext.Params = append(mockContext.Params, gin.Param{Key: "id", Value: mockID})
		mockRepo.On("GetById", mockResult.ID).Return(mockResult, nil)

		// Call the GetById method on your service with the mock context
		result, _ := tugasService.GetById(mockContext)

		// Assertions
		// assert.Nil(t, err)
		assert.NotNil(t, result)

		// Verify that the repository method was called with the expected input
		mockRepo.AssertExpectations(t)

		// Additional assertions based on your requirements
		// assert.IsType(t, respons.GetIdTugasRespon{}, result)
		assert.EqualValues(t, mockResult.ID, result.ID)
		assert.EqualValues(t, mockResult.Task, result.Task)
		assert.EqualValues(t, mockResult.Task, result.Task)
		assert.EqualValues(t, mockResult.Task, result.Task)
		assert.EqualValues(t, mockResult.Task, result.Task)
	})

	// t.Run("Failed Param", func(t *testing.T) {
	// 	// Set up the mock context with an invalid ID parameter
	// 	mockContext.Params = append(mockContext.Params, gin.Param{Key: "id", Value: "InvalidID"})

	// 	// Call the GetById method on your service with the mock context
	// 	_, err := tugasService.GetById(mockContext)

	// 	// Assertions
	// 	assert.Nilf(t, err, "InvalidID") // Ensure that an error is returned
	// 	// mockRepo.AssertExpectations(t)

	// })

}

func Test_Update(t *testing.T) {
	mockRepo := new(mocks.TugasRepository)
	tugasService := NewTugasService(mockRepo)
	mockID := "1"
	mockResult := models.Tugas{
		Model: gorm.Model{
			ID:        1,
			UpdatedAt: time.Now(),
		},
		Task:        "Sample Task",
		Level:       "medium",
		Deadline:    "tomorrow",
		Description: "Sample Description",
		Status:      true,
	}

	mockContext, _ := gin.CreateTestContext(nil)
	mockContext.Params = append(mockContext.Params, gin.Param{Key: "id", Value: mockID})

	mockRepo.On("GetById", mockResult.ID).Return(mockResult, nil).Once()

	mockRepo.On("Update").Return(mockResult, nil).Once()

	defer mockRepo.AssertExpectations(t)

	result, _ := tugasService.Update(mockContext)

	// Assert that no error occurred
	// assert.NoError(t, err)

	// // Assert that result is not nil
	assert.Nil(t, result)

	assert.EqualValues(t, mockResult.ID, result.ID)
	assert.EqualValues(t, mockResult.Task, result.Task)
	assert.EqualValues(t, mockResult.Level, result.Level)
	assert.EqualValues(t, mockResult.Deadline, result.Deadline)
	assert.EqualValues(t, mockResult.Description, result.Description)
	assert.EqualValues(t, mockResult.Status, result.Status)
	assert.EqualValues(t, mockResult.UpdatedAt, result.Update_at)
}
