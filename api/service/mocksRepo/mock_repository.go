package mocksRepo

import (
	"Todo/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockRespository struct {
	mock.Mock
}

// Created implements repository.TugasRepository.
func (m *MockRespository) Created(tugas *models.Tugas) (*models.Tugas, error) {
	arg := m.Called(tugas)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	} else {
		tugasbeda := arg.Get(0).(*models.Tugas)
		return tugasbeda, nil
	}
}

// Delete implements repository.TugasRepository.
func (m *MockRespository) Delete(tugas *models.Tugas) (*models.Tugas, error) {
	arg := m.Called(tugas)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	} else {
		tugasbeda := arg.Get(0).(*models.Tugas)
		return tugasbeda, nil
	}
}

// GetByDeadline implements repository.TugasRepository.
func (m *MockRespository) GetByDeadline(ded string, page int, perPage int) ([]*models.Tugas, error) {
	arg := m.Called(ded, page, perPage)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	} else {
		tugasbeda := arg.Get(0).([]*models.Tugas)
		return tugasbeda, nil
	}
}

// GetById implements repository.TugasRepository.
func (m *MockRespository) GetById(id uint) (*models.Tugas, error) {
	arg := m.Called(id)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	} else {
		tugasbeda := arg.Get(0).(*models.Tugas)
		return tugasbeda, nil
	}
}

// GetByStatus implements repository.TugasRepository.
func (m *MockRespository) GetByStatus(sts bool, page int, perPage int) ([]*models.Tugas, error) {
	arg := m.Called(sts, page, perPage)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	} else {
		tugasbeda := arg.Get(0).([]*models.Tugas)
		return tugasbeda, nil
	}
}

// GetBylevel implements repository.TugasRepository.
func (m *MockRespository) GetBylevel(lvl string, page int, perPage int) ([]*models.Tugas, error) {
	arg := m.Called(lvl, page, perPage)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	} else {
		tugasbeda := arg.Get(0).([]*models.Tugas)
		return tugasbeda, nil
	}
}

// Update implements repository.TugasRepository.
func (m *MockRespository) Update(tugas *models.Tugas) (*models.Tugas, error) {
	arg := m.Called(tugas)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	} else {
		tugasbeda := arg.Get(0).(*models.Tugas)
		return tugasbeda, nil
	}
}

func (m *MockRespository) GetAll(ctx *gin.Context, page int, perPage int) []models.Tugas {
	arg := m.Called(ctx, page, perPage)
	if arg.Get(0) == nil {
		return nil
	} else {
		tugas := arg.Get(0).([]models.Tugas)
		return tugas
	}
}
