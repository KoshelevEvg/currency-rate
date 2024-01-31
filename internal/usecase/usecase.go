package usecase

import (
	"time"
)

type CurrencyUseCase interface {
	GetCurrency(date time.Time, curName string) (*CurrencyDTO, error)
}

type CurrencyGateway interface {
	GetQuotes(date string) (*AllCur, error)
}

type CurrencyCache interface {
	GetCurrencyDate(date string, charName string) (*CurrencyDTO, error)
	InsertCurrencyDate(value []CurrencyDTO) error
}
