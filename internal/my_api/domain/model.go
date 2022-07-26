package domain

import "sync"

type Api struct {
	Id   uint64 `json:"id"`
	Data string `json:"data"`
}

type ApiRepository interface {
	Get() (*[]Api, error)
	Create(data string) (*Api, error)
}

type ApiService interface {
	ProcessData(data string) (string, error)
	GetCurrency(currency string, ch chan<- Result, wg *sync.WaitGroup) ([]Currency, error)
	GetCurrencies(currency ...string) ([]Currency, error)
}

type Currency struct {
	Id          uint64
	Code        string `json:"code,omitempty"`
	Codein      string `json:"codein,omitempty"`
	Name        string `json:"name,omitempty"`
	High        string `json:"high,omitempty"`
	Low         string `json:"low,omitempty"`
	PctChange   string `json:"pctChange,omitempty"`
	Bid         string `json:"bid,omitempty"`
	Ask         string `json:"ask,omitempty"`
	VarBid      string `json:"varBid,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	Create_data string `json:"create_date,omitempty"`
	Partial     bool   `json:"partial"`
}

type Result struct {
	Data []Currency
	Err  error
}
