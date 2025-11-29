package handlers

import (
	"backend-fullstack-app/database"
	"backend-fullstack-app/dto"
	"backend-fullstack-app/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRecipes(c *gin.Context) {
	data := models.Recipes{}

	database.DB.Preload("Category").Find(&data)

	// Debug: print loaded data
	for i, recipe := range data {
		fmt.Printf("Recipe %d: ID=%d, CategoryId=%d, Category.Id=%d, Category.Name=%s\n",
			i, recipe.Id, recipe.CategoryId, recipe.Category.Id, recipe.Category.Name)
	}

	var parseData = dto.ReceipeResponses{}

	for _, element := range data {
		parseData = append(parseData, dto.ReceipeResponse{
			ID:          element.Id,
			Name:        element.Name,
			Slug:        element.Slug,
			CategoryId:  element.CategoryId,
			Category:    element.Category.Name,
			Time:        element.Time,
			Description: element.Description,
			Photo:       element.Photo,
			Date:        element.Date.String(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Recipes retrieved successfully",
		"data":    parseData,
	})
}
