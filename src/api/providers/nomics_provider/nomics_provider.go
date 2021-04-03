package nomics_provider

import (
	"github.com/yuricampolongo/crypto-monitoring/src/api/clients/restclient"
	"github.com/yuricampolongo/crypto-monitoring/src/api/domain/nomics"
)

const (
	apiKey             = "nomics_api_key_here"
	sucessfullResponse = 200
	urlGetCurrencies   = ""
)

func GetCurrencies(cryptoCurrencyRequest nomics.CurrencyTickerRequest) (*nomics.CurrencyTickerResponse, error) {
	restclient.Get(urlGetCurrencies)
}
