package service

import (
	"CoinTransfer/pkg/model"
	"CoinTransfer/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (string, error)
	CreateToken(username, password string) (string, error)
}

type Transfer interface {
	SendCoins(fromUserId int, toUser string, amount int) error
}

type Info interface {
	GetInfo(UserId int) (model.Info, error)
}

type BuyItem interface {
	BuyItemByName(ItemName string, UserId int) (model.Item, error)
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
