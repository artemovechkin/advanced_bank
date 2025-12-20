package handlers

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	store IStorage
}

func New(store IStorage) *Handler {
	return &Handler{store: store}
}

func Init(r *gin.Engine, h *Handler) {
	r.POST("/account/create", h.CreateAccount)
	r.POST("/account/close/:email", h.CloseAccount)

	r.GET("/balance/:email", h.GetBalance)
	r.POST("/amount/:email", h.AmountOperations)

	r.POST("/transfer", h.Transfer)
}
