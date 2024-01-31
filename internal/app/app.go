package app

import (
	"currency-rate/config"
	"currency-rate/internal/controller/http/v1"
	"currency-rate/internal/usecase/interactor"
	"currency-rate/internal/usecase/repo"
	"currency-rate/internal/usecase/webapi"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Run(cfg *config.Config) error {
	db := NewConnectDB(cfg.DB)
	router := gin.New()
	api := webapi.NewWeb(cfg.Address)
	repoDB := repo.NewCurrencyDB(db)
	getter := interactor.NewGetCurrency(api, repoDB)
	v1.Register(router, getter)

	srv := &http.Server{
		Addr:           cfg.ServerAddress,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return srv.ListenAndServe()
}
