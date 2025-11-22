package routes

import (
	"github.com/BimoAtaullahR/penugasan-gdgoc/config"
	"github.com/BimoAtaullahR/penugasan-gdgoc/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine{
	config.ConnectDatabase()
	r := gin.Default()
	r.POST("/menu", controllers.CreateMenu)
	r.GET("/menu", controllers.ListMenu)

	return r
}