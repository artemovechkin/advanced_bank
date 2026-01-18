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

func (s *Service) GetAccount(email string) (models.BankAccount, customerror.Error) {
	account, err := s.store.GetAccount(email)
	if err != nil {
		return models.BankAccount{}, &customerror.CustomError{
			State:   http.StatusNotFound,
			Message: fmt.Sprintf("failed to get account: %v", err.Error()),
		}
	}

	return account, nil
}

func (s *Service) AmountOperation(operation string, amount float64, account models.BankAccount) customerror.Error {
	switch operation {
	case "withdraw":
		err := account.Withdraw(amount)
		if err != nil {
			return &customerror.CustomError{
				State:   http.StatusBadRequest,
				Message: fmt.Sprintf("failed to withdraw: %v", err.Error()),
			}
		}
	case "deposit":
		err := account.Deposit(amount)
		if err != nil {
			return &customerror.CustomError{
				State:   http.StatusBadRequest,
				Message: fmt.Sprintf("failed to deposit: %v", err.Error()),
			}
		}
	}

	err := s.store.UpdateAccount(account)
	if err != nil {
		return &customerror.CustomError{
			State:   http.StatusInternalServerError,
			Message: fmt.Sprintf("failed to update account: %v", err.Error()),
		}
	}

	return nil
}

func (s *Service) Transfer(req models.TransferRequest) customerror.Error {
	senderAccount, customErr := s.store.GetAccount(req.EmailFrom)
	if customErr != nil {
		return &customerror.CustomError{
			State:   http.StatusNotFound,
			Message: fmt.Sprintf("opened account not found: %v", customErr.Error()),
		}
	}

	receiverAccount, customErr := s.store.GetAccount(req.EmailTo)
	if customErr != nil {
		return &customerror.CustomError{
			State:   http.StatusNotFound,
			Message: fmt.Sprintf("opened account not found: %v", customErr.Error()),
		}
	}

	err := senderAccount.Transfer(req.Amount, &receiverAccount)
	if err != nil {
		return &customerror.CustomError{
			State:   http.StatusBadRequest,
			Message: fmt.Sprintf("failed to transfer: %v", err.Error()),
		}
	}

	err = s.store.UpdateAccount(receiverAccount)
	if err != nil {
		return &customerror.CustomError{
			State:   http.StatusInternalServerError,
			Message: fmt.Sprintf("failed to update account: %v", err.Error()),
		}
	}

	err = s.store.UpdateAccount(senderAccount)
	if err != nil {
		return &customerror.CustomError{
			State:   http.StatusInternalServerError,
			Message: fmt.Sprintf("failed to update account: %v", err.Error()),
		}
	}

	return nil
}
