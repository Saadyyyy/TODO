package controller

import (
	"Todo/feature/tugas/service"
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
	data := uc.TugasService.GetAll()

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Succes",
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

func (uc *TugasController) GetByStatus(c *gin.Context) {
	// Mengambil nilai status dari URL
	statusParam := c.Param("status")

	// Mengkonversi statusParam ke boolean
	status, err := strconv.ParseBool(statusParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status parameter"})
		return
	}

	// Memanggil metode GetByStatus dari service
	result, err := uc.TugasService.GetByStatus(c, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengirimkan respons ke client
	c.JSON(http.StatusOK, result)
}
