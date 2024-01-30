package webapi

type CurrencyGateway interface {
	Get(b string, date string) (*AllCur, error)
}
