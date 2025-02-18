package service

import (
	"CoinTransfer/pkg/repository"
)

type InfoService struct {
	repo repository.Info
}

func NewInfoService(repo repository.Info) *InfoService {
	return &InfoService{repo: repo}
}

func (s *InfoService) GetInfo(UserId int) error {
	return nil
}
