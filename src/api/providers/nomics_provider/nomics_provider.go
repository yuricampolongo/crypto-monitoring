package nomics_provider

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/yuricampolongo/crypto-monitoring/src/api/clients/restclient"
	"github.com/yuricampolongo/crypto-monitoring/src/api/domain/nomics"
)

const (
	apiKey                = "1b1a19cbba9fe8adce2d3e34dc5a5fd3"
	sucessfullResponse    = 200
	endpointGetCurrencies = "https://api.nomics.com"
	pathGetCurrencies     = "/v1/currencies/ticker"
)

func GetCurrencies(cryptoCurrencyRequest nomics.CurrencyTickerRequest) (*[]nomics.CurrencyTickerResponse, error) {
	params := map[string]string{
		"key":      apiKey,
		"ids":      cryptoCurrencyRequest.Ids,
		"convert":  cryptoCurrencyRequest.Convert,
		"interval": cryptoCurrencyRequest.Interval,
	}

	response, err := restclient.Get(endpointGetCurrencies, pathGetCurrencies, params)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != sucessfullResponse {
		return nil, errors.New("Nomics API error")
	}

	var result []nomics.CurrencyTickerResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, errors.New("Error to unmarshal response body")
	}

	return &result, nil
}
