package handlers

import (
	"advancedbank/internal/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBalance(c *gin.Context) {
	email := c.Param("email")

	account, err := h.service.GetAccount(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "opened account not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": account.GetBalance()})
}

func (h *Handler) AmountOperation(c *gin.Context) {
	email := c.Param("email")
	account, customErr := h.service.GetAccount(email)
	if customErr != nil {
		c.JSON(customErr.Status(), gin.H{"error": customErr.Error()})
		return
	}

	var req models.AmountOperationsRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validateOperation(req.Operation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customErr = h.service.AmountOperation(req.Operation, req.Amount, account)
	if customErr != nil {
		c.JSON(customErr.Status(), gin.H{"error": customErr.Error()})
		return
	}
}

func validateOperation(operation string) error {
	if operation != "withdraw" && operation != "deposit" {
		return errors.New("invalid operation")
	}

	return nil
}
