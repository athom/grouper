package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Binding from JSON
type MakeFriendRequest struct {
	Friends []string `json:"friends" binding:"required"`
}

func makeFriendController(c *gin.Context) {
	var input MakeFriendRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ok)
}
