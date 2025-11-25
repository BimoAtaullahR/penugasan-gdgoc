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
	r.GET("/menu/:id", controllers.GetMenuByID)
	r.PUT("/menu/:id", controllers.UpdateMenuByID)
	r.DELETE("/menu/:id", controllers.DeleteMenuByID)
	r.GET("/menu/group-by-category", controllers.GroupByCategory)
	r.GET("/menu/search", controllers.SearchByText)
	
	return r
}