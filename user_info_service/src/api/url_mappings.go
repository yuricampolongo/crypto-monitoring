package api

import "github.com/yuricampolongo/crypto-monitoring/user_info_service/src/api/controllers"

func mapUrls() {
	router.POST("/user/add", controllers.AddUser)
	router.POST("/user/monitor/crypto", controllers.MonitorCrypto)
}
