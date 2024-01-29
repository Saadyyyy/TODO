package main

import (
	"Todo/apps/config"
	"Todo/apps/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	db := config.InitializeDatabase()

	routes.Api(r, db)
	// routes.ImagesRouter(r, db)
	r.Run()
}
