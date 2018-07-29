package web

import (
	"net/http"

	"github.com/athom/grouper"
	"github.com/gin-gonic/gin"
)

// Binding from JSON
//type CommonFriendsRequest struct {
//	Friends []string `json:"friends" binding:"required"`
//}

//type CommonFriendsOutput struct {
//	Success bool     `json:"success"`
//	Friends []string `json:"friends"`
//	Count   int      `json:"count"`
//}

type CommonFriendsRequest ConnectRequest

type CommonFriendsOutput FindFriendsOutput

func (this *Controller) commonFriendsController(c *gin.Context) {
	var input CommonFriendsRequest
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
	ids, err := this.core.CommonFriends(id1, id2)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var emails = make([]string, 0)
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
