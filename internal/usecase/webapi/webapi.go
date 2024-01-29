package webapi

import (
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

func (this *WebServer) Get(b string, date string) (*AllCur, error) {
	var str string
	if b == "Today" {
		str = fmt.Sprintf("%s/daily_json.js", this.Addr)
	} else {
		str = fmt.Sprintf("%s/archive/%s/daily_json.js", this.Addr, date)
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
