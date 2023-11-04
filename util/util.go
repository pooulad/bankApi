package util

import (
	"database/sql"

	"github.com/pooulad/bankApi/model"
)

func ScanSqlRows(rows *sql.Rows) (*model.Account, error) {
	account := new(model.Account)
	err := rows.Scan(&account.ID, &account.FirstName, &account.LastName,
		&account.Number, &account.Balance, &account.CreatedAt)

	return account, err
}
