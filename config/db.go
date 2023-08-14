package config

import (
	"github.com/shamysharma/task1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "host=localhost user=postgres password=shamy@123 dbname=task1 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
}
