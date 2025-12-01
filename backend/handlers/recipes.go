package handlers

import (
	"backend-fullstack-app/database"
	"backend-fullstack-app/dto"
	"backend-fullstack-app/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
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

	scheme := "http"

	if c.Request.TLS != nil {
		scheme = "https"
	}

	for _, element := range data {

		date := fmt.Sprintf("%d-%02d-%02d", element.Date.Day(), element.Date.Month(), element.Date.Year())

		photo := fmt.Sprintf("%s://%s/public/uploads/recipes/%s", scheme, c.Request.Host, element.Photo)

		parseData = append(parseData, dto.ReceipeResponse{
			ID:          element.Id,
			Name:        element.Name,
			Slug:        element.Slug,
			CategoryId:  element.CategoryId,
			Category:    element.Category.Name,
			Time:        element.Time,
			Description: element.Description,
			Photo:       photo,
			Date:        date,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Recipes retrieved successfully",
		"data":    parseData,
	})
}

func GetRecipeById(c *gin.Context) {
	id := c.Param("id")

	data := models.Recipe{}

	if err := database.DB.Preload("Category").First(&data, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "ERROR",
			"message": "Recipe not found",
			"err":     err.Error(),
		})
		return
	}

	scheme := "http"

	if c.Request.TLS != nil {
		scheme = "https"
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Recipe retrieved successfully",
		"data": dto.ReceipeResponse{
			ID:          data.Id,
			Name:        data.Name,
			Slug:        data.Slug,
			CategoryId:  data.CategoryId,
			Category:    data.Category.Name,
			Time:        data.Time,
			Description: data.Description,
			Photo:       fmt.Sprintf("%s://%s/public/uploads/recipes/%s", scheme, c.Request.Host, data.Photo),
			Date:        fmt.Sprintf("%d-%02d-%02d", data.Date.Day(), data.Date.Month(), data.Date.Year()),
		},
	})
}

func CreateRecipe(c *gin.Context) {
	fileForm, err := c.FormFile("photo")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "Failed to upload photo",
			"err":     err.Error(),
		})
		return
	}

	// check mimetype
	if fileForm.Header.Get("Content-Type") != "image/jpeg" && fileForm.Header.Get("Content-Type") != "image/png" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "Invalid file type",
			"err":     "Invalid file type",
		})
		return
	}

	var extension = strings.Split(fileForm.Filename, ".")[1]
	timeParts := strings.Split(time.Now().String(), " ")
	photo := string(timeParts[4][6:14] + "." + extension)

	var file string = "public/uploads/recipes/" + photo

	c.SaveUploadedFile(fileForm, file)

	category_id, _ := strconv.ParseUint(c.PostForm("category_id"), 10, 64)

	data := models.Recipe{
		CategoryId:  uint(category_id),
		Name:        c.PostForm("name"),
		Slug:        slug.Make(c.PostForm("name")),
		Time:        c.PostForm("time"),
		Description: c.PostForm("description"),
		Photo:       photo,
		Date:        time.Now(),
	}

	database.DB.Create(&data)

	c.JSON(http.StatusCreated, gin.H{
		"status":  "OK",
		"message": "Recipe created successfully",
		"data":    data,
	})
}
