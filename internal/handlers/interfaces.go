package handlers

import (
	"advancedbank/internal/models"

	"modernc.org/sqlite"
)

type IStorage interface {
	GetAccount(email string) (models.BankAccount, error)
	SetAccount(account models.BankAccount) *sqlite.Error
	UpdateAccount(account models.BankAccount) error
}
