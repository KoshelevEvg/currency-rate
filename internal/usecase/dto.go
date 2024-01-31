package usecase

import "time"

type Valute struct {
	ID       string  `json:"ID"`
	NumCode  string  `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Nominal  int     `json:"Nominal"`
	Name     string  `json:"Name"`
	Value    float64 `json:"Value"`
	Previous float64 `json:"Previous"`
}

type AllCur struct {
	Date         time.Time `json:"Date"`
	PreviousDate time.Time `json:"PreviousDate"`
	PreviousURL  string    `json:"PreviousURL"`
	Timestamp    time.Time `json:"Timestamp"`
	Valute       map[string]Valute
}

type CurrencyDTO struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Name      string    `json:"name"`
	CharCode  string    `json:"char_code"`
	Value     float64   `json:"value"`
}
