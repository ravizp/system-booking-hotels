package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Connect() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
