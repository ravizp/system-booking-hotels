package models

import "time"

type Refund struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	BookingID    uint      `json:"bookingId"`
	Amount       int       `json:"amount"`
	Reason       string    `json:"reason"`
	Status       string    `json:"status"` // pending, approved, rejected
	RequestDate  time.Time `json:"request_date"`
	ApprovedDate *time.Time `json:"approved_date,omitempty"`
}
