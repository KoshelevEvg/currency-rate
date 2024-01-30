package usecase

import (
	"strings"
	"time"
)

type GetCurrencyWithDate struct {
	w CurrencyGateway
	r CurrencyWorkerWithDB
}

func NewGetCurrency(w CurrencyGateway, r CurrencyWorkerWithDB) *GetCurrencyWithDate {
	return &GetCurrencyWithDate{
		w: w,
		r: r,
	}
}

func (this *GetCurrencyWithDate) GetCurrency(date time.Time, nameCur string) (*CurrencyDTO, error) {
	fmtString := replacingDash(date.String())
	currency, err := this.r.GetCurrencyDate(fmtString, nameCur)
	if err != nil {
		currencyList := make([]CurrencyDTO, 0)
		response, err := this.w.GetQuotes(fmtString)
		if _, ok := response.Valute[nameCur]; !ok {
			return nil, err //TODO добавить описание ошибки
		}
		if err != nil {
			return nil, err
		}
		for _, v := range response.Valute {
			cur := CurrencyDTO{
				StartDate: response.Date,
				EndDate:   response.PreviousDate,
				Name:      v.Name,
				CharCode:  v.CharCode,
				Value:     v.Value,
			}
			currencyList = append(currencyList, cur)
		}

		err = this.r.InsertCurrencyDate(currencyList)
		if err != nil {
			return nil, err
		}
		return &CurrencyDTO{
			StartDate: response.Date,
			EndDate:   response.PreviousDate,
			Name:      response.Valute[nameCur].Name,
			CharCode:  response.Valute[nameCur].CharCode,
			Value:     response.Valute[nameCur].Value,
		}, nil

	} else {
		return &CurrencyDTO{
			StartDate: currency.StartDate,
			EndDate:   currency.EndDate,
			Name:      currency.Name,
			CharCode:  currency.CharCode,
			Value:     currency.Value,
		}, nil
	}
}

func replacingDash(b string) string {
	b = strings.Replace(b, "-", "/", 2)
	fmtString := strings.Split(b, " ")
	return fmtString[0]
}
