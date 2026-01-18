package handlers

import (
	"advancedbank/internal/customerror"
	"advancedbank/internal/models"
	"advancedbank/internal/service"

	"github.com/gin-gonic/gin"
)

type IService interface {
	CreateAccount(req models.CreateAccountRequest) customerror.Error
	CloseAccount(email string) customerror.Error
	GetAccount(email string) (models.BankAccount, customerror.Error)

	AmountOperation(operation string, amount float64, account models.BankAccount) customerror.Error
	Transfer(req models.TransferRequest) customerror.Error
}

type Handler struct {
	store service.IStorage // todo временное поле

	service IService
}

func New(service IService) *Handler {
	return &Handler{
		service: service,
	}
}

func Init(r *gin.Engine, h *Handler) {
	r.POST("/account/create", h.CreateAccount)
	r.POST("/account/close/:email", h.CloseAccount)

	r.GET("/balance/:email", h.GetBalance)
	r.POST("/amount/:email", h.AmountOperation)

	r.POST("/transfer", h.Transfer)
}
