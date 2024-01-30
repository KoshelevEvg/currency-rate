package usecase

import (
	"time"
)

type Currency struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Name      string    `json:"name"`
	CharCode  string    `json:"char_code"`
	Value     float64   `json:"value"`
}

type CurrencyUseCase interface {
	GetCurrency(date time.Time, curName string) (*Currency, error)
}
