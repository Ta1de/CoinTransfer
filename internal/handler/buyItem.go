package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) buyItem(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	itemName := c.Param("item")

	item, err := h.services.BuyItemByName(itemName, userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Item purchased successfully",
		"item":    item,
	})
}
