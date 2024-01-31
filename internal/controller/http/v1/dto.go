package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type errorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Name      string    `json:"name"`
	CharCode  string    `json:"char_code"`
	Value     float64   `json:"value"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
