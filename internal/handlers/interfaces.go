package handlers

import "advancedbank/internal/models"

type IStorage interface {
	LoadAccounts() error
	SaveAccounts()
	GetAccount(email string) (*models.BankAccount, bool)
	SetAccount(email string, b *models.BankAccount)
}
