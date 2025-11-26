package handlers

import (
	"backend-fullstack-app/database"
	"backend-fullstack-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
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
