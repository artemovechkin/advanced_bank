package storage

import (
	"advancedbank/internal/models"
	"advancedbank/migrations"
	"database/sql"

	"github.com/pressly/goose/v3"
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
	db, err := sql.Open("sqlite", "bank_new.db")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	goose.SetBaseFS(migrations.EmbedMigrations)

	if err = goose.SetDialect("sqlite"); err != nil {
		panic(err)
	}

	err = goose.Up(db, ".")
	if err != nil {
		panic(err)
	}

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
