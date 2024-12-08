package models

import "time"

type Booking struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `json:"userId"`
	RoomID       uint      `json:"roomId"`
	CheckInDate  time.Time `json:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date"`
	Status       string    `json:"status" gorm:"default:pending"` // pending, confirmed, cancelled
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
