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
	TransferHistory(UserId int) (model.History, error)
	getReceivedTransactions(UserId int) ([]model.ReceivedTransaction, error)
	getSentTransactions(UserId int) ([]model.SentTransaction, error)
	GetCoins(userID int) (int, error)
}

type BuyItem interface {
	GetItem(ItemName string) (model.Item, error)
}

type Repository struct {
	Authorization
	Transfer
	Info
	BuyItem
}

func NewRepositore(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Transfer:      NewTransferPostgres(db),
		Info:          NewInfoPostgres(db),
		BuyItem:       NewBuyItemPostgres(db),
	}
}
