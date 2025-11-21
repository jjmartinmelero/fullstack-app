package main

import (
	"backend-fullstack-app/handlers"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	
	var prefix = "/api/v1/"
	// routes
	router.GET(prefix + "/example", handlers.Example_Get)

	router.Run(":1024") // listens on 0.0.0.0:8080 by default
}
