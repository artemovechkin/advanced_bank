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

	customErr := h.service.Transfer(req)
	if customErr != nil {
		c.JSON(customErr.Status(), gin.H{"error": customErr.Error()})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"message": fmt.Sprintf("transferred %.2f from %s to %s", req.Amount, req.EmailFrom, req.EmailTo)},
	)
}
