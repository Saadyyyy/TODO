package service

import (
	"Todo/api/repository"
	respons "Todo/api/service/Respons"
	"Todo/models"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TugasService interface {
	GetAll(ctx *gin.Context, page int, perPage int) []interface{}
	GetById(ctx *gin.Context) (interface{}, error)
	Create(ctx *gin.Context) (interface{}, error)
	Update(ctx *gin.Context) (interface{}, error)
	Delete(ctx *gin.Context) (interface{}, error)
	GetByStatus(ctx *gin.Context, bol bool, page int, perPage int) (interface{}, error)
	GetBylevel(ctx *gin.Context, level string, page int, perPage int) (interface{}, error)
	GetByDeadline(ctx *gin.Context, level string, page int, perPage int) (interface{}, error)
}

type TugasServiceImpl struct {
	repo repository.TugasRepository
}

func NewTugasService(repo repository.TugasRepository) TugasService {
	return &TugasServiceImpl{repo: repo}
}

//get all tugas
func (us *TugasServiceImpl) GetAll(ctx *gin.Context, page int, perPage int) []interface{} {
	result := us.repo.GetAll(ctx, page, perPage)
	if result == nil {
		return nil
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

	return []interface{}{respon}
}

//get tugas by id
func (us *TugasServiceImpl) GetById(ctx *gin.Context) (interface{}, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return nil, err
	}

	result, err := us.repo.GetById(uint(id))
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

	return respons, nil
}

// create tugas
func (us *TugasServiceImpl) Create(ctx *gin.Context) (interface{}, error) {
	input := models.Tugas{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validator := validator.New()
	if err := validator.Struct(input); err != nil {
		return nil, err
	}

	result, err := us.repo.Created(input)
	if err != nil {
		return nil, err
	}
	respon := respons.CreateTugasRespon{
		ID:          result.ID,
		Task:        result.Task,
		Level:       result.Level,
		Deadline:    result.Deadline,
		Description: result.Description,
		Status:      result.Status,
		Created_at:  result.CreatedAt,
	}

	return respon, nil
}

//Update tugas
func (us *TugasServiceImpl) Update(ctx *gin.Context) (interface{}, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return gin.H{"message": "ID not found"}, nil
	}

	GetId, err := us.repo.GetById(uint(id))
	if err != nil {
		return nil, err
	}

	if err := ctx.ShouldBindJSON(&GetId); err != nil {
		return nil, err
	}

	validator := validator.New()
	if err := validator.Struct(GetId); err != nil {
		return nil, err
	}

	result, err := us.repo.Update(GetId)
	if err != nil {
		return nil, err
	}

	respon := respons.UpdateTugasRespon{
		ID:          uint(id),
		Task:        result.Task,
		Level:       result.Level,
		Deadline:    result.Deadline,
		Description: result.Description,
		Status:      result.Status,
		Update_at:   time.Now(),
	}

	return respon, nil
}

// delete tugas
func (us *TugasServiceImpl) Delete(ctx *gin.Context) (interface{}, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return nil, err
	}

	GetId, err := us.repo.GetById(uint(id))
	if err != nil {
		return nil, err
	}

	result, err := us.repo.Delete(GetId)
	if err != nil {
		return nil, err
	}

	respon := respons.DeleteTugasRespon{
		ID:         result.ID,
		Task:       result.Task,
		Deleted_at: result.DeletedAt.Time,
	}

	return respon, nil
}

// Get tugas by status

func (us *TugasServiceImpl) GetByStatus(ctx *gin.Context, bol bool, page int, perPage int) (interface{}, error) {
	status, err := strconv.ParseBool(ctx.Param("status"))
	if err != nil {
		return gin.H{"message": "ID not found"}, nil
	}
	result, err := us.repo.GetByStatus(status, page, perPage)
	if err != nil {
		return gin.H{"message": "ID not found"}, nil

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

	return []interface{}{respon}, nil
}

// Logic get all level

func (us *TugasServiceImpl) GetBylevel(ctx *gin.Context, level string, page int, perPage int) (interface{}, error) {
	result, err := us.repo.GetBylevel(level, page, perPage)

	if err != nil {
		return gin.H{"message": "level not found"}, nil
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

	return []interface{}{respon}, nil
}

func (us *TugasServiceImpl) GetByDeadline(ctx *gin.Context, deadline string, page int, perPage int) (interface{}, error) {
	result, err := us.repo.GetByDeadline(deadline, page, perPage)

	if err != nil {
		// Log the error for debugging purposes
		fmt.Println("Error in GetByDeadline:", err)
		return gin.H{"message": "error retrieving tasks"}, err
	}

	respon := make([]respons.GetIdTugasRespon, 0, len(result))
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
	return []interface{}{respon}, nil

}
