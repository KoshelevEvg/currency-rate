package usecase

import (
	"time"
)

type Currency struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type CurrencyGetter interface {
	GetCurrency(date time.Time, curName string) (*Currency, error)
}
