package service

import (
	"CoinTransfer/pkg/model"
	"CoinTransfer/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
