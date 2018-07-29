package web

import (
	"net/http"

	"github.com/athom/grouper"
	"github.com/gin-gonic/gin"
)

// Binding from JSON
type GetRecipientsRequest struct {
	Sender string `json:"sender" binding:"required"`
	Text   string `json:"text"`
}

type GetRecipientsOutput struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients" binding:"required"`
}

func (this *Controller) getRecipientsController(c *gin.Context) {
	var input GetRecipientsRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id1 := grouper.FriendId(input.Sender)
	ids, err := this.core.Receipients(id1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	var recipients = make([]string, 0)
	for _, id := range ids {
		recipients = append(recipients, id.String())
	}
	output := &GetRecipientsOutput{
		Success:    true,
		Recipients: recipients,
	}
	c.JSON(http.StatusOK, output)
}
