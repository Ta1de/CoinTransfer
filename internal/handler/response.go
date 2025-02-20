package handler

import (
	"CoinTransfer/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, models.Error{
		StatusCode: statusCode,
		Message:    message,
	})
}
