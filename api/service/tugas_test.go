package service

import (
	"Todo/api/service/mocksRepo"
	"Todo/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func Test_TugasGetAll(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	service := NewTugasService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		mockData := []models.Tugas{
			{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: time.Now(),
				},
				Task:        "ayam",
				Level:       "easy",
				Deadline:    "besok",
				Description: "apalah dia ni",
				Status:      true,
			},
		}
		mockRepo.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(mockData)
		result := service.GetAll(nil, 1, 2)
		assert.NotNil(t, result)
		assert.Len(t, result, len(mockData))

		mockRepo.AssertExpectations(t)
	})

}

func Test_CreateTugas(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	serv := NewTugasService(mockRepo)

	t.Run("Succes", func(t *testing.T) {
		data := models.Tugas{
			Model: gorm.Model{
				ID:        2,
				CreatedAt: time.Now(),
			},
			Task:        "mm",
			Level:       "mm",
			Deadline:    "mm",
			Description: "apalah dia ni",
			Status:      true,
		}

		mockRepo.On("Created", mock.AnythingOfType("*models.Tugas")).Return(&data, nil)
		result, err := serv.Create(&data)

		assert.NotNil(t, result)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})
}
