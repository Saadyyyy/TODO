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
func (m *MockRespository) Delete(*models.Tugas) (*models.Tugas, error) {
	panic("unimplemented")
}

// GetByDeadline implements repository.TugasRepository.
func (m *MockRespository) GetByDeadline(ded string, page int, perPage int) ([]*models.Tugas, error) {
	panic("unimplemented")
}

// GetById implements repository.TugasRepository.
func (m *MockRespository) GetById(id uint) (*models.Tugas, error) {
	panic("unimplemented")
}

// GetByStatus implements repository.TugasRepository.
func (m *MockRespository) GetByStatus(sts bool, page int, perPage int) ([]*models.Tugas, error) {
	panic("unimplemented")
}

// GetBylevel implements repository.TugasRepository.
func (m *MockRespository) GetBylevel(lvl string, page int, perPage int) ([]*models.Tugas, error) {
	panic("unimplemented")
}

// Update implements repository.TugasRepository.
func (m *MockRespository) Update(*models.Tugas) (*models.Tugas, error) {
	panic("unimplemented")
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
