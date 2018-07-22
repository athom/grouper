package web

import "github.com/gin-gonic/gin"

func route() (r *gin.Engine) {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		gFriends := v1.Group("/friends")

		gFriends.POST("/connect", makeFriendController)
		gFriends.POST("/subscribe", makeFriendController)
		gFriends.POST("/find", getFriendsController)

	}
	return router
}

func Run() {
	r := route()
	r.Run(":7200")
}
