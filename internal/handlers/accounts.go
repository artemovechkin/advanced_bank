package handlers

import (
	"advancedbank/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	sqlite3 "modernc.org/sqlite/lib"
)

func (h *Handler) CreateAccount(c *gin.Context) {
	var req models.CreateAccountRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := models.NewBankAccount(models.AccountOwner{
		Name:  req.Name,
		Age:   req.Age,
		Email: req.Email,
	}, req.InitialBalance)

	sqliteErr := h.store.SetAccount(*account)
	if sqliteErr != nil {
		if sqliteErr.Code() == sqlite3.SQLITE_CONSTRAINT_PRIMARYKEY {
			c.JSON(http.StatusBadRequest, gin.H{"error": "account already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": sqliteErr.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) CloseAccount(c *gin.Context) {
	email := c.Param("email")

	account, _ := h.store.GetAccount(email)
	//if !exists {
	//	c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
	//	return
	//}

	err := account.CloseAccount()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
