package providers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yuricampolongo/crypto-monitoring/clients/restclient"
	"github.com/yuricampolongo/crypto-monitoring/currencies_service/src/api/domain"
)

func TestCurrenciesGetErrorFromAPI(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRest := restclient.NewMockRestInterface(mockCtrl)
	restclient.Do = mockRest

	params := map[string]string{
		"key":      apiKey,
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

func TestCurrenciesReadBodyError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRest := restclient.NewMockRestInterface(mockCtrl)
	restclient.Do = mockRest

	params := map[string]string{
		"key":      apiKey,
		"ids":      "BTC",
		"convert":  "BRL",
		"interval": "1h",
	}

	mockReadCloser := mockReadCloser{}
	// if Read is called, it will return error
	mockReadCloser.On("Read", mock.AnythingOfType("[]uint8")).Return(0, fmt.Errorf("error reading"))
	// if Close is called, it will return error
	mockReadCloser.On("Close").Return(fmt.Errorf("error closing"))

	mockedResponse := &http.Response{
		StatusCode: http.StatusOK,
		//Body:       ioutil.NopCloser(strings.NewReader("invalid json response")),
		Body: &mockReadCloser,
	}

	mockRest.EXPECT().Get("https://api.nomics.com", "/v1/currencies/ticker", params).Return(mockedResponse, nil).Times(1)

	resp, err := Currencies.Get(domain.CurrencyRequest{
		Ids:      params["ids"],
		Convert:  params["convert"],
		Interval: params["interval"],
	})

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid response body", err.Error())
}

func TestCurrenciesApiError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRest := restclient.NewMockRestInterface(mockCtrl)
	restclient.Do = mockRest

	params := map[string]string{
		"key":      apiKey,
		"ids":      "BTC",
		"convert":  "BRL",
		"interval": "1h",
	}

	mockedResponse := &http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader("error from API")),
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
		"key":      apiKey,
		"ids":      "BTC",
		"convert":  "BRL",
		"interval": "1h",
	}

	mockedResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(strings.NewReader("{")),
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
		"key":      apiKey,
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

	mockedResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader(jsonString)),
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

type mockReadCloser struct {
	mock.Mock
}

func (m *mockReadCloser) Read(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

func (m *mockReadCloser) Close() error {
	args := m.Called()
	return args.Error(0)
}
