package util

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pooulad/bankApi/model"
)

func ScanSqlRows(rows *sql.Rows) (*model.Account, error) {
	account := new(model.Account)
	err := rows.Scan(&account.ID, &account.FirstName, &account.LastName,
		&account.Number, &account.Balance, &account.CreatedAt)

	return account, err
}

func GetAccountParameterId(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given in url:%d", id)
	}

	return id, nil
}