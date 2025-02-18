package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) getInfo(c *gin.Context) {
	// userId, err := getUserIdFromContext(c)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	// 	return
	// }

	// coins, err := h.services.GetCoins(userId)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, "Failed to get user coins")
	// 	return
	// }

	// inventory, err := h.services.GetInventory(userId)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, "Failed to get user inventory")
	// 	return
	// }

	// receivedHistory, err := h.services.GetReceivedHistory(userId)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, "Failed to get received coin history")
	// 	return
	// }

	// sentHistory, err := h.services.GetSentHistory(userId)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, "Failed to get sent coin history")
	// 	return
	// }

}
