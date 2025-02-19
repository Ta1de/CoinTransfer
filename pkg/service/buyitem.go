package service

import (
	"CoinTransfer/pkg/repository"
)

type BuyItemService struct {
	repo repository.BuyItem
}

func NewBuyItemService(repo repository.BuyItem) *BuyItemService {
	return &BuyItemService{repo: repo}
}

func (s *BuyItemService) BuyItem() (string, error) {

}
