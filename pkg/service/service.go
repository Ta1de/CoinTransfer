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

type Service struct {
	Authorization
	Transfer
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Transfer:      NewTransferService(repos.Transfer),
	}
}
