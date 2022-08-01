package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/anagiorgiani/onbording-it/internal/my_api/domain"
)

type apiService struct {
	apirepository domain.ApiRepository
}

func NewService(r domain.ApiRepository) domain.ApiService {
	return &apiService{
		apirepository: r,
	}
}

func (s *apiService) ProcessData(data string) (string, error) {
	_, err := s.apirepository.Create(data)
	if err != nil {
		return "", err
	}
	return data, nil
}

func (s *apiService) GetCurrency(currency string, ch chan<- domain.Result, wg *sync.WaitGroup) ([]domain.Currency, error) {
	getRes := domain.Result{}

	response := []domain.Currency{}
	resp, err := http.Get("https://economia.awesomeapi.com.br/" + currency + "/1")
	if err != nil {
		getRes.Err = err
		getRes.Data = response
		ch <- getRes
		wg.Done()
		return response, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		getRes.Err = err
		getRes.Data = response
		ch <- getRes
		wg.Done()
		return response, err
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		getRes.Err = err
		getRes.Data = response
		ch <- getRes
		wg.Done()
		return response, err
	}

	if len(response) == 0 {
		errresp := domain.Currency{}
		errresp.Partial = true
		response = append(response, errresp)
	}
	getRes.Data = response
	ch <- getRes
	wg.Done()
	return response, nil
}

func (s *apiService) GetCurrencies(currency ...string) ([]domain.Currency, error) {

	var (
		chs = make(chan domain.Result, len(currency))
		wg  sync.WaitGroup
	)
	wg.Add(len(currency))

	currencyList := []domain.Currency{}

	for _, row := range currency {
		defer wg.Done()
		go s.GetCurrency(row, chs, &wg)
	}

	wg.Wait()
	close(chs)
	for ch := range chs {
		if ch.Err != nil {
			return []domain.Currency{}, ch.Err
		}

		currencyList = append(currencyList, ch.Data...)
	}
	return currencyList, nil

}
