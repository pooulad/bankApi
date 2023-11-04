package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pooulad/bankApi/config"
	"github.com/pooulad/bankApi/model"
)

type Storage interface {
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

func (s *PostgresStore) CreateAccount(*model.Account) error {
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
