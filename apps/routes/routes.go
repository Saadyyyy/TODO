package routes

import (
	"Todo/feature/tugas/controller"
	"Todo/feature/tugas/repository"
	"Todo/feature/tugas/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ctx *gin.Context
)

func Api(r *gin.Engine, db *gorm.DB) {
	repository := repository.NewTugasReporitory(db)
	service := service.NewTugasService(repository)
	controller := controller.NewTugasController(service, ctx)

	Tugas := r.Group("tugas")
	{
		Tugas.GET("/", controller.GetAll)
		Tugas.GET("/:id", controller.GetByID)
		Tugas.POST("/create", controller.Create)
		Tugas.PATCH("/:id", controller.Update)
		Tugas.DELETE("/:id", controller.Delete)
		Tugas.GET("/status/:status", controller.GetByStatus)
		Tugas.GET("/level/:level", controller.GetBylevel)
	}
}
