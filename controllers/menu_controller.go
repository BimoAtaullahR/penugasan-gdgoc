package controllers

import (
	"net/http"

	"github.com/BimoAtaullahR/penugasan-gdgoc/config"
	"github.com/BimoAtaullahR/penugasan-gdgoc/models"
	"github.com/gin-gonic/gin"
)

func CreateMenu(c *gin.Context){
	var menu models.Menu

	//Binding JSON body ke struct
	if err := c.ShouldBindBodyWithJSON(&menu); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//simpan ke database
	if err := config.DB.Create(&menu).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Success", "data": menu})
}

func ListMenu(c *gin.Context){
	var menu models.Menu

	result := config.DB.Find(&menu)
	c.JSON(http.StatusOK, gin.H{"data": result})
}