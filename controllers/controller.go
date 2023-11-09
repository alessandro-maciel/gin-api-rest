package controllers

import (
	"net/http"

	"github.com/alessandro-maciel/gin-api-rest/database"
	"github.com/alessandro-maciel/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func StudentIndex(c *gin.Context) {
	cpf := c.Query("cpf")

	var students []models.Student
	database.DB.Where(&models.Student{CPF: cpf}).Find(&students)

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

	if err := models.ValidateData(&student); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	database.DB.Create(&student)

	c.JSON(http.StatusCreated, student)
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

func StudentUpdate(c *gin.Context) {
	var student models.Student

	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	var new_student models.Student

	if err := c.ShouldBindJSON(&new_student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateData(&new_student); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	database.DB.Model(&student).Updates(new_student)

	c.JSON(http.StatusOK, student)
}

func StudentDelete(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student
	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
