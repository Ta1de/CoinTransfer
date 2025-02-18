package service

import (
	"CoinTransfer/pkg/repository"
	"fmt"
)

type TransferService struct {
	repo repository.Transfer
}

func NewTransferService(repo repository.Transfer) *TransferService {
	return &TransferService{repo: repo}
}

func (s *TransferService) SendCoins(fromUserId int, toUser string, amount int) error {

	toUserId, err := s.repo.GetUserIDByUsername(toUser)
	if err != nil {
		return err
	}

	balance, err := s.repo.GetUserBalance(fromUserId)
	if err != nil {
		return err
	}

	if balance < amount {
		return fmt.Errorf("not enough balance")
	}

	err = s.repo.UpdateBalances(fromUserId, toUserId, amount)
	if err != nil {
		return err
	}

	err = s.repo.SaveTransfer(fromUserId, toUserId, amount)
	if err != nil {
		return err
	}

	return nil
}
