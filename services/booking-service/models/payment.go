package models

import "time"

type Payment struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	BookingID     uint      `json:"bookingId"`
	Amount        int       `json:"amount"`
	PaymentMethod string    `json:"payment_method"` // e.g., card, paypal
	Status        string    `json:"status"`         // pending, completed, failed
	PaymentDate   time.Time `json:"payment_date"`
}
