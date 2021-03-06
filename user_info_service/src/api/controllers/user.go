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
	userTable = "tUser"
)

func init() {
	db.Provider().CreateTable(userTable)
}

func AddUser(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "invalid request body")
	}
	var user domain.User
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error to unmarshal json body request")
	}

	err = db.Provider().Save(user, userTable)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error to save user")
	}
}
