package service

import (
	"CoinTransfer/pkg/model"
	"CoinTransfer/pkg/repository"
)

type InfoService struct {
	repo repository.Info
}

func NewInfoService(repo repository.Info) *InfoService {
	return &InfoService{repo: repo}
}

func (s *InfoService) GetInfo(UserId int) (model.Info, error) {
	var info model.Info

	coins, err := s.repo.GetCoins(UserId)
	if err != nil {
		return info, err
	}
	info.Coins = coins

	items, err := s.repo.GetInventory(UserId)
	if err != nil {
		return info, err
	}
	info.Inventory = items

	history, err := s.repo.TransferHistory(UserId)
	if err != nil {
		return info, err
	}
	info.CoinHistory = history

	return info, nil
}
