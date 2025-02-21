package service_test

import (
	"CoinTransfer/internal/models"
	"CoinTransfer/internal/services"
	"CoinTransfer/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInfo_Success(t *testing.T) {
	mockRepo := new(mocks.MockInfo)
	service := services.NewInfoService(mockRepo)

	userID := 1
	expectedCoins := 100
	expectedInventory := models.InventoryItems{
		Inventory: []models.InventoryItem{
			{Type: "sword", Quantity: 1},
			{Type: "shield", Quantity: 2},
		},
	}
	expectedHistory := models.History{
		Received: []models.ReceivedTransaction{
			{FromUser: "Alice", Amount: 50},
		},
		Sent: []models.SentTransaction{
			{ToUser: "Bob", Amount: 20},
		},
	}

	// Настраиваем мок-методы
	mockRepo.On("GetCoins", userID).Return(expectedCoins, nil)
	mockRepo.On("GetInventory", userID).Return(expectedInventory, nil)
	mockRepo.On("TransferHistory", userID).Return(expectedHistory, nil)

	// Вызываем тестируемый метод
	info, err := service.GetInfo(userID)

	// Проверяем результаты
	assert.NoError(t, err)
	assert.Equal(t, expectedCoins, info.Coins)
	assert.Equal(t, expectedInventory, info.Inventory)
	assert.Equal(t, expectedHistory, info.CoinHistory)

	// Проверяем, что все вызовы были сделаны
	mockRepo.AssertExpectations(t)
}

func TestGetInfo_Failure_GetCoins(t *testing.T) {
	mockRepo := new(mocks.MockInfo)
	service := services.NewInfoService(mockRepo)

	userID := 1
	mockRepo.On("GetCoins", userID).Return(0, assert.AnError)

	info, err := service.GetInfo(userID)

	assert.Error(t, err)
	assert.Empty(t, info)
	mockRepo.AssertExpectations(t)
}

func TestGetInfo_Failure_GetInventory(t *testing.T) {
	mockRepo := new(mocks.MockInfo)
	service := services.NewInfoService(mockRepo)

	userID := 1
	expectedCoins := 100
	expectedError := errors.New("failed to fetch inventory")

	// Настраиваем мок-методы
	mockRepo.On("GetCoins", userID).Return(expectedCoins, nil)
	mockRepo.On("GetInventory", userID).Return(models.InventoryItems{}, expectedError)

	// Вызываем тестируемый метод
	info, err := service.GetInfo(userID)

	// Проверяем, что ошибка была возвращена
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)

	// Проверяем, что возвращена пустая структура
	assert.Equal(t, models.Info{}, info)

	// Проверяем вызовы мок-методов
	mockRepo.AssertExpectations(t)
}

func TestGetInfo_Failure_TransferHistory(t *testing.T) {
	mockRepo := new(mocks.MockInfo)
	service := services.NewInfoService(mockRepo)

	userID := 1
	expectedCoins := 100
	expectedInventory := models.InventoryItems{
		Inventory: []models.InventoryItem{{Type: "sword", Quantity: 1}},
	}
	expectedError := errors.New("failed to fetch transfer history")

	// Настраиваем мок-методы
	mockRepo.On("GetCoins", userID).Return(expectedCoins, nil)
	mockRepo.On("GetInventory", userID).Return(expectedInventory, nil)
	mockRepo.On("TransferHistory", userID).Return(models.History{}, expectedError)

	// Вызываем тестируемый метод
	info, err := service.GetInfo(userID)

	// Проверяем, что ошибка была возвращена
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)

	// Проверяем, что возвращена пустая структура
	assert.Equal(t, models.Info{}, info)

	// Проверяем вызовы мок-методов
	mockRepo.AssertExpectations(t)
}
