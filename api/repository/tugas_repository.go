package repository

import (
	"Todo/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TugasRepository interface {
	GetAll(ctx *gin.Context, page int, perPage int) []models.Tugas
	GetById(id uint) (*models.Tugas, error)
	Created(*models.Tugas) (*models.Tugas, error)
	Update(*models.Tugas) (*models.Tugas, error)
	Delete(*models.Tugas) (*models.Tugas, error)
	GetByStatus(sts bool, page int, perPage int) ([]*models.Tugas, error)
	GetBylevel(lvl string, page int, perPage int) ([]*models.Tugas, error)
	GetByDeadline(ded string, page int, perPage int) ([]*models.Tugas, error)
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

func (ur *TugasRepositoryImp) GetById(id uint) (*models.Tugas, error) {
	user := models.Tugas{}
	ur.db.First(&user, id)

	return &user, nil
}

func (ur *TugasRepositoryImp) Created(user *models.Tugas) (*models.Tugas, error) {
	if err := ur.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *TugasRepositoryImp) Update(tugas *models.Tugas) (*models.Tugas, error) {
	if err := ur.db.Model(&models.Tugas{}).Where("id = ?", tugas.ID).Update("status", tugas.Status).Error; err != nil {
		return nil, err
	}
	return tugas, nil
}

// Delete user
func (ur *TugasRepositoryImp) Delete(user *models.Tugas) (*models.Tugas, error) {
	if err := ur.db.Model(&user).Where("id = ?", user.ID).Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetByStatus implements TugasRepository.
func (ur *TugasRepositoryImp) GetByStatus(status bool, page int, perPage int) ([]*models.Tugas, error) {
	tugas := []*models.Tugas{}
	offsets := (page - 1) * perPage
	err := ur.db.Model(&tugas).Where("status = ?", status).Limit(perPage).Offset(offsets).Find(&tugas).Error
	if err != nil {
		return nil, err
	}
	return tugas, nil
}

// Get tugas by level
func (ur *TugasRepositoryImp) GetBylevel(level string, page int, perPage int) ([]*models.Tugas, error) {
	tugas := []*models.Tugas{}
	offsets := (page - 1) * perPage
	err := ur.db.Model(&tugas).Where("level LIKE ?", "%"+level+"%").Limit(perPage).Offset(offsets).Find(&tugas).Error
	if err != nil {
		return nil, err
	}
	return tugas, nil
}

func (ur *TugasRepositoryImp) GetByDeadline(deadline string, page int, perPage int) ([]*models.Tugas, error) {
	tugas := []*models.Tugas{}
	offsets := (page - 1) * perPage
	err := ur.db.Model(&tugas).Where("deadline LIKE ?", "%"+deadline+"%").Limit(perPage).Offset(offsets).Find(&tugas).Error
	if err != nil {
		return nil, err
	}
	return tugas, nil
}
