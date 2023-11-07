package database

import (
	"log"

	"github.com/alessandro-maciel/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	connect := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(connect))
	if err != nil {
		log.Panic("Error connecting to database")
	}

	DB.AutoMigrate(&models.Student{})
}
