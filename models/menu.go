package models

import(
	"time"
	"gorm.io/gorm"
)

type Menu struct{
	ID uint `gorm:"primaryKey" json:"id"`
	Name string `json:"name" binding:"required"`
	Category string `json:"category" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	Calories float64 `json:"calories"`

	Ingredients []string `json:"ingredients" gorm:"type:text;serializer:json"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}