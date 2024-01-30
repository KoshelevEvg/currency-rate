package usecase

import (
	"currency-rate/internal/usecase/webapi"
	"strings"
	"time"
)

type GetC struct {
	w webapi.CurrencyGateway
}

func NewGetCurrency(w webapi.CurrencyGateway) *GetC {
	return &GetC{
		w: w,
	}
}

func (this *GetC) GetCurrency(date time.Time, nameCur string) (*Currency, error) {
	var valStr string
	year, month, day := date.Date()
	yearToDay, monthToDay, dayToDay := time.Now().Date()

	fmtString := replace(date.String())

	if year == yearToDay && month == monthToDay && day == dayToDay {
		valStr = "Today"
	} else {
		valStr = "Not today"
	}

	currency, err := this.w.Get(valStr, fmtString)
	if err != nil {
		return nil, err
	}

	if _, ok := currency.Valute[nameCur]; !ok {
		return nil, err //TODO добавить описание ошибки
	}

	return &Currency{
		StartDate: currency.Date,
		EndDate:   currency.PreviousDate,
		Name:      currency.Valute[nameCur].Name,
		CharCode:  currency.Valute[nameCur].CharCode,
		Value:     currency.Valute[nameCur].Value,
	}, nil
}

func replace(b string) string {
	b = strings.Replace(b, "-", "/", 2)
	fmtString := strings.Split(b, " ")
	return fmtString[0]
}
