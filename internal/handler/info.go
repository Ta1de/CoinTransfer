package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getInfo(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	info, err := h.services.GetInfo(userId)
	if err != nil {
		log.Printf("Error in getInfo: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, "Failed to get user info")
		return
	}

	c.JSON(http.StatusOK, info)
}
