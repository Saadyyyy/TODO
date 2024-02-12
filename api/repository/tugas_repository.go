package repository

import (
	"Todo/models"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TugasRepository interface {
	GetAll(ctx *gin.Context, page int, perPage int) []models.Tugas
	GetById(id uint) (models.Tugas, error)
	Created(models.Tugas) (models.Tugas, error)
	Update(models.Tugas) (models.Tugas, error)
	Delete(models.Tugas) (models.Tugas, error)
	GetByStatus(sts bool, page int, perPage int) ([]models.Tugas, error)
	GetBylevel(lvl string, page int, perPage int) ([]models.Tugas, error)
	GetByDeadline(ded string, page int, perPage int) ([]models.Tugas, error)
}

type TugasRepositoryImp struct {
	db *gorm.DB
}

func NewTugasReporitory(db *gorm.DB) TugasRepository {
	return &TugasRepositoryImp{db: db}
}

func (ur *TugasRepositoryImp) GetAll(ctx *gin.Context, page int, perPage int) []models.Tugas {
	user := []models.Tugas{}

	offsets := (page - 1) * perPage
	ur.db.Limit(perPage).Offset(offsets).Find(&user)

	return user
}

func (ur *TugasRepositoryImp) GetById(id uint) (models.Tugas, error) {
	user := models.Tugas{}
	ur.db.First(&user, id)

	return user, nil
}

func (ur *TugasRepositoryImp) Created(user models.Tugas) (models.Tugas, error) {
	err := ur.db.Create(&user)

	if err.Error != nil {
		return models.Tugas{}, err.Error
	}

	return user, nil
}

func (ur *TugasRepositoryImp) Update(models.Tugas) (models.Tugas, error) {
	tugas := models.Tugas{}
	// Pastikan tugas memiliki ID yang tidak kosong
	if tugas.ID == 0 {
		return models.Tugas{}, errors.New("ID tugas tidak valid")
	}

	// Lakukan pembaruan hanya pada field yang ingin diubah
	err := ur.db.Model(&models.Tugas{}).Where("id = ?", tugas.ID).Updates(tugas).Error
	if err != nil {
		return models.Tugas{}, err
	}

	return tugas, nil
}

// Delete user
func (ur *TugasRepositoryImp) Delete(user models.Tugas) (models.Tugas, error) {
	err := ur.db.Delete(&user).Error

	if err != nil {
		return models.Tugas{}, err
	}

	return user, nil
}

// GetByStatus implements TugasRepository.
func (ur *TugasRepositoryImp) GetByStatus(status bool, page int, perPage int) ([]models.Tugas, error) {
	tugas := []models.Tugas{}
	offsets := (page - 1) * perPage
	err := ur.db.Model(&tugas).Where("status = ?", status).Limit(perPage).Offset(offsets).Find(&tugas).Error
	if err != nil {
		return nil, err
	}
	return tugas, nil
}

// Get tugas by level
func (ur *TugasRepositoryImp) GetBylevel(level string, page int, perPage int) ([]models.Tugas, error) {
	tugas := []models.Tugas{}
	offsets := (page - 1) * perPage
	err := ur.db.Model(&tugas).Where("level LIKE ?", "%"+level+"%").Limit(perPage).Offset(offsets).Find(&tugas).Error
	if err != nil {
		return nil, err
	}
	return tugas, nil
}

func (ur *TugasRepositoryImp) GetByDeadline(deadline string, page int, perPage int) ([]models.Tugas, error) {
	tugas := []models.Tugas{}
	offsets := (page - 1) * perPage
	err := ur.db.Model(&tugas).Where("deadline LIKE ?", "%"+deadline+"%").Limit(perPage).Offset(offsets).Find(&tugas).Error
	if err != nil {
		return nil, err
	}
	return tugas, nil
}
