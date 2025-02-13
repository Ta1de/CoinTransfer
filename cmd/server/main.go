package main

import (
	"CoinTransfer/internal/handler"
	"CoinTransfer/pkg/database"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	r := mux.NewRouter()

	// Маршрут для аутентификации
	r.HandleFunc("/api/auth", handler.AuthHandler).Methods("POST")

	// Запуск сервера
	http.ListenAndServe(":8080", r)
}
