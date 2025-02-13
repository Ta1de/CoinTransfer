package handler

import (
	"CoinTransfer/internal/model"
	"CoinTransfer/internal/repository"
	"CoinTransfer/internal/utils"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Обработчик аутентификации и регистрации пользователя
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ищем пользователя в базе
	existingUser, err := repository.GetUserByUsername(user.Username)
	if err != nil {
		// Если пользователь не найден, создаем нового
		existingUser = &model.User{
			Username: user.Username,
		}

		// Хешируем пароль
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}
		existingUser.Password = string(hashedPassword)

		// Сохраняем нового пользователя в базе
		err = repository.CreateUser(existingUser)
		if err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}
	}

	// Проверка пароля
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Генерация JWT
	token, err := utils.GenerateJWT(existingUser.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Отправка токена
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
