package storage

import (
	"advancedbank/internal/models"
	"encoding/json"
	"log/slog"
	"os"
)

type Storage struct {
	accounts map[string]*models.BankAccount
}

func New() *Storage {
	return &Storage{make(map[string]*models.BankAccount)}
}

func (s *Storage) LoadAccounts() error {
	bytes, err := os.ReadFile("accounts.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &s.accounts)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) SaveAccounts() {
	bytes, err := json.MarshalIndent(s.accounts, "", "  ")
	if err != nil {
		slog.Error("failed to marshal accounts", "error", err)
		return
	}

	err = os.WriteFile("accounts.json", bytes, os.ModePerm)
	if err != nil {
		slog.Error("failed to write file", "error", err)
		return
	}

	slog.Info("saved accounts to file success")
}

func (s *Storage) GetAccount(email string) (*models.BankAccount, bool) {
	account, exists := s.accounts[email]
	return account, exists
}

func (s *Storage) SetAccount(email string, account *models.BankAccount) {
	s.accounts[email] = account
}
