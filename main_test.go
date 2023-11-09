package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/alessandro-maciel/gin-api-rest/controllers"
	"github.com/alessandro-maciel/gin-api-rest/database"
	"github.com/alessandro-maciel/gin-api-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var Students []models.Student
var Routes *gin.Engine

func Setup() {
	database.ConnectDatabaseSqlite()

	database.DB.Migrator().DropTable(&models.Student{})
	database.DB.AutoMigrate(&models.Student{})

	gin.SetMode(gin.ReleaseMode)
	Routes = gin.Default()
}

func TestListStudents(t *testing.T) {
	Setup()

	CreateStudentsMock()

	Routes.GET("/students", controllers.StudentIndex)

	request, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()

	Routes.ServeHTTP(response, request)

	var students_response []models.Student
	json.Unmarshal([]byte(response.Body.Bytes()), &students_response)

	assert.Equal(t, http.StatusOK, response.Code, "Expected 200 and returned %d", response.Code)
	assert.Equal(t, Students, students_response)
}

func TestCreateStudent(t *testing.T) {
	Setup()

	Routes.POST("/students", controllers.StudentCreate)

	var student = models.Student{
		Name: "Student 1", CPF: "11111111111", RG: "111111111",
	}

	data, _ := json.Marshal(student)

	request, _ := http.NewRequest("POST", "/students", bytes.NewBuffer(data))
	response := httptest.NewRecorder()

	Routes.ServeHTTP(response, request)

	var student_response models.Student
	json.Unmarshal([]byte(response.Body.Bytes()), &student_response)

	assert.Equal(t, http.StatusCreated, response.Code, "Expected 201 and returned %d", response.Code)
	assert.Equal(t, student.Name, student_response.Name)
	assert.Equal(t, student.CPF, student_response.CPF)
	assert.Equal(t, student.RG, student_response.RG)

	var student_database models.Student
	database.DB.First(&student_database, student_response.ID)

	assert.Equal(t, student_database.ID, student_response.ID)
	assert.Equal(t, student_database.Name, student_response.Name)
	assert.Equal(t, student_database.CPF, student_response.CPF)
	assert.Equal(t, student_database.RG, student_response.RG)
}

func TestShowStudent(t *testing.T) {
	Setup()

	CreateStudentsMock()

	Routes.GET("/students/:id", controllers.StudentShow)

	student := Students[len(Students)-1]
	student_id := student.ID
	path := "/students/" + strconv.Itoa(int(student_id))
	request, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()

	Routes.ServeHTTP(response, request)

	var student_response models.Student
	json.Unmarshal([]byte(response.Body.Bytes()), &student_response)

	assert.Equal(t, http.StatusOK, response.Code, "Expected 200 and returned %d", response.Code)
	assert.Equal(t, student.ID, student_response.ID)
	assert.Equal(t, student.Name, student_response.Name)
	assert.Equal(t, student.CPF, student_response.CPF)
	assert.Equal(t, student.RG, student_response.RG)
}

func TestUpdateStudent(t *testing.T) {
	Setup()

	CreateStudentsMock()

	Routes.PUT("/students/:id", controllers.StudentUpdate)

	updated_student := models.Student{
		Name: "updated student", CPF: "33333333333", RG: "333333333",
	}
	data, _ := json.Marshal(updated_student)

	student := Students[len(Students)-1]
	student_id := student.ID
	path := "/students/" + strconv.Itoa(int(student_id))
	request, _ := http.NewRequest("PUT", path, bytes.NewBuffer(data))
	response := httptest.NewRecorder()

	Routes.ServeHTTP(response, request)

	var student_response models.Student
	json.Unmarshal([]byte(response.Body.Bytes()), &student_response)

	assert.Equal(t, http.StatusOK, response.Code, "Expected 200 and returned %d", response.Code)

	assert.Equal(t, student.ID, student_response.ID)
	assert.Equal(t, updated_student.Name, student_response.Name)
	assert.Equal(t, updated_student.CPF, student_response.CPF)
	assert.Equal(t, updated_student.RG, student_response.RG)

	var student_database models.Student
	database.DB.First(&student_database, student.ID)

	assert.Equal(t, student_id, student_database.ID)
	assert.Equal(t, updated_student.Name, student_database.Name)
	assert.Equal(t, updated_student.CPF, student_database.CPF)
	assert.Equal(t, updated_student.RG, student_database.RG)
}

func TestDeleteStudent(t *testing.T) {
	Setup()

	CreateStudentsMock()

	Routes.DELETE("/students/:id", controllers.StudentDelete)

	student := Students[len(Students)-1]
	student_id := student.ID
	path := "/students/" + strconv.Itoa(int(student_id))
	request, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()

	Routes.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code, "Expected 200 and returned %d", response.Code)
	assert.Equal(t, `{"success":true}`, response.Body.String())

	var student_database models.Student
	database.DB.First(&student_database, student.ID)

	assert.Equal(t, 0, int(student_database.ID))
}

func CreateStudentsMock() {
	var students = []models.Student{
		{Name: "Student 1", CPF: "11111111", RG: "111111111"},
		{Name: "Student 2", CPF: "22222222", RG: "222222222"},
	}

	for _, student := range students {
		database.DB.Create(&student)
		Students = append(Students, student)
	}
}
