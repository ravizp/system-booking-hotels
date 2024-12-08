package database

import (
	"booking-service/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(host, user, password, dbname, port string) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected")
	DB = db

	// Migrate tables
	err = DB.AutoMigrate(
		&models.Booking{},
		&models.Payment{},
		&models.Refund{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

