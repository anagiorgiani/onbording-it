package controller

import (
	"sync"

	"github.com/anagiorgiani/onbording-it/internal/my_api/domain"
)

type mockCurrencyService struct {
	result interface{}
	err    error
}

func (m mockCurrencyService) GetCurrency(currency string, ch chan<- domain.Result, wg *sync.WaitGroup) ([]domain.Currency, error) {
	if m.err != nil {
		return []domain.Currency{}, m.err
	}

	return m.result.([]domain.Currency), nil
}

func (m mockCurrencyService) GetCurrencies(currency ...string) ([]domain.Currency, error) {
	if m.err != nil {
		return []domain.Currency{}, m.err
	}

	return m.result.([]domain.Currency), nil
}

func (m mockCurrencyService) ProcessData(data string) (string, error) {
	if m.err != nil {
		return "", m.err
	}
	return data, nil
}
