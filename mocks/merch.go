package mocks

import (
	"CoinTransfer/internal/models"

	"github.com/stretchr/testify/mock"
)

// MockInfo — заглушка для интерфейса repository.Info
type MockInfo struct {
	mock.Mock
}

func (m *MockInfo) GetCoins(userID int) (int, error) {
	args := m.Called(userID)
	return args.Int(0), args.Error(1)
}

func (m *MockInfo) GetInventory(userID int) (models.InventoryItems, error) {
	args := m.Called(userID)
	return args.Get(0).(models.InventoryItems), args.Error(1)
}

func (m *MockInfo) TransferHistory(userID int) (models.History, error) {
	args := m.Called(userID)
	return args.Get(0).(models.History), args.Error(1)
}
