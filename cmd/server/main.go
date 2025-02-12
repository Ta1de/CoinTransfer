package main

import (
	"log"
	"net/http"

	"CoinTransfer/internal/config"
	"CoinTransfer/internal/handler"
	"CoinTransfer/internal/middleware"
	"CoinTransfer/internal/repository"
	"CoinTransfer/internal/service"
	"CoinTransfer/pkg/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	router := http.NewServeMux()
	router.HandleFunc("/api/auth", authHandler.Authenticate)

	authMiddleware := middleware.AuthMiddleware(authService)

	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: authMiddleware(router),
	}

	log.Printf("Server starting on %s", cfg.ServerAddress)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
