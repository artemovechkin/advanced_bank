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

	alterTableAccounts := `
		ALTER TABLE accounts ADD COLUMN is_active BOOLEAN DEFAULT true;
`
	// todo переписать на миграции
	db.Exec(alterTableAccounts)

	return db
}

func (s *Storage) GetAccount(email string) (account models.BankAccount, err error) {
	err = s.db.QueryRow(`SELECT * FROM accounts WHERE email = ? AND is_active = true`, email).Scan(
		&account.Owner.Email, &account.Owner.Name, &account.Owner.Age, &account.Balance, &account.IsActive)
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

func (s *Storage) UpdateAccount(account models.BankAccount) error {
	updateQuery := `UPDATE accounts SET name = ?, age = ?, balance = ?, is_active = ? WHERE email = ?;`

	_, err := s.db.Exec(updateQuery, account.Owner.Name, account.Owner.Age, account.Balance, account.IsActive, account.Owner.Email)
	if err != nil {
		return err
	}

	return nil
}
