package repository

import (
	"CoinTransfer/internal/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) error
	GetUser(username, password string) (models.User, error)
}

type Transfer interface {
	GetUserIDByUsername(username string) (int, error)
	GetUserBalance(userID int) (int, error)
	UpdateBalances(senderID, receiverID, amount int) error
	SaveTransfer(fromUserId, toUser, amount int) error
}

type Info interface {
	GetInventory(UserId int) (models.InventoryItems, error)
	TransferHistory(UserId int) (models.History, error)
	getReceivedTransactions(UserId int) ([]models.ReceivedTransaction, error)
	getSentTransactions(UserId int) ([]models.SentTransaction, error)
	GetCoins(userID int) (int, error)
}

type BuyItem interface {
	GetItem(ItemName string) (models.Item, error)
	GetBalance(userID int) (int, error)
	AddToInventory(userID int, itemName string) error
	UpdateBalance(senderID, amount int) error
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
