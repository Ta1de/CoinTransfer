package model

type SendCoinRequest struct {
	ToUser string `json:"ToUser" binding:"required"`
	Amount int    `json:"amount" binding:"required,gt=0"`
}
