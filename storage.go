package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=goBank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err 
	}
	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}


func (s *PostgresStore) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY, 
		firstName VARCHAR(50),
		lastName VARCHAR(50),
		number serial, 
		balance FLOAT,
		created_at timestamp
		)`
	_, err := s.db.Exec(query)
	return err

}


func (s *PostgresStore) CreateAccount(a *Account) error {
	query := `INSERT INTO account 
		(firstName, lastName, number, balance, created_at) 
		VALUES ($1, $2, $3, $4, $5)`
	resp, err := s.db.Query(query, a.FirstName, a.LastName, a.Number, a.Balance, a.CreatedAt)
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}


func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}


func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}


func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(&account.ID, 
				&account.FirstName, 
				&account.LastName, 
				&account.Number, 
				&account.Balance, 
				&account.CreatedAt) 
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}