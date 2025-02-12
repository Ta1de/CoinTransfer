package model

import "time"

type Transaction struct {
	ID         int       `json:"id"`
	FromUserID int       `json:"from_user_id"`
	ToUserID   int       `json:"to_user_id"`
	Amount     int       `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
}

type TransactionHistory struct {
	Received []Transaction `json:"received"`
	Sent     []Transaction `json:"sent"`
}
