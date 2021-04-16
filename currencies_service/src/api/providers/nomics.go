package providers

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/yuricampolongo/crypto-monitoring/clients/restclient"
	"github.com/yuricampolongo/crypto-monitoring/currencies_service/src/api/domain"
)

const (
	sucessfullResponse    = 200
	endpointGetCurrencies = "https://api.nomics.com"
	pathGetCurrencies     = "/v1/currencies/ticker"
)

var (
	apiKey     *string
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
	checkApiKey()

	params := map[string]string{
		"key":      (*apiKey),
		"ids":      cryptoCurrencyRequest.Ids,
		"convert":  cryptoCurrencyRequest.Convert,
		"interval": cryptoCurrencyRequest.Interval,
	}

	response, err := restclient.Do.Get(endpointGetCurrencies, pathGetCurrencies, params)
	if err != nil {
		return nil, errors.New("error to get currencies from API")
	}

	if response.StatusCode != sucessfullResponse {
		return nil, errors.New("nomics API error")
	}

	var result []domain.CurrencyResponse
	if err := json.Unmarshal([]byte(response.Body), &result); err != nil {
		return nil, errors.New("error to unmarshal response body")
	}

	return &result, nil
}

func checkApiKey() {
	value, present := os.LookupEnv("NOMICS_API_KEY")
	if !present {
		log.Fatal("no Nomics API Key set. Please set the key in the environment variable $NOMICS_API_KEY")
		apiKey = nil
		return
	}
	apiKey = &value
}
