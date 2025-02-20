package services

import (
	"CoinTransfer/internal/models"
	"CoinTransfer/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	CreateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Transfer interface {
	SendCoins(fromUserId int, toUser string, amount int) error
}

type Info interface {
	GetInfo(UserId int) (models.Info, error)
}

type BuyItem interface {
	BuyItemByName(ItemName string, UserId int) (models.Item, error)
}

type Service struct {
	Authorization
	Transfer
	Info
	BuyItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Transfer:      NewTransferService(repos.Transfer),
		Info:          NewInfoService(repos.Info),
		BuyItem:       NewBuyItemService(repos.BuyItem),
	}
}
