package handler

import (
	"CoinTransfer/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/auth", h.auth)
		api.POST("/sendCoin", h.SendCoin)
		api.GET("/info", h.getInfo)
		api.GET("/buy/{item}", h.buyItem)
	}
	return router
}
