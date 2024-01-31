package interactor

import (
	"currency-rate/internal/usecase"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type GetCurrencyWithDate struct {
	w usecase.CurrencyGateway
	r usecase.CurrencyCache
}

func NewGetCurrency(w usecase.CurrencyGateway, r usecase.CurrencyCache) *GetCurrencyWithDate {
	return &GetCurrencyWithDate{
		w: w,
		r: r,
	}
}

func (g GetCurrencyWithDate) GetCurrency(date time.Time, nameCur string) (*usecase.CurrencyDTO, error) {
	fmtString := replacingDash(date.String())
	currency, err := g.r.GetCurrencyDate(fmtString, nameCur)
	if err != nil {
		currencyList := make([]usecase.CurrencyDTO, 0)
		response, err := g.w.GetQuotes(fmtString)
		if err != nil {
			return nil, err
		}
		if _, ok := response.Valute[nameCur]; !ok {
			return nil, err //TODO Создать ошибку, что нет валюты
		}
		for _, v := range response.Valute {
			cur := usecase.CurrencyDTO{
				StartDate: response.Date,
				EndDate:   response.PreviousDate,
				Name:      v.Name,
				CharCode:  v.CharCode,
				Value:     v.Value,
			}
			currencyList = append(currencyList, cur)
		}

		if err = g.r.InsertCurrencyDate(currencyList); err != nil {
			logrus.Error(err) //TODO добавить контекст
		}
		return &usecase.CurrencyDTO{
			StartDate: response.Date,
			EndDate:   response.PreviousDate,
			Name:      response.Valute[nameCur].Name,
			CharCode:  response.Valute[nameCur].CharCode,
			Value:     response.Valute[nameCur].Value,
		}, nil
	}

	return &usecase.CurrencyDTO{
		StartDate: currency.StartDate,
		EndDate:   currency.EndDate,
		Name:      currency.Name,
		CharCode:  currency.CharCode,
		Value:     currency.Value,
	}, nil
}

func replacingDash(b string) string {
	b = strings.Replace(b, "-", "/", 2)
	fmtString := strings.Split(b, " ")
	return fmtString[0]
}
