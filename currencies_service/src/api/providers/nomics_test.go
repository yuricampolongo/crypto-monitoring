package providers

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yuricampolongo/crypto-monitoring/clients/restclient"
	"github.com/yuricampolongo/crypto-monitoring/currencies_service/src/api/domain"
)

func TestCurrenciesGetErrorFromAPI(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRest := restclient.NewMockRestInterface(mockCtrl)
	restclient.Do = mockRest

	params := map[string]string{
		"key":      "",
		"ids":      "BTC",
		"convert":  "BRL",
		"interval": "1h",
	}

	mockRest.EXPECT().Get("https://api.nomics.com", "/v1/currencies/ticker", params).Return(nil, errors.New("")).Times(1)

	resp, err := Currencies.Get(domain.CurrencyRequest{
		Ids:      params["ids"],
		Convert:  params["convert"],
		Interval: params["interval"],
	})

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, "error to get currencies from API", err.Error())
}

func TestCurrenciesApiError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRest := restclient.NewMockRestInterface(mockCtrl)
	restclient.Do = mockRest

	params := map[string]string{
		"key":      "",
		"ids":      "BTC",
		"convert":  "BRL",
		"interval": "1h",
	}

	mockedResponse := &restclient.Response{
		StatusCode: http.StatusBadRequest,
		Body:       "error from API",
	}

	mockRest.EXPECT().Get("https://api.nomics.com", "/v1/currencies/ticker", params).Return(mockedResponse, nil).Times(1)

	resp, err := Currencies.Get(domain.CurrencyRequest{
		Ids:      params["ids"],
		Convert:  params["convert"],
		Interval: params["interval"],
	})

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, "nomics API error", err.Error())
}

func TestCurrenciesInvalidResponseBody(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRest := restclient.NewMockRestInterface(mockCtrl)
	restclient.Do = mockRest

	params := map[string]string{
		"key":      "apiKey",
		"ids":      "BTC",
		"convert":  "BRL",
		"interval": "1h",
	}

	mockedResponse := &restclient.Response{
		StatusCode: http.StatusOK,
		Body:       "{",
	}

	mockRest.EXPECT().Get("https://api.nomics.com", "/v1/currencies/ticker", params).Return(mockedResponse, nil).Times(1)

	resp, err := Currencies.Get(domain.CurrencyRequest{
		Ids:      params["ids"],
		Convert:  params["convert"],
		Interval: params["interval"],
	})

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, "error to unmarshal response body", err.Error())
}

func TestCurrencie(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRest := restclient.NewMockRestInterface(mockCtrl)
	restclient.Do = mockRest

	params := map[string]string{
		"key":      "apiKey",
		"ids":      "BTC",
		"convert":  "BRL",
		"interval": "1h",
	}

	mockedCurrencyResponse := []domain.CurrencyResponse{
		{
			Id:        "BTC",
			Currency:  "BRL",
			Name:      "Bitcoin",
			Logo:      "empty.jpg",
			Price:     "72000",
			PriceDate: "",
		},
	}

	jsonString, _ := json.Marshal(mockedCurrencyResponse)

	mockedResponse := &restclient.Response{
		StatusCode: http.StatusOK,
		Body:       string(jsonString),
	}

	mockRest.EXPECT().Get("https://api.nomics.com", "/v1/currencies/ticker", params).Return(mockedResponse, nil).Times(1)

	resp, err := Currencies.Get(domain.CurrencyRequest{
		Ids:      params["ids"],
		Convert:  params["convert"],
		Interval: params["interval"],
	})

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, mockedCurrencyResponse[0].Id, (*resp)[0].Id)
}
