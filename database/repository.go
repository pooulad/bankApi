package database

import (
	"fmt"

	"github.com/pooulad/bankApi/model"
	"github.com/pooulad/bankApi/util"
)

type Storage interface {
	GetAccounts() ([]*model.Account, error)
	CreateAccount(*model.Account) error
	DeleteAccount(int) error
	UpdateAccount(*model.Account) error
	GetAccountById(int) (*model.Account, error)
	GetAccountByNumber(int) (*model.Account, error)
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(100),
		last_name varchar(100),
		number serial,
		encrypted_password varchar(100),
		balance serial,
		created_at timestamp
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(acc *model.Account) error {
	query := `insert into account
		(first_name,last_name,number,encrypted_password,balance,created_at)
		values ($1,$2,$3,$4,$5,$6)`

	resp, err := s.db.Query(query, acc.FirstName, acc.LastName, acc.Number, acc.EncryptedPassword, acc.Balance, acc.CreatedAt)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

func (s *PostgresStore) UpdateAccount(*model.Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Query("DELETE FROM account WHERE id = $1", id)
	return err
}

func (s *PostgresStore) GetAccountByNumber(number int) (*model.Account, error) {
	rows, err := s.db.Query("SELECT * FROM account WHERE number = $1", number)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return util.ScanSqlRows(rows)
	}
	return nil, fmt.Errorf("account with number [%d] not found", number)
}
func (s *PostgresStore) GetAccountById(id int) (*model.Account, error) {
	rows, err := s.db.Query("SELECT * FROM account WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return util.ScanSqlRows(rows)
	}
	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) GetAccounts() ([]*model.Account, error) {
	rows, err := s.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}

	accounts := []*model.Account{}

	for rows.Next() {
		account, err := util.ScanSqlRows(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	if err = rows.Err(); err != nil {
		return accounts, err
	}

	return accounts, nil

}
