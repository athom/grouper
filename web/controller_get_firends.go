package web

import (
	"net/http"

	"github.com/athom/grouper"
	"github.com/gin-gonic/gin"
)

// Binding from JSON
type FindFriendsRequest struct {
	Email string `json:"email" binding:"required"`
}

type FindFriendsOutput struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

func (this *Controller) getFriendsController(c *gin.Context) {
	var input FindFriendsRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := grouper.FriendId(input.Email)
	ids, err := this.core.ListFriends(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var emails []string
	for _, id := range ids {
		emails = append(emails, id.String())
	}

	output := &FindFriendsOutput{
		Success: true,
		Count:   len(emails),
		Friends: emails,
	}

	c.JSON(http.StatusOK, output)
}
