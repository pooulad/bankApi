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
	GetAccountById(int) (*model.Account, error)
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