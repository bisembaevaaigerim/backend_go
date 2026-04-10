package database

import (
	"backend-go/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=13252021aigerim dbname=events port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	db.AutoMigrate(&models.Event{}, &models.Booking{}, &models.User{})

	DB = db
	log.Println("DB connected")
}
