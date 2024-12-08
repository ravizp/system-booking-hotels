package models

import (
	"fmt"
	"time"
	"user-service/config"

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
	UpdatedAt time.Time `json:"updated_at"`
}

var DB *gorm.DB

func InitDatabase() {
	config.LoadEnv() // Load environment variables

	var dbHost, dbPort, dbUser, dbPassword, dbName string
	if config.GetEnv("USE_LOCAL_DB", "false") == "true" {
		dbHost = config.GetEnv("DB_HOST_LOCAL", "localhost")
		dbPort = config.GetEnv("DB_PORT_LOCAL", "5432")
		dbUser = config.GetEnv("DB_USER_LOCAL", "postgres")
		dbPassword = config.GetEnv("DB_PASSWORD_LOCAL", "postgres")
		dbName = config.GetEnv("DB_NAME_LOCAL", "postgres")
	} else {
		dbHost = config.GetEnv("DB_HOST", "postgres")
		dbPort = config.GetEnv("DB_PORT", "5432")
		dbUser = config.GetEnv("DB_USER", "postgres")
		dbPassword = config.GetEnv("DB_PASSWORD", "postgres")
		dbName = config.GetEnv("DB_NAME", "SystemHotelDB")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	var err error
	for i := 1; i <= 5; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Database connected!")
			break
		}
		fmt.Printf("Failed to connect to database (%d/5): %s\n", i, err.Error())
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
}
