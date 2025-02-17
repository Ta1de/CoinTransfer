package service

import (
	"CoinTransfer/pkg/repository"
)

type TransferService struct {
	repo repository.Transfer
}

func NewTransferService(repo repository.Transfer) *TransferService {
	return &TransferService{repo: repo}
}

func (s *TransferService) SendCoins(fromUserId int, toUser string, amount int) error {

	// Сохраняем информацию о транзакции в базу данных
	err := s.repo.SaveTransfer(fromUserId, toUser, amount)
	if err != nil {
		return err
	}

	return nil
}
