package handlers

import "github.com/gin-gonic/gin"

func Example_Get(c *gin.Context){
	c.JSON(200, gin.H{
		"status": "OK",
		"message": "message from my example route",
	})
}