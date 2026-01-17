package service

import (
	"advancedbank/internal/customerror"
	"advancedbank/internal/models"
	"fmt"
	"net/http"

	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

type IStorage interface {
	GetAccount(email string) (models.BankAccount, error)
	SetAccount(account models.BankAccount) *sqlite.Error
	UpdateAccount(account models.BankAccount) error
}

type Service struct {
	store IStorage
}

func New(store IStorage) *Service {
	return &Service{store: store}
}

func (s *Service) CreateAccount(req models.CreateAccountRequest) customerror.Error {
	account := models.NewBankAccount(models.AccountOwner{
		Name:  req.Name,
		Age:   req.Age,
		Email: req.Email,
	}, req.InitialBalance)

	sqliteErr := s.store.SetAccount(*account)
	if sqliteErr != nil {
		if sqliteErr.Code() == sqlite3.SQLITE_CONSTRAINT_PRIMARYKEY {
			return &customerror.CustomError{
				State:   http.StatusBadRequest,
				Message: fmt.Sprintf("account already exists: %v", sqliteErr.Error()),
			}
		}

		return &customerror.CustomError{
			State:   http.StatusInternalServerError,
			Message: fmt.Sprintf("failed to set account: %v", sqliteErr.Error()),
		}
	}

	return nil
}

func (s *Service) CloseAccount(email string) customerror.Error {
	account, err := s.store.GetAccount(email)
	if err != nil {
		return &customerror.CustomError{
			State:   http.StatusNotFound,
			Message: fmt.Sprintf("opened account not found: %v", err.Error()),
		}
	}

	err = account.CloseAccount()
	if err != nil {
		return &customerror.CustomError{
			State:   http.StatusBadRequest,
			Message: fmt.Sprintf("bad request: %v", err.Error()),
		}
	}

	err = s.store.UpdateAccount(account)
	if err != nil {
		return &customerror.CustomError{
			State:   http.StatusInternalServerError,
			Message: fmt.Sprintf("failed to update account: %v", err.Error()),
		}
	}

	return nil
}
