package api

import "github.com/yuricampolongo/crypto-monitoring/src/api/controllers"

func mapUrls() {
	router.GET("/crypto/currency/:ids/:convert/:interval", controllers.GetCryptoCurrencies)
}
