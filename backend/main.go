package main

import (
	"backend-fullstack-app/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	// static files
	router.Static("/public", "./public")

	// home path
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "Welcome to the API",
		})
	})

	// path not found
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "ERROR",
			"message": "Path not found",
		})
	})

	var prefix = "/api/v1/"
	// routes
	router.GET(prefix+"/example", handlers.Example_get)
	router.GET(prefix+"/example/:id", handlers.Example_get_with_params)
	router.POST(prefix+"/example", handlers.Example_post)
	router.PUT(prefix+"/example/:id", handlers.Example_put)
	router.DELETE(prefix+"/example/:id", handlers.Example_delete)
	router.GET(prefix+"/example/querystring", handlers.Example_querystring)
	router.POST(prefix+"/example/upload", handlers.Example_upload)

	router.Run(":1024") // listens on 0.0.0.0:8080 by default
}
