package web

import (
	"net/http"

	"github.com/athom/grouper"
	"github.com/gin-gonic/gin"
)

// Binding from JSON
type ConnectRequest struct {
	Friends []string `json:"friends" binding:"required"`
}

func (this *Controller) connectController(c *gin.Context) {
	var input ConnectRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(input.Friends) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "length of friends must be 2"})
		return
	}

	id1 := grouper.FriendId(input.Friends[0])
	id2 := grouper.FriendId(input.Friends[1])
	err := this.core.MakeFriend(id1, id2)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ok)
}
