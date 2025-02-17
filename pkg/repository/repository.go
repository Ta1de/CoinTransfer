package repository

import (
	"CoinTransfer/pkg/model"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) error
	GetUser(username, password string) (model.User, error)
}

type Transfer interface {
	GetUserIDByUsername(username string) (int, error)
	GetUserBalance(userID int) (float64, error)
	UpdateBalances(senderID, receiverID int, amount float64) error
	SaveTransfer(fromUserId int, toUser string, amount int) error
}

type Repository struct {
	Authorization
	Transfer
}

func NewRepositore(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Transfer:      NewTransferPostgres(db),
	}
}
