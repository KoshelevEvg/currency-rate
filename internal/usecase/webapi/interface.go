package webapi

type GetList interface {
	Get(b string, date string) (*AllCur, error)
}
