package routes

import (
	"github.com/alessandro-maciel/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controllers.DisplaysAllStudents)

	r.Run()
}