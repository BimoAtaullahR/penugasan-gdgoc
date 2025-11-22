package config

import (
	"github.com/BimoAtaullahR/penugasan-gdgoc/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	database, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil{
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Menu{})
	DB = database
}