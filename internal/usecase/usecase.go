package usecase

import (
	"time"
)

//go:generate go run go.uber.org/mock/mockgen -source=usecase.go -destination=mocks/mock_interface.go

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
