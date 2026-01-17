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

	senderAccount, err := h.store.GetAccount(req.EmailFrom)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "opened account not found"})
		return
	}

	receiverAccount, err := h.store.GetAccount(req.EmailTo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "opened account not found"})
		return
	}

	err = senderAccount.Transfer(req.Amount, &receiverAccount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.store.UpdateAccount(receiverAccount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.store.UpdateAccount(senderAccount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"message": fmt.Sprintf("transferred %.2f from %s to %s", req.Amount, req.EmailFrom, req.EmailTo)},
	)
}
