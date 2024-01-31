package v1

import (
	"currency-rate/internal/usecase"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, currencyGetter usecase.CurrencyUseCase) {
	controller := NewCurrencyController(currencyGetter)
	api := router.Group("/api/v1")
	api.GET("/currency", controller.GetCur)
}
