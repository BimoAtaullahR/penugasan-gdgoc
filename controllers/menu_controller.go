package controllers

import (
	"net/http"
	"strconv"

	"github.com/BimoAtaullahR/penugasan-gdgoc/config"
	"github.com/BimoAtaullahR/penugasan-gdgoc/models"
	"github.com/BimoAtaullahR/penugasan-gdgoc/services"
	"github.com/gin-gonic/gin"
)

func CreateMenu(c *gin.Context) {
	var menu models.Menu

	//Binding JSON body ke struct
	if err := c.ShouldBindBodyWithJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if menu.Description == ""{
		desc, err := services.GenerateDescription(c, menu.Name, menu.Ingredients)
		if err==nil{
			menu.Description = desc
		}
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
		if params.Page > 0 {
			page = params.Page
		}
		if params.PerPage > 0 {
			perPage = params.PerPage
		}
	}

	offset := (page - 1) * perPage

	query = query.Limit(perPage).Offset(offset)

	//menerapkan sorting
	if params.Sort == "price:asc" {
		query = query.Order("price asc")
	} else if params.Sort == "price:desc" {
		query = query.Order("price desc")
	} else {
		query = query.Order("created_at desc")
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
			"total":       jumlah,
			"page":        page,
			"per_page":    perPage,
			"total_pages": (jumlah + int64(perPage) - 1) / int64(perPage),
		},
	})
}

func GetMenuByID(c *gin.Context) {
	id := c.Param("id")
	var menu models.Menu

	if err := config.DB.First(&menu, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": menu})
}

func UpdateMenuByID(c *gin.Context) {
	id := c.Param("id")
	var menu models.Menu
	if err := config.DB.First(&menu, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var inputData models.Menu
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&menu).Updates(inputData)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    menu,
	})
}

func DeleteMenuByID(c *gin.Context) {
	id := c.Param("id")
	var menu models.Menu
	if err := config.DB.First(&menu, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Delete(&menu)
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func GroupByCategory(c *gin.Context) {
	var menus models.Menu
	mode := c.Query("mode")

	type CategoryCount struct {
		Category string
		Total    int64
	}
	var categoryMenu []CategoryCount

	switch mode {
	case "count":
		result := make(map[string]int64)
		config.DB.Model(&menus).Select("category, count(*) as total").Group("category").Scan(&categoryMenu)
		for _, item := range categoryMenu {
			result[item.Category] = item.Total
		}
		c.JSON(http.StatusOK, gin.H{"data": result})

	case "list":
		var listCategories []struct{
			Category string
		}
		type MenuResponseSimple struct{
			ID uint `json:"id"`
			Name string `json:"name"`
			Category string `json:"category"`
			Price float64 `json:"price"`
		}
		config.DB.Model(&menus).Distinct("category").Group("category").Scan(&listCategories)

		results := make(map[string][]MenuResponseSimple)
		perPage := c.DefaultQuery("per_category", "5")
		limitPerPage, err := strconv.Atoi(perPage)
		if err != nil{limitPerPage = 5}
		
		for _, item := range listCategories{
			var data []MenuResponseSimple
			config.DB.Model(&models.Menu{}).Where("category = ?", item.Category).Limit(limitPerPage).Find(&data)
			results[item.Category] = data
		}
		
		c.JSON(http.StatusOK, gin.H{"data": results})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mode tidak valid. Gunakan ?mode=count atau ?mode=list"})
	}
}


func SearchByText(c * gin.Context){
	query := config.DB.Model(&models.Menu{})

	if q := c.Query("q"); q != ""{
		query = query.Where("name LIKE ? OR category LIKE ?", "%"+q+"%", "%"+q+"%")
	}
	var jumlah int64
	query.Count(&jumlah)

	type QueryParams struct{
		Page int `form:"page, default=1"`
		PerPage int `form:"per_page, default=5"`
	}
	var params QueryParams
	page := 1
	perPage := 5
	if err := c.ShouldBindQuery(&params); err == nil{
		if params.Page > 0 {page = params.Page}
		if params.PerPage > 0 {perPage = params.PerPage}
	}
	offset := (page - 1)*perPage
	query = query.Limit(perPage).Offset(offset)

	type Response struct{
		ID uint `gorm:"primaryKey" json:"id"`
		Name string `json:"name" binding:"required"`
		Category string `json:"category" binding:"required"`
		Price float64 `json:"price" binding:"required"`
		Calories float64 `json:"calories"`
		Ingredients []string `json:"ingredients" gorm:"type:text;serializer:json"`
		Description string `json:"description"`
	}
	var results []Response
	query.Find(&results)

	c.JSON(http.StatusOK, gin.H{
		"data" : results,
		"pagination": gin.H{
			"total": jumlah,
			"page": page,
			"per_page": perPage,
		},
	})
}
