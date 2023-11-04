package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pooulad/bankApi/config"
	"github.com/pooulad/bankApi/model"
)

type Storage interface {
	GetAccounts([]*model.Account) error
	CreateAccount(*model.Account) error
	DeleteAccount(int) error
	UpdateAccount(*model.Account) error
	GetAccountById(int) error
}

type PostgresStore struct {
	db *sql.DB
}

func ConnectDB(cfg *config.PostgresConfig) (*PostgresStore, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Name, cfg.SSL)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("postgres connected")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial,
		created_at timestamp
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(acc *model.Account) error {
	query := `insert into account
		(first_name,last_name,number,balance,created_at)
		values ($1,$2,$3,$4,$5)`

	resp, err := s.db.Query(query, acc.FirstName, acc.LastName, acc.Number, acc.Balance, acc.CreatedAt)
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
	return nil
}

func (s *PostgresStore) GetAccountById(id int) error {
	return nil
}
