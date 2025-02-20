package handler

import (
	"CoinTransfer/internal/middleware"
	"CoinTransfer/internal/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(authService *services.AuthService) *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/auth", h.auth)

		protected := api.Group("/", middleware.JWTAuthMiddleware(authService))
		{
			protected.POST("/sendCoin", h.SendCoin)
			protected.GET("/info", h.getInfo)
			protected.GET("/buy/:item", h.buyItem)
		}
	}
	return router
}

func getUserIdFromContext(c *gin.Context) (int, error) {
	userId, exists := c.Get("userId")
	if !exists {
		return 0, fmt.Errorf("user ID not found in context")
	}

	id, ok := userId.(int)
	if !ok {
		return 0, fmt.Errorf("user ID is not of type int")
	}

	return id, nil
}
