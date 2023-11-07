package main

import (
	"github.com/alessandro-maciel/gin-api-rest/database"
	"github.com/alessandro-maciel/gin-api-rest/routes"
)

func main() {
	database.ConnectDatabase()
	routes.HandleRequests()
}
