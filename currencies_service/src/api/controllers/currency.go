package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuricampolongo/crypto-monitoring/currencies_service/src/api/domain"
	"github.com/yuricampolongo/crypto-monitoring/currencies_service/src/api/providers"
)

func GetCryptoCurrencies(c *gin.Context) {
	response, err := providers.Currencies.Get(domain.CurrencyRequest{
		Ids:      c.Param("ids"),
		Convert:  c.Param("convert"),
		Interval: c.Param("interval"),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, "An error occurred to get currencies")
		return
	}

	c.JSON(http.StatusOK, response)
}
