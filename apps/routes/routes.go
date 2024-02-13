package routes

import (
	"Todo/api/controller"
	"Todo/api/repository"
	"Todo/api/service"

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
		Tugas.PUT("/:id", controller.Update)
		Tugas.DELETE("/:id", controller.Delete)
		Tugas.GET("/status/:status", controller.GetByStatus)
		Tugas.GET("/level/:level", controller.GetBylevel)
		Tugas.GET("/deadline/:deadline", controller.GetByDeadline)
	}
}
