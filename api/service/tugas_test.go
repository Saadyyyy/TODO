package service

import (
	respons "Todo/api/service/Respons"
	"Todo/api/service/mocksRepo"
	"Todo/models"
	"errors"
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

func Test_CreateTugasSuccess(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	serv := NewTugasService(mockRepo)

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
}

func Test_CreateTugasFailed(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	serv := NewTugasService(mockRepo)
	data := models.Tugas{
		Model: gorm.Model{
			ID:        3,
			CreatedAt: time.Now(),
		},
		Task:        "mm",
		Level:       "mm",
		Deadline:    "mm",
		Description: "apalah dia ni",
		Status:      true,
	}

	mockRepo.On("Created", mock.AnythingOfType("*models.Tugas")).Return(nil, errors.New("mock error"))
	_, err := serv.Create(&data)

	// assert.Nil(t, result)
	assert.NotNil(t, err)
	mockRepo.AssertCalled(t, "Created", &data)
}

func Test_GetByIdSucces(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	serv := NewTugasService(mockRepo)
	data := models.Tugas{
		Model: gorm.Model{
			ID: 3,
		},
	}

	mockRepo.On("GetById", data.ID).Return(&data, nil)
	result, err := serv.GetById(data.ID)
	assert.NotNil(t, result)
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func Test_GetByIdFailed(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	serv := NewTugasService(mockRepo)
	data := models.Tugas{
		Model: gorm.Model{
			ID: 3,
		},
	}

	mockRepo.On("GetById", data.ID).Return(nil, errors.New("mock id error"))
	result, err := serv.GetById(data.ID)
	assert.Nil(t, result)
	assert.NotNil(t, err)

	mockRepo.AssertExpectations(t)
}

func Test_UpdateSuccess(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	serv := NewTugasService(mockRepo)
	data := models.Tugas{
		Model: gorm.Model{
			ID: 3,
		},
	}
	var respon respons.UpdateTugasRespon
	mockRepo.On("Update", mock.AnythingOfType("*models.Tugas")).Return(&data, nil)
	mockRepo.On("GetById", data.ID).Return(&data, nil)

	result, err := serv.Update(data.ID, respon)

	assert.NotNil(t, result)
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func Test_UpdateFailedUpdate(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	serv := NewTugasService(mockRepo)
	data := models.Tugas{
		Model: gorm.Model{
			ID: 4,
		},
	}
	var respon respons.UpdateTugasRespon
	mockRepo.On("Update", mock.AnythingOfType("*models.Tugas")).Return(nil, errors.New("Mock id Error"))
	mockRepo.On("GetById", data.ID).Return(&data, nil)

	result, err := serv.Update(data.ID, respon)

	assert.Nil(t, result)
	assert.NotNil(t, err)

	mockRepo.AssertExpectations(t)
}

func Test_UpdateFailedId(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	serv := NewTugasService(mockRepo)
	data := models.Tugas{
		Model: gorm.Model{
			ID: 6,
		},
	}
	var respon respons.UpdateTugasRespon
	mockRepo.On("Update", mock.AnythingOfType("*models.Tugas")).Return(&data, nil).Maybe()

	mockRepo.On("GetById", data.ID).Return(nil, errors.New("mock id error"))

	result, err := serv.Update(data.ID, respon)

	assert.Nil(t, result)
	assert.NotNil(t, err)

	mockRepo.AssertExpectations(t)
}

func Test_TugasDeleteSuccess(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	serv := NewTugasService(mockRepo)
	data := models.Tugas{
		Model: gorm.Model{
			ID: 3,
		},
	}
	// var respon respons.UpdateTugasRespon
	mockRepo.On("Delete", mock.AnythingOfType("*models.Tugas")).Return(&data, nil)
	mockRepo.On("GetById", data.ID).Return(&data, nil)

	result, err := serv.Delete(data.ID)

	assert.NotNil(t, result)
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func Test_UpdateFailedDeletes(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	serv := NewTugasService(mockRepo)
	data := models.Tugas{
		Model: gorm.Model{
			ID: 4,
		},
	}
	mockRepo.On("Delete", mock.AnythingOfType("*models.Tugas")).Return(nil, errors.New("Mock id Error"))
	mockRepo.On("GetById", data.ID).Return(&data, nil)

	result, err := serv.Delete(data.ID)

	assert.Nil(t, result)
	assert.NotNil(t, err)

	mockRepo.AssertExpectations(t)
}

func Test_UpdateFailedDeletesId(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	serv := NewTugasService(mockRepo)
	data := models.Tugas{
		Model: gorm.Model{
			ID: 7,
		},
	}
	mockRepo.On("Delete", mock.AnythingOfType("*models.Tugas")).Return(&data, nil).Maybe()
	mockRepo.On("GetById", data.ID).Return(nil, errors.New("mock erorr delete id"))

	result, err := serv.Delete(data.ID)

	assert.Nil(t, result)
	assert.NotNil(t, err)

	mockRepo.AssertExpectations(t)
}

func Test_TugasGetByStatusSuccess(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	serv := NewTugasService(mockRepo)
	status := true
	mockData := []*models.Tugas{
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
	mockRepo.On("GetByStatus", status, mock.Anything, mock.Anything).Return(mockData, nil)
	result, err := serv.GetByStatus(status, 1, 2)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func Test_TugasGetByStatusFailed(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	serv := NewTugasService(mockRepo)
	status := true
	mockRepo.On("GetByStatus", status, mock.Anything, mock.Anything).Return(nil, errors.New("Failed to mock data"))
	result, _ := serv.GetByStatus(status, 1, 2)
	assert.Nil(t, result)
	// assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func Test_TugasGetByLevelSuccess(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	serv := NewTugasService(mockRepo)
	mockData := []*models.Tugas{
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
	mockRepo.On("GetBylevel", mock.Anything, mock.Anything, mock.Anything).Return(mockData, nil)
	result, err := serv.GetBylevel("", 1, 2)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func Test_TugasGetByLevelFailed(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	serv := NewTugasService(mockRepo)
	mockRepo.On("GetBylevel", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("Failed to mock data"))
	result, _ := serv.GetBylevel("", 1, 2)
	assert.Nil(t, result)
	// assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func Test_TugasGetByDeadlineSuccess(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	serv := NewTugasService(mockRepo)
	mockData := []*models.Tugas{
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
	mockRepo.On("GetByDeadline", mock.Anything, mock.Anything, mock.Anything).Return(mockData, nil)
	result, err := serv.GetByDeadline("", 1, 2)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func Test_TugasGetByDeadlineFailed(t *testing.T) {
	mockRepo := new(mocksRepo.MockRespository)
	serv := NewTugasService(mockRepo)
	mockRepo.On("GetByDeadline", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("Failed to mock data"))
	result, _ := serv.GetByDeadline("", 1, 2)
	assert.Nil(t, result)
	// assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
