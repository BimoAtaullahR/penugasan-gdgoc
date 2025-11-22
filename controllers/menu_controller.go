package controllers

import (
	"net/http"

	"github.com/BimoAtaullahR/penugasan-gdgoc/config"
	"github.com/BimoAtaullahR/penugasan-gdgoc/models"
	"github.com/gin-gonic/gin"
)

func CreateMenu(c *gin.Context) {
	var menu models.Menu

	//Binding JSON body ke struct
	if err := c.ShouldBindBodyWithJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//simpan ke database
	if err := config.DB.Create(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Success", "data": menu})
}

func ListMenu(c *gin.Context) {
	var menus []models.Menu

	//inisialisasi query dasar
	query := config.DB.Model(&models.Menu{})

	//FILTERING (WHERE)
	//filter search (q) (mencari berdasarkan nama atau deskripsi, SQL LIKE)
	if q := c.Query("q"); q != "" {
		likeQuery := "%" + q + "%"
		query = query.Where("name LIKE ? OR description LIKE ?", likeQuery, likeQuery)
	}

	//filter kategori
	if category := c.Query("category"); category != "" {
		query = query.Where("category = ?", category)
	}

	//filter minimum price
	if minPrice := c.Query("min_price"); minPrice != "" {
		query = query.Where("price >= ?", minPrice)
	}

	//filter maximum price
	if maxPrice := c.Query("max_price"); maxPrice != "" {
		query = query.Where("price <= ?", maxPrice)
	}

	//filter max calories
	if maxCalories := c.Query("max_cal"); maxCalories != "" {
		query = query.Where("calories <= ?", maxCalories)
	}

	//PAGINATION (LIMIT DAN OFFSET)
	var jumlah int64
	query.Count(&jumlah)

	type QueryParams struct {
		PerPage int    `form:"per_page, default=10"`
		Page    int    `form:"page, default=1"`
		Sort    string `form:"sort"`
	}

	var params QueryParams
	page := 1
	perPage := 10
	if err := c.ShouldBindQuery(&params); err == nil {
		page = params.Page
		perPage = params.PerPage
	}

	offset := (page - 1) * perPage

	query = query.Limit(perPage).Offset(offset)

	//menerapkan sorting
	if params.Sort == "price:asc" {
		query = query.Order("price asc")
	} else if params.Sort == "price:dsc" {
		query = query.Order("price dsc")
	} else {
		query = query.Order("created_at dsc")
	}

	//EKSEKUSI QUERY
	if err := query.Find(&menus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//FORMAT RESPONSE
	c.JSON(http.StatusOK, gin.H{
		"data": menus,
		"pagination": gin.H{
			"total": jumlah,
			"page": page,
			"per_page" : perPage,
			"total_pages": (jumlah + int64(perPage) - 1)/int64(perPage),
		},
	})
}
