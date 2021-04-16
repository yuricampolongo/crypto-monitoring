package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuricampolongo/crypto-monitoring/user_info_service/src/api/domain"
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

	fmt.Println(user)
}
