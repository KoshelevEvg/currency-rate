package usecase

import (
	"currency-rate/internal/usecase/webapi"
	"strings"
	"time"
)

func GetCur(date time.Time, nameCur string) (*Currency, error) {
	var valStr string
	year, month, day := date.Date()
	yearToDay, monthToDay, dayToDay := time.Now().Date()

	fmtString := replace(date.String())

	if year == yearToDay && month == monthToDay && day == dayToDay {
		valStr = "Today"
	} else {
		valStr = "Not today"
	}
	a, err := webapi.Get(valStr, fmtString)

	if err != nil {
		return nil, err
	}

	if _, ok := a.Valute[nameCur]; !ok {
		return nil, err
	}

	return &Currency{
		Name:  a.Valute[nameCur].Name,
		Value: a.Valute[nameCur].Value,
	}, nil
}

func replace(b string) string {
	b = strings.Replace(b, "-", "/", 2)
	fmtString := strings.Split(b, " ")
	return fmtString[0]
}
