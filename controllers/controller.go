package controllers

import (
	"net/http"

	"github.com/alessandro-maciel/gin-api-rest/database"
	"github.com/alessandro-maciel/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func StudentIndex(c *gin.Context) {
	database.DB.Find(&models.Students)

	c.JSON(http.StatusOK, models.Students)
}

func StudentCreate(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&student)

	c.JSON(http.StatusOK, student)
}
