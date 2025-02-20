package services

import (
	"CoinTransfer/internal/models"
	"CoinTransfer/internal/repository"
	"fmt"
)

type BuyItemService struct {
	repo repository.BuyItem
}

func NewBuyItemService(repo repository.BuyItem) *BuyItemService {
	return &BuyItemService{repo: repo}
}

func (s *BuyItemService) BuyItemByName(ItemName string, UserId int) (models.Item, error) {
	var item models.Item
	item, err := s.repo.GetItem(ItemName)
	if err != nil {
		return item, err
	}

	balance, err := s.repo.GetBalance(UserId)
	if err != nil {
		return item, err
	}

	if item.Price > balance {
		return item, fmt.Errorf("not enough balance to buy %s", ItemName)
	}

	err = s.repo.UpdateBalance(UserId, item.Price)
	if err != nil {
		return item, err
	}

	err = s.repo.AddToInventory(UserId, ItemName)
	if err != nil {
		return item, err
	}

	return item, err
}
