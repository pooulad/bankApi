package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pooulad/bankApi/database"
	"github.com/pooulad/bankApi/model"
)

type ApiServer struct {
	listenAddr string
	store      database.Storage
}

func NewApiServer(listenAddr string, store database.Storage) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (a *ApiServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", MakeHttpHandleFunc(a.handleAccount))
	router.HandleFunc("/account/{id}", MakeHttpHandleFunc(a.handleGetAccountById))

	log.Println("JSON api server running on port:", a.listenAddr)

	http.ListenAndServe(a.listenAddr, router)
}

func (a *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return a.handleGetAccounts(w, r)
	}
	if r.Method == "POST" {
		return a.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return a.handleDeleteAccount(w, r)
	}

	return nil
}

func (a *ApiServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := a.store.GetAccounts()
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, accounts)
}

func (a *ApiServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid id given in url:%d", id)
	}

	account, err := a.store.GetAccountById(id)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, account)
}

func (a *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountReq := new(model.CreateAccountRequest)

	err := json.NewDecoder(r.Body).Decode(createAccountReq)
	if err != nil {
		return err
	}

	account := model.NewAccount(createAccountReq.FirstName, createAccountReq.LastName)

	err = a.store.CreateAccount(account)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, account)
}

func (a *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
