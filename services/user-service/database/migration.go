package database

import "user-service/models"

func RunMigrations() {
	models.DB.AutoMigrate(&models.User{})
}	