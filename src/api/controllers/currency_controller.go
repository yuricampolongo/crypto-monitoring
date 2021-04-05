package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuricampolongo/crypto-monitoring/src/api/domain/nomics"
	"github.com/yuricampolongo/crypto-monitoring/src/api/providers/nomics_provider"
)

func GetCryptoCurrencies(c *gin.Context) {
	response, err := nomics_provider.GetCurrencies(nomics.CurrencyTickerRequest{
		Ids:      c.Param("ids"),
		Convert:  c.Param("convert"),
		Interval: c.Param("interval"),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, "An error occurred to get currencies")
	}

	c.JSON(http.StatusOK, response)
}
