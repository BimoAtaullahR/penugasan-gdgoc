package config

import (
	"os"

	"github.com/BimoAtaullahR/penugasan-gdgoc/models"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Menu{})
	DB = database
}