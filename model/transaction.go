package model

import "time"

type Transaction struct {
	ID              int     `json:"id"`
	Booking         Booking `json:"booking_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	TransactionDate time.Time
	Status          string  `json:"status"`
}
