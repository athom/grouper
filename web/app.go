package web

import (
	"encoding/json"
	"fmt"

	"strings"

	"github.com/DisposaBoy/JsonConfigReader"
	"github.com/gin-gonic/gin"
)

type MySqlSettings struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type config struct {
	Port          int           `json:"port"`
	StorageType   string        `json:"storage_type"`
	MySqlSettings MySqlSettings `json:"mysql_settings"`
}

func route(controller *Controller) (r *gin.Engine) {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		gFriends := v1.Group("/friends")

		gFriends.POST("/connect", controller.connectController)
		gFriends.POST("/subscribe", controller.subscribeController)
		gFriends.POST("/block", controller.blockController)

		gFriends.POST("/find", controller.getFriendsController)
		gFriends.POST("/common", controller.commonFriendsController)
		gFriends.POST("/recipients", controller.getRecipientsController)
	}
	return router
}

func Run(conf string) {
	jsonReader := JsonConfigReader.New(strings.NewReader(conf))
	config := &config{}
	err := json.NewDecoder(jsonReader).Decode(config)
	if err != nil {
		panic(err)
	}

	controller := &Controller{}
	controller.Setup(config)

	r := route(controller)

	port := fmt.Sprintf(":%v", config.Port)
	r.Run(port)
}
