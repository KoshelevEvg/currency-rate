package webapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Get(b string, date string) (*AllCur, error) {
	var str string
	if b == "Today" {
		str = fmt.Sprint("https://www.cbr-xml-daily.ru/daily_json.js")
	} else {
		str = fmt.Sprintf("https://www.cbr-xml-daily.ru/archive/%s/daily_json.js", date)
	}
	a, err := http.Get(str)
	if err != nil {
		return nil, err
	}

	structResponse := new(AllCur)
	decoder := json.NewDecoder(a.Body)
	decoder.UseNumber()
	err = decoder.Decode(structResponse)
	if err != nil {
		return nil, err
	}
	return structResponse, err
}
