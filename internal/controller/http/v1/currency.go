package v1

import (
	"currency-rate/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetCur(c *gin.Context) {

	dateNew, err := time.Parse("2006/01/02", c.Query("date"))
	if err != nil {
		dateNew = time.Now()
	}

	val := c.Query("val")

	a, err := usecase.GetCur(dateNew, val)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
	}

	c.JSON(http.StatusOK, a)

}
