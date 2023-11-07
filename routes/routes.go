package routes

import (
	"github.com/alessandro-maciel/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controllers.StudentIndex)
	r.POST("/students", controllers.StudentCreate)
	r.GET("/students/:id", controllers.StudentShow)
	r.PUT("/students/:id", controllers.StudentUpdate)
	r.DELETE("/students/:id", controllers.StudentDelete)

	r.Run()
}
