package usecase

import (
	"currency-rate/internal/usecase/webapi"
	"time"
)

type CurrencyDTO struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Name      string    `json:"name"`
	CharCode  string    `json:"char_code"`
	Value     float64   `json:"value"`
}

type CurrencyUseCase interface {
	GetCurrency(date time.Time, curName string) (*CurrencyDTO, error)
}

type CurrencyGateway interface {
	GetQuotes(date string) (*webapi.AllCur, error)
}

type CurrencyWorkerWithDB interface {
	GetCurrencyDate(date string, charName string) (*CurrencyDTO, error)
	InsertCurrencyDate(value []CurrencyDTO) error
}
