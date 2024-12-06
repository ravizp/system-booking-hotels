package models

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Role      string    `json:"role" gorm:"default:'customer'"`
	CreatedAt time.Time `json:"created_at"`
	UpadtedAt time.Time `json:"updated_at"`
}

var DB *gorm.DB

func init() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=SystemHotelDB port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&User{})
}