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

	customErr := h.service.CreateAccount(req)
	if customErr != nil {
		c.JSON(customErr.Status(), gin.H{"error": customErr.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) CloseAccount(c *gin.Context) {
	email := c.Param("email")

	err := h.service.CloseAccount(email)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
