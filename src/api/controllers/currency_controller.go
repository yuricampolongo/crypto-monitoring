package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuricampolongo/crypto-monitoring/src/api/domain/nomics"
	"github.com/yuricampolongo/crypto-monitoring/src/api/providers/nomics_provider"
)

func GetCryptoCurrencies(c *gin.Context) {
	nomics_provider.GetCurrencies(nomics.CurrencyTickerRequest{
		Ids:      c.Query("ids"),
		Convert:  c.Query("convert"),
		Interval: c.Query("interval"),
	})
}
