package handlers

import (
	"backend-fullstack-app/database"
	"backend-fullstack-app/dto"
	"backend-fullstack-app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Contact(c *gin.Context) {
	var contact dto.ContactDTO

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": "Failed to bind JSON",
			"err":     err.Error(),
		})
		return
	}

	data := models.Contact{
		Name:    contact.Name,
		Email:   contact.Email,
		Phone:   contact.Phone,
		Message: contact.Message,
		Date:    time.Now(),
	}

	database.DB.Create(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Contact created successfully",
		"data":    data,
	})
}
