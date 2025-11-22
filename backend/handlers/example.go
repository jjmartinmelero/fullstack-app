package handlers

import (
	"backend-fullstack-app/dto"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Example_get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "message get v2",
	})
}

func Example_get_with_params(c *gin.Context) {
	fmt.Println(c.Param("id"))
	fmt.Println(reflect.TypeOf(c.Param("id")))
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Get - Id is : " + c.Param("id"),
	})
}

func Example_post(c *gin.Context) {

	var dto dto.ExampleDto

	// Validate that JSON can be parsed correctly
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "Invalid data: " + err.Error(),
		})
		return
	}

	// Validate that required fields are not empty
	if dto.Email == "" || dto.Password == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "ERROR",
			"message": "Email and password fields are required",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "OK",
		"message": "Post - message post v2",
		"data":    dto,
	})
}

func Example_put(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Put - Id is : " + c.Param("id"),
	})
}

func Example_delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Delete - Id is : " + c.Param("id"),
	})
}

func Example_querystring(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Querystring - Id is : " + c.Query("id"),
	})
}

func Example_upload(c *gin.Context) {

	photo, err := c.FormFile("photo")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "Invalid data: " + err.Error(),
		})
		return
	}

	var extension = strings.Split(photo.Filename, ".")[1]
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fileToSave := timestamp + "." + extension

	c.SaveUploadedFile(photo, "public/uploads/images/"+fileToSave)

	c.JSON(http.StatusCreated, gin.H{
		"status":  "OK",
		"message": "Upload - message upload v2",
		"data":    photo,
	})
}
