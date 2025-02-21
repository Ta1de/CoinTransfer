package services

import (
	"CoinTransfer/internal/models"
	"CoinTransfer/internal/repository"
)

type InfoService struct {
	repo repository.Info
}

func NewInfoService(repo repository.Info) *InfoService {
	return &InfoService{repo: repo}
}

func (s *InfoService) GetInfo(UserId int) (models.Info, error) {
	var info models.Info

	coins, err := s.repo.GetCoins(UserId)
	if err != nil {
		return info, err
	}

	items, err := s.repo.GetInventory(UserId)
	if err != nil {
		return info, err
	}

	history, err := s.repo.TransferHistory(UserId)
	if err != nil {
		return info, err
	}
	info.Coins = coins
	info.Inventory = items
	info.CoinHistory = history

	return info, nil
}
