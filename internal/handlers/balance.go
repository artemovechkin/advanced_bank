package handlers

import (
	"advancedbank/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBalance(c *gin.Context) {
	email := c.Param("email")

	account, err := h.store.GetAccount(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "opened account not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": account.GetBalance()})
}

func (h *Handler) AmountOperations(c *gin.Context) {
	email := c.Param("email")

	account, err := h.store.GetAccount(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "opened account not found"})
		return
	}

	var req models.AmountOperationsRequest

	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch req.Operation {
	case "withdraw":
		err = account.Withdraw(req.Amount)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"status": "successfully withdrawn"})

	case "deposit":
		err = account.Deposit(req.Amount)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"status": "successfully deposited"})

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid operation"})
	}
}
