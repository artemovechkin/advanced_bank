package handlers

import (
	"advancedbank/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateAccount(c *gin.Context) {
	var req models.CreateAccountRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, exists := h.store.GetAccount(req.Email)
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account already exists"})
		return
	}

	account := models.NewBankAccount(models.AccountOwner{
		Name:  req.Name,
		Age:   req.Age,
		Email: req.Email,
	}, req.InitialBalance)

	h.store.SetAccount(account.Owner.Email, account)

	c.Status(http.StatusCreated)
}

func (h *Handler) CloseAccount(c *gin.Context) {
	email := c.Param("email")

	account, exists := h.store.GetAccount(email)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	err := account.CloseAccount()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
