package storage

import (
	"advancedbank/internal/models"
	"database/sql"

	"modernc.org/sqlite"
	_ "modernc.org/sqlite"
)

type Storage struct {
	db *sql.DB
}

func New() *Storage {
	return &Storage{
		db: OpenConnection(),
	}
}

func OpenConnection() *sql.DB {
	db, err := sql.Open("sqlite", "bank.db")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	createTableAccounts := `
	CREATE TABLE IF NOT EXISTS accounts (
 		email TEXT PRIMARY KEY,
  		name TEXT,
  		age int,
  		balance float
);`

	_, err = db.Exec(createTableAccounts)
	if err != nil {
		panic(err)
	}

	return db
}

func (s *Storage) GetAccount(email string) (account models.BankAccount, err error) {
	err = s.db.QueryRow(`SELECT * FROM accounts WHERE email = ?`, email).Scan(
		&account.Owner.Email, &account.Owner.Name, &account.Owner.Age, &account.Balance)
	if err != nil {
		return models.BankAccount{}, err
	}

	return account, nil
}

func (s *Storage) SetAccount(account models.BankAccount) *sqlite.Error {
	queryInsert := `INSERT INTO accounts (email, name, age, balance) VALUES (?, ?, ?, ?);`

	_, err := s.db.Exec(queryInsert, account.Owner.Email, account.Owner.Name, account.Owner.Age, account.Balance)
	if err != nil {
		return err.(*sqlite.Error)
	}

	return nil
}
