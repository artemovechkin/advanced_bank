package models

import (
	"errors"
)

type AccountOwner struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type BankAccount struct {
	Owner   AccountOwner `json:"owner"`
	Balance float64      `json:"balance"`
	Closed  bool         `json:"closed"`
}

func NewBankAccount(owner AccountOwner, initialBalance float64) *BankAccount {
	return &BankAccount{
		Owner:   owner,
		Balance: initialBalance,
		Closed:  false,
	}
}

func (ba *BankAccount) GetBalance() float64 {
	return ba.Balance
}

func (ba *BankAccount) Deposit(amount float64) error {
	if ba.Closed {
		return errors.New("account is closed")
	}
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	ba.Balance += amount
	return nil
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if ba.Closed {
		return errors.New("account is closed")
	}
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if ba.Balance < amount {
		return errors.New("insufficient funds")
	}
	ba.Balance -= amount
	return nil
}

func (ba *BankAccount) Transfer(amount float64, receiver *BankAccount) error {
	if ba.Closed {
		return errors.New("sender account is closed")
	}
	if receiver.Closed {
		return errors.New("receiver account is closed")
	}
	if amount <= 0 {
		return errors.New("transfer amount must be positive")
	}
	if ba.Balance < amount {
		return errors.New("insufficient funds")
	}
	ba.Balance -= amount
	receiver.Balance += amount
	return nil
}

func (ba *BankAccount) CloseAccount() error {
	if ba.Closed {
		return errors.New("account is already closed")
	}
	if ba.Balance != 0 {
		return errors.New("cannot close account with non-zero balance")
	}
	ba.Closed = true
	return nil
}

type CreateAccountRequest struct {
	Name           string  `json:"name" binding:"required"`
	Age            int     `json:"age" binding:"required"`
	Email          string  `json:"email" binding:"required"`
	InitialBalance float64 `json:"initial_balance"`
}

type AmountOperationsRequest struct {
	Operation string  `json:"operation" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
}

type TransferRequest struct {
	EmailFrom string  `json:"email_from" binding:"required"`
	EmailTo   string  `json:"email_to" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
}
