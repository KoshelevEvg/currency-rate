package v1

import (
	"currency-rate/internal/constant"
	"currency-rate/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CurrencyController struct {
	getter usecase.CurrencyUseCase
}

func NewCurrencyController(getter usecase.CurrencyUseCase) *CurrencyController {
	return &CurrencyController{getter: getter}
}

func (c CurrencyController) GetCur(ctx *gin.Context) {
	dateNew, err := time.Parse(constant.CustomTimeLayout, ctx.Query("date"))
	if err != nil {
		dateNew = time.Now()
	}

	val := ctx.Query("val")

	currency, err := c.getter.GetCurrency(dateNew, val)
	if err != nil {
		newErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	resp := &SuccessResponse{
		StartDate: currency.StartDate,
		EndDate:   currency.EndDate,
		Name:      currency.Name,
		CharCode:  currency.CharCode,
		Value:     currency.Value,
	}

	ctx.JSON(http.StatusOK, resp)

}
