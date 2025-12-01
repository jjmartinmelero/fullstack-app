package main

import (
	"backend-fullstack-app/handlers"
	"backend-fullstack-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Run database migrations
	models.Migration()

	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	// exec migrations
	models.Migration()

	// static files
	router.Static("/public", "./public")

	// home path
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "Welcome to the API",
		})
	})

	// 69u94J1AIc

	// path not found
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "ERROR",
			"message": "Path not found",
		})
	})

	var prefix = "/api/v1/"

	router.GET(prefix+"/categories", handlers.GetCategories)
	router.GET(prefix+"/categories/:id", handlers.GetCategoryById)
	router.POST(prefix+"/categories", handlers.CreateCategory)
	router.PUT(prefix+"/categories/:id", handlers.EditCategory)
	router.DELETE(prefix+"/categories/:id", handlers.DeleteCategory)

	router.GET(prefix+"/recipes", handlers.GetRecipes)
	router.GET(prefix+"/recipes/:id", handlers.GetRecipeById)

	router.Run(":1024") // listens on 0.0.0.0:8080 by default
}
