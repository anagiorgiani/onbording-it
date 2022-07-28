package controller

import (
	"net/http/httptest"

	"net/http"
	"testing"

	"github.com/anagiorgiani/onbording-it/internal/my_api/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Get(t *testing.T) {

	currencyList := []domain.Currency{
		{
			Id:          1,
			Code:        "AAAA",
			Codein:      "AAAA",
			Name:        "AAAA",
			High:        "AAAA",
			Low:         "AAAA",
			PctChange:   "AAAA",
			Bid:         "AAAA",
			Ask:         "AAAA",
			VarBid:      "AAAA",
			Timestamp:   "AAAA",
			Create_data: "AAAA",
			Partial:     true,
		},
	}

	mockService := mockCurrencyService{
		result: currencyList,
		err:    nil,
	}

	router := setupRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/myApi", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
}

func setupRouter(mockService mockCurrencyService) *gin.Engine {
	controller := NewController(mockService)

	router := gin.Default()
	router.GET("/api/myApi", controller.Get())

	return router
}
