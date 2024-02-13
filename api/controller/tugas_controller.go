package controller

import (
	"Todo/api/service"
	respons "Todo/api/service/Respons"
	"Todo/models"
	"Todo/utils/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TugasController struct {
	TugasService service.TugasService
}

func NewTugasController(tugasService service.TugasService, ctx *gin.Context) *TugasController {
	return &TugasController{
		TugasService: tugasService,
	}
}

func (uc *TugasController) GetAll(ctx *gin.Context) {
	page, perPage := helper.Pagination(ctx)

	data := uc.TugasService.GetAll(ctx, page, perPage)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    data,
	})
}

func (uc *TugasController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Error",
			"Data":    err,
		})
	}
	data, err := uc.TugasService.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Error",
			"Data":    err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Succes",
		"data":    data,
	})
}

func (uc *TugasController) Create(ctx *gin.Context) {
	tugas := &models.Tugas{}
	ctx.ShouldBindJSON(&tugas)
	data, err := uc.TugasService.Create(tugas)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Status Internal Server Error",
			"error":   err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Status ok",
		"data":    data,
	})
}

func (uc *TugasController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
			"error":   err.Error(),
		})
		return
	}
	tugas := respons.UpdateTugasRespon{}
	if err := ctx.ShouldBindJSON(&tugas); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	data, err := uc.TugasService.Update(id, tugas)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Status Internal Server Error",
			"error":   err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Status ok",
		"data":    data,
	})
}
func (uc *TugasController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
			"error":   err.Error(),
		})
		return
	}
	data, err := uc.TugasService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Status Internal Server Error",
			"Data":    err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Status ok",
		"data":    data,
	})
}

func (uc *TugasController) GetByStatus(ctx *gin.Context) {
	page, perPage := helper.Pagination(ctx)
	// Mengambil nilai status dari URL
	statusParam := ctx.Param("status")

	// Mengkonversi statusParam ke boolean
	status, err := strconv.ParseBool(statusParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status parameter"})
		return
	}

	// Memanggil metode GetByStatus dari service
	result, err := uc.TugasService.GetByStatus(status, page, perPage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengirimkan respons ke client
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Status ok",
		"data":    result,
	})
}

// controller get tugas deaGetByDeadline
func (uc *TugasController) GetBylevel(ctx *gin.Context) {
	page, perPage := helper.Pagination(ctx)
	// Mengambil nilai status dari URL
	deaGetBylevel := ctx.Param("level")

	// Memanggil metode GetBylevel dari service
	result, err := uc.TugasService.GetBylevel(deaGetBylevel, page, perPage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengirimkan respons ke client
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Status ok",
		"data":    result,
	})
}

// controller get tugas deaGetByDeadline
func (uc *TugasController) GetByDeadline(ctx *gin.Context) {
	page, perPage := helper.Pagination(ctx)
	// Mengambil nilai status dari URL
	deaGetBylevel := ctx.Param("deadline")

	// Memanggil metode GetBylevel dari service
	result, err := uc.TugasService.GetByDeadline(deaGetBylevel, page, perPage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengirimkan respons ke client
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Status ok",
		"data":    result,
	})
}
