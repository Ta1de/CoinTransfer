package mocks

import (
	"CoinTransfer/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockBuyItemRepository struct {
	mock.Mock
}

func (m *MockBuyItemRepository) GetItem(name string) (models.Item, error) {
	args := m.Called(name)
	return args.Get(0).(models.Item), args.Error(1)
}

func (m *MockBuyItemRepository) GetBalance(userID int) (int, error) {
	args := m.Called(userID)
	return args.Int(0), args.Error(1)
}

func (m *MockBuyItemRepository) UpdateBalance(userID int, amount int) error {
	args := m.Called(userID, amount)
	return args.Error(0)
}

func (m *MockBuyItemRepository) AddToInventory(userID int, itemName string) error {
	args := m.Called(userID, itemName)
	return args.Error(0)
}
