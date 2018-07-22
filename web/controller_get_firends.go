package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Binding from JSON
type FindFriendsRequest struct {
	Email []string `json:"email" binding:"required"`
}

func getFriendsController(c *gin.Context) {
	var input FindFriendsRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ok)
}
