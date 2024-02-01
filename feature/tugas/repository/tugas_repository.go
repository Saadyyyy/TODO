package repository

import (
	"Todo/models"
	"errors"

	"gorm.io/gorm"
)

type TugasRepository interface {
	GetAll() []models.Tugas
	GetById(id uint) (models.Tugas, error)
	Created(models.Tugas) (*models.Tugas, error)
	Update(models.Tugas) (*models.Tugas, error)
	Delete(models.Tugas) (*models.Tugas, error)
	GetByStatus(bool) ([]models.Tugas, error)
	GetBylevel(string) ([]models.Tugas, error)
	GetByDeadline(string) ([]models.Tugas, error)
}

type TugasRepositoryImp struct {
	db *gorm.DB
}

func NewTugasReporitory(db *gorm.DB) TugasRepository {
	return &TugasRepositoryImp{db: db}
}

func (ur *TugasRepositoryImp) GetAll() []models.Tugas {
	user := []models.Tugas{}
	ur.db.Find(&user)

	return user
}

func (ur *TugasRepositoryImp) GetById(id uint) (models.Tugas, error) {
	user := models.Tugas{}
	ur.db.First(&user, id)

	return user, nil
}

func (ur *TugasRepositoryImp) Created(user models.Tugas) (*models.Tugas, error) {
	err := ur.db.Create(&user)

	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}

func (ur *TugasRepositoryImp) Update(tugas models.Tugas) (*models.Tugas, error) {
	// Pastikan tugas memiliki ID yang tidak kosong
	if tugas.ID == 0 {
		return nil, errors.New("ID tugas tidak valid")
	}

	// Lakukan pembaruan hanya pada field yang ingin diubah
	err := ur.db.Model(&models.Tugas{}).Where("id = ?", tugas.ID).Updates(tugas).Error
	if err != nil {
		return nil, err
	}

	return &tugas, nil
}

// Delete user
func (ur *TugasRepositoryImp) Delete(user models.Tugas) (*models.Tugas, error) {
	err := ur.db.Delete(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetByStatus implements TugasRepository.
func (ur *TugasRepositoryImp) GetByStatus(status bool) ([]models.Tugas, error) {
	user := []models.Tugas{}
	err := ur.db.Model(&user).Where("level = ?", status).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

//Get tugas by level
func (ur *TugasRepositoryImp) GetBylevel(level string) ([]models.Tugas, error) {
	tugas := []models.Tugas{}
	err := ur.db.Model(&tugas).Where("level LIKE ?", "%"+level+"%").Find(&tugas).Error
	if err != nil {
		return nil, err
	}
	return tugas, nil
}

func (ur *TugasRepositoryImp) GetByDeadline(deadline string) ([]models.Tugas, error) {
	tugas := []models.Tugas{}
	err := ur.db.Model(&tugas).Where("deadline LIKE ?", "%"+deadline+"%").Find(&tugas).Error
	if err != nil {
		return nil, err
	}
	return tugas, nil
}
