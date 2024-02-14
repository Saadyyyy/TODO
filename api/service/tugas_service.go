package service

import (
	"Todo/api/repository"
	respons "Todo/api/service/Respons"
	"Todo/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TugasService interface {
	GetAll(ctx *gin.Context, page int, perPage int) []respons.GetIdTugasRespon
	GetById(id uint) (*respons.GetIdTugasRespon, error)
	Create(*models.Tugas) (*respons.CreateTugasRespon, error)
	Update(id uint, up respons.UpdateTugasRespon) (*respons.UpdateTugasRespon, error)
	Delete(id uint) (*respons.DeleteTugasRespon, error)
	GetByStatus(bol bool, page int, perPage int) ([]respons.GetIdTugasRespon, error)
	GetBylevel(level string, page int, perPage int) ([]respons.GetIdTugasRespon, error)
	GetByDeadline(level string, page int, perPage int) ([]respons.GetIdTugasRespon, error)
}

type TugasServiceImpl struct {
	repo repository.TugasRepository
}

func NewTugasService(repo repository.TugasRepository) TugasService {
	return &TugasServiceImpl{repo: repo}
}

// get all tugas
func (us *TugasServiceImpl) GetAll(ctx *gin.Context, page int, perPage int) []respons.GetIdTugasRespon {
	result := us.repo.GetAll(ctx, page, perPage)

	respon := []respons.GetIdTugasRespon{}
	for _, tugas := range result {
		respons := respons.GetIdTugasRespon{
			ID:          tugas.ID,
			Task:        tugas.Task,
			Level:       tugas.Level,
			Deadline:    tugas.Deadline,
			Description: tugas.Description,
			Status:      tugas.Status,
		}
		respon = append(respon, respons)
	}

	return respon
}

// get tugas by id
func (us *TugasServiceImpl) GetById(id uint) (*respons.GetIdTugasRespon, error) {
	result, err := us.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	respons := respons.GetIdTugasRespon{
		ID:          result.ID,
		Task:        result.Task,
		Level:       result.Level,
		Deadline:    result.Deadline,
		Description: result.Description,
		Status:      result.Status,
	}

	return &respons, nil
}

// create tugas
func (us *TugasServiceImpl) Create(input *models.Tugas) (*respons.CreateTugasRespon, error) {
	validator := validator.New()
	validator.Struct(input)

	result, err := us.repo.Created(input)
	if err != nil {
		return nil, err
	}
	data := respons.CreateTugasRespon{
		ID:          result.ID,
		Task:        result.Task,
		Level:       result.Level,
		Deadline:    result.Deadline,
		Description: result.Description,
		Status:      result.Status,
		Created_at:  result.CreatedAt,
	}

	return &data, nil
}

// Update tugas
func (us *TugasServiceImpl) Update(ids uint, up respons.UpdateTugasRespon) (*respons.UpdateTugasRespon, error) {
	getId, err := us.repo.GetById(ids)
	if err != nil {
		return nil, err
	}

	// Update data dengan nilai yang diperbarui
	getId.Status = up.Status

	// Panggil metode Update pada repositori
	result, err := us.repo.Update(getId)
	if err != nil {
		return nil, err
	}
	respon := respons.UpdateTugasRespon{
		ID:          result.ID,
		Task:        result.Task,
		Level:       result.Level,
		Deadline:    result.Deadline,
		Description: result.Description,
		Status:      result.Status,
		Update_at:   time.Now(),
	}

	return &respon, nil
}

// delete tugas
func (us *TugasServiceImpl) Delete(ids uint) (*respons.DeleteTugasRespon, error) {
	getId, err := us.repo.GetById(ids)
	if err != nil {
		return nil, err
	}
	result, err := us.repo.Delete(getId)
	if err != nil {
		return nil, err
	}

	respon := respons.DeleteTugasRespon{
		ID:         result.ID,
		Task:       result.Task,
		Deleted_at: result.DeletedAt.Time,
	}

	return &respon, nil
}

// Get tugas by status

func (us *TugasServiceImpl) GetByStatus(bol bool, page int, perPage int) ([]respons.GetIdTugasRespon, error) {

	result, err := us.repo.GetByStatus(bol, page, perPage)
	if err != nil {
		return nil, nil

	}
	respon := []respons.GetIdTugasRespon{}
	for _, tugas := range result {
		respons := respons.GetIdTugasRespon{
			ID:          tugas.ID,
			Task:        tugas.Task,
			Level:       tugas.Level,
			Deadline:    tugas.Deadline,
			Description: tugas.Description,
			Status:      tugas.Status,
		}
		respon = append(respon, respons)
	}

	return respon, nil
}

// Logic get all level

func (us *TugasServiceImpl) GetBylevel(level string, page int, perPage int) ([]respons.GetIdTugasRespon, error) {
	result, err := us.repo.GetBylevel(level, page, perPage)

	if err != nil {
		return nil, err
	}
	respon := []respons.GetIdTugasRespon{}
	for _, tugas := range result {
		respons := respons.GetIdTugasRespon{
			ID:          tugas.ID,
			Task:        tugas.Task,
			Level:       tugas.Level,
			Deadline:    tugas.Deadline,
			Description: tugas.Description,
			Status:      tugas.Status,
		}
		respon = append(respon, respons)
	}

	return respon, nil
}

func (us *TugasServiceImpl) GetByDeadline(deadline string, page int, perPage int) ([]respons.GetIdTugasRespon, error) {
	result, err := us.repo.GetByDeadline(deadline, page, perPage)

	if err != nil {
		return nil, err
	}

	respon := []respons.GetIdTugasRespon{}
	for _, tugas := range result {
		response := respons.GetIdTugasRespon{
			ID:          tugas.ID,
			Task:        tugas.Task,
			Level:       tugas.Level,
			Deadline:    tugas.Deadline,
			Description: tugas.Description,
			Status:      tugas.Status,
		}
		respon = append(respon, response)
	}
	return respon, nil

}
