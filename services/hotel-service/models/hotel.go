package models

import "time"

type Hotel struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Address     string    `gorm:"not null" json:"address"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Rooms       []Room    `gorm:"foreignKey:HotelID" json:"rooms"`
}

type Room struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	HotelID   uint      `json:"hotel_id"`
	Number    string    `gorm:"size:10;not null" json:"number"`
	Type      string    `gorm:"size:50;not null" json:"type"`
	Price     float64   `gorm:"not null" json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}