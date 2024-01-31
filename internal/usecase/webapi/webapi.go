package webapi

import (
	"currency-rate/internal/usecase"
	"encoding/json"
	"fmt"
	"net/http"
)

type WebServer struct {
	Addr string
}

func NewWeb(addr string) *WebServer {
	return &WebServer{
		Addr: addr,
	}
}

func (this *WebServer) GetQuotes(date string) (*usecase.AllCur, error) {
	str := fmt.Sprintf("%s/archive/%s/daily_json.js", this.Addr, date)
	response, err := http.Get(str)
	if err != nil {
		return nil, err
	}

	structResponse := new(usecase.AllCur)
	decoder := json.NewDecoder(response.Body)
	decoder.UseNumber()
	if err = decoder.Decode(structResponse); err != nil {
		return nil, err
	}
	return structResponse, err
}
