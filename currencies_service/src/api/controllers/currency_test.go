package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yuricampolongo/crypto-monitoring/currencies_service/src/api/domain"
	"github.com/yuricampolongo/crypto-monitoring/currencies_service/src/api/providers"
)

func TestGetCryptoCurrenciesError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockCurrencies := providers.NewMockCurrenciesInterface(mockCtrl)
	providers.Currencies = mockCurrencies

	request := domain.CurrencyRequest{
		Ids:      "BTC",
		Convert:  "BRL",
		Interval: "1h",
	}

	mockCurrencies.EXPECT().Get(request).Return(nil, errors.New("invalid request")).Times(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "ids", Value: request.Ids},
		{Key: "convert", Value: request.Convert},
		{Key: "interval", Value: request.Interval},
	}

	GetCryptoCurrencies(c)

	assert.EqualValues(t, http.StatusInternalServerError, w.Code)
	var result string
	json.NewDecoder(w.Body).Decode(&result)
	assert.EqualValues(t, "An error occurred to get currencies", result)
}

func TestGetCryptoCurrencies(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockCurrencies := providers.NewMockCurrenciesInterface(mockCtrl)
	providers.Currencies = mockCurrencies

	request := domain.CurrencyRequest{
		Ids:      "BTC",
		Convert:  "BRL",
		Interval: "1h",
	}

	mockedResponse := []domain.CurrencyResponse{
		{
			Id:        "BTC",
			Currency:  "BRL",
			Name:      "Bitcoin",
			Logo:      "empty.jpg",
			Price:     "72000",
			PriceDate: "",
		},
	}

	mockCurrencies.EXPECT().Get(request).Return(&mockedResponse, nil).Times(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "ids", Value: request.Ids},
		{Key: "convert", Value: request.Convert},
		{Key: "interval", Value: request.Interval},
	}

	GetCryptoCurrencies(c)

	assert.EqualValues(t, http.StatusOK, w.Code)
	var result []domain.CurrencyResponse
	json.NewDecoder(w.Body).Decode(&result)
	assert.EqualValues(t, mockedResponse[0].Id, result[0].Id)
	assert.EqualValues(t, mockedResponse[0].Currency, result[0].Currency)
	assert.EqualValues(t, mockedResponse[0].Logo, result[0].Logo)
	assert.EqualValues(t, mockedResponse[0].Name, result[0].Name)
	assert.EqualValues(t, mockedResponse[0].Price, result[0].Price)
	assert.EqualValues(t, mockedResponse[0].PriceDate, result[0].PriceDate)
}
