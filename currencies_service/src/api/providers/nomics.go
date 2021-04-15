package providers

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/yuricampolongo/crypto-monitoring/clients/restclient"
	"github.com/yuricampolongo/crypto-monitoring/currencies_service/src/api/domain"
)

const (
	apiKey                = "1b1a19cbba9fe8adce2d3e34dc5a5fd3"
	sucessfullResponse    = 200
	endpointGetCurrencies = "https://api.nomics.com"
	pathGetCurrencies     = "/v1/currencies/ticker"
)

var (
	Currencies CurrenciesInterface
)

type CurrenciesInterface interface {
	Get(cryptoCurrencyRequest domain.CurrencyRequest) (*[]domain.CurrencyResponse, error)
}

type currencies struct{}

func init() {
	Currencies = &currencies{}
}

func (c *currencies) Get(cryptoCurrencyRequest domain.CurrencyRequest) (*[]domain.CurrencyResponse, error) {
	params := map[string]string{
		"key":      apiKey,
		"ids":      cryptoCurrencyRequest.Ids,
		"convert":  cryptoCurrencyRequest.Convert,
		"interval": cryptoCurrencyRequest.Interval,
	}

	response, err := restclient.Do.Get(endpointGetCurrencies, pathGetCurrencies, params)
	if err != nil {
		return nil, errors.New("error to get currencies from API")
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("invalid response body")
	}
	defer response.Body.Close()

	if response.StatusCode != sucessfullResponse {
		return nil, errors.New("nomics API error")
	}

	var result []domain.CurrencyResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, errors.New("error to unmarshal response body")
	}

	return &result, nil
}
