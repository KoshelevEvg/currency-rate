package v1

import (
	"currency-rate/internal/usecase"
	"currency-rate/internal/usecase/webapi"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetCur(c *gin.Context) {

	dateNew, err := time.Parse("2006/01/02", c.Query("date"))
	if err != nil {
		dateNew = time.Now()
	}

	api := webapi.NewWeb("https://www.cbr-xml-daily.ru")

	s := usecase.NewGetCurrency(api)
	val := c.Query("val")

	a, err := s.GetCurrency(dateNew, val)

	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
	}

	c.JSON(http.StatusOK, a)

}
