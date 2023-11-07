package controllers

import (
	"net/http"

	"github.com/alessandro-maciel/gin-api-rest/database"
	"github.com/alessandro-maciel/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func StudentIndex(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)

	c.JSON(http.StatusOK, students)
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

func StudentShow(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "student not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}
