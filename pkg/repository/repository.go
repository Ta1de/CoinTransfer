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
	GetUserBalance(userID int) (int, error)
	UpdateBalances(senderID, receiverID, amount int) error
	SaveTransfer(fromUserId, toUser, amount int) error
}

type Info interface {
	GetInventory(UserId int) (model.InventoryItems, error)
}

type Repository struct {
	Authorization
	Transfer
	Info
}

func NewRepositore(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Transfer:      NewTransferPostgres(db),
		Info:          NewInfoPostgres(db),
	}
}
