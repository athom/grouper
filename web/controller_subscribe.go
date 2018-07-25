package web

import (
	"net/http"

	"github.com/athom/grouper"
	"github.com/gin-gonic/gin"
)

// Binding from JSON
type SubscribeRequest struct {
	Requestor string `json:"requestor" binding:"required"`
	Target    string `json:"target" binding:"required"`
}

func (this *Controller) subscribeController(c *gin.Context) {
	var input SubscribeRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id1 := grouper.FriendId(input.Requestor)
	id2 := grouper.FriendId(input.Target)
	err := this.core.Subscribe(id1, id2)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ok)
}
