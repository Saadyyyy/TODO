package controller

import (
	"Todo/feature/tugas/service"
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
	_, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Error",
			"Data":    err,
		})
	}
	data, err := uc.TugasService.GetById(ctx)
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
	data, err := uc.TugasService.Create(ctx)
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

	data, err := uc.TugasService.Update(ctx)
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
	data, err := uc.TugasService.Delete(ctx)
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
	// Mengambil nilai status dari URL
	statusParam := ctx.Param("status")

	// Mengkonversi statusParam ke boolean
	status, err := strconv.ParseBool(statusParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status parameter"})
		return
	}

	// Memanggil metode GetByStatus dari service
	result, err := uc.TugasService.GetByStatus(ctx, status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengirimkan respons ke client
	ctx.JSON(http.StatusOK, result)
}

// controller get tugas deaGetByDeadline
func (uc *TugasController) GetBylevel(ctx *gin.Context) {
	// Mengambil nilai status dari URL
	deaGetBylevel := ctx.Param("deaGetBylevel")

	// Memanggil metode GetBylevel dari service
	result, err := uc.TugasService.GetBylevel(ctx, deaGetBylevel)
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
	// Mengambil nilai status dari URL
	deaGetByDeadline := ctx.Param("deadline")

	// Memanggil metode GetByDeadline dari service
	result, err := uc.TugasService.GetByDeadline(ctx, deaGetByDeadline)
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
