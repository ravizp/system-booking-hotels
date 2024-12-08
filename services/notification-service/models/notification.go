package models

import "time"

type Notification struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"userId"`
	Type      string    `json:"type"` // e.g., "booking", "payment", etc.
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
