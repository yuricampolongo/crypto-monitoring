package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuricampolongo/crypto-monitoring/user_info_service/src/api/domain"
	"github.com/yuricampolongo/crypto-monitoring/user_info_service/src/api/providers"
)

func AddUser(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "invalid request body")
	}
	var user *domain.User
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error to unmarshal json body request")
	}

	err = providers.DynamoDB.Save((*user))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error to save user")
	}

	fmt.Println(user)
}
