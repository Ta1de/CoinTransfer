package models

type User struct {
	ID       int    `json:"-" db:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Coins    int    `json:"coins"`
}
