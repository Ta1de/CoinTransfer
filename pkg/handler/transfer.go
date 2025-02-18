// handler/send_coin.go
package handler

import (
	"CoinTransfer/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SendCoin(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var input model.SendCoinRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Transfer.SendCoins(userId, input.ToUser, input.Amount)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Coins sent successfully"})
}
