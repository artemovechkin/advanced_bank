package handlers

import (
	"advancedbank/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Transfer(c *gin.Context) {
	var req models.TransferRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	senderAccount, _ := h.store.GetAccount(req.EmailFrom)
	//if !exists {
	//	c.JSON(http.StatusNotFound, gin.H{"error": "sender account not found"})
	//	return
	//}

	receiverAccount, _ := h.store.GetAccount(req.EmailTo)
	//if !exists {
	//	c.JSON(http.StatusNotFound, gin.H{"error": "receiver account not found"})
	//	return
	//}

	err := senderAccount.Transfer(req.Amount, &receiverAccount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"message": fmt.Sprintf("transferred %.2f from %s to %s", req.Amount, req.EmailFrom, req.EmailTo)},
	)

}
