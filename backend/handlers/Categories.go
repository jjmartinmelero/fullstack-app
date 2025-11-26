package handlers

import (
	"backend-fullstack-app/database"
	"backend-fullstack-app/dto"
	"backend-fullstack-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func GetCategories(c *gin.Context) {
	data := models.Categories{}

	database.DB.Order("id desc").Find(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Categories retrieved successfully",
		"data":    data,
	})

}

func GetCategoryById(c *gin.Context) {
	id := c.Param("id")

	data := models.Category{}

	if err := database.DB.First(&data, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "ERROR",
			"message": "Category not found",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Category retrieved successfully",
		"data":    data,
	})
}

func CreateCategory(c *gin.Context) {
	var body dto.CategoryDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "Invalid request body",
			"err":     err.Error(),
		})
		return
	}

	// check if category already exists
	if err := database.DB.Where("name = ?", body.Name).First(&models.Category{}).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":  "ERROR",
			"message": "Category already exists",
		})
		return
	}

	data := models.Category{
		Name: body.Name,
		Slug: slug.Make(body.Name),
	}

	database.DB.Save(&data)

	c.JSON(http.StatusCreated, gin.H{
		"status":  "OK",
		"message": "Category created successfully",
		"data":    data,
	})

}

func EditCategory(c *gin.Context) {
	id := c.Param("id")

	var body dto.CategoryDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "Invalid request body",
			"err":     err.Error(),
		})
		return
	}

	data := models.Category{}

	if err := database.DB.First(&data, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "ERROR",
			"message": "Category not found",
			"err":     err.Error(),
		})
		return
	}

	data.Name = body.Name
	data.Slug = slug.Make(body.Name)

	database.DB.Save(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Category updated successfully",
		"data":    data,
	})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	data := models.Category{}

	if err := database.DB.First(&data, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "ERROR",
			"message": "Category not found",
			"err":     err.Error(),
		})
		return
	}

	database.DB.Delete(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Category deleted successfully",
		"data":    data,
	})
}
