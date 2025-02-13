package repository

import (
	"CoinTransfer/internal/model"
	"CoinTransfer/pkg/database"
	"log"
)

// Получение пользователя по имени
func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func CreateUser(user *model.User) error {
	result := database.DB.Create(&user)
	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error) // Логирование ошибки
		return result.Error
	}
	log.Printf("User %s created successfully", user.Username) // Логирование успешного создания
	return nil
}
