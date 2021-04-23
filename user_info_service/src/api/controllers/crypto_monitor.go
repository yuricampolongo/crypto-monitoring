package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuricampolongo/crypto-monitoring/user_info_service/src/api/db"
	"github.com/yuricampolongo/crypto-monitoring/user_info_service/src/api/domain"
)

const (
	userCryptoMonitorTable = "tUser_crypto"
)

func init() {
	db.Provider().CreateTable(userCryptoMonitorTable)
}

func MonitorCrypto(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "invalid request body")
	}
	var cryptos *domain.Crypto
	err = json.Unmarshal(jsonData, &cryptos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error to unmarshal json body request")
	}

	err = db.Provider().Save(cryptos, userCryptoMonitorTable)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error to save crypto monitor information")
	}
}
