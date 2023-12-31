package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pooulad/bankApi/database"
	"github.com/pooulad/bankApi/model"
	"github.com/pooulad/bankApi/util"
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

	router.HandleFunc("/login", MakeHttpHandleFunc(a.handleLogin))
	router.HandleFunc("/account", MakeHttpHandleFunc(a.handleAccount))
	router.HandleFunc("/account/{id}", withJwtAuth(MakeHttpHandleFunc(a.handleAccountById), a.store))
	router.HandleFunc("/transfer", MakeHttpHandleFunc(a.handleTransfer))

	log.Println("JSON api server running on port:", a.listenAddr)

	http.ListenAndServe(a.listenAddr, router)
}

func (a *ApiServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("method not allowd %s", r.Method)
	}

	var loginRequest model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		return err
	}

	acc, err := a.store.GetAccountByNumber(int(loginRequest.Number))
	if err != nil {
		return err
	}

	if !acc.ValidPassword(loginRequest.Password) {
		return fmt.Errorf("not authenticated")
	}

	token, err := createJwt(acc)
	if err != nil {
		return err
	}

	loginResponse := model.LoginResponse{
		Token:  token,
		Number: acc.Number,
	}

	return WriteJson(w, http.StatusOK, loginResponse)
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

func (a *ApiServer) handleAccountById(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return a.handleGetAccountById(w, r)
	}
	if r.Method == "DELETE" {
		return a.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowd %s", r.Method)
}

func (a *ApiServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	id, err := util.GetAccountParameterId(r)
	if err != nil {
		return err
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
	defer r.Body.Close()

	account, err := model.NewAccount(createAccountReq.FirstName, createAccountReq.LastName, createAccountReq.Password)
	if err != nil {
		return err
	}

	err = a.store.CreateAccount(account)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, account)
}

func (a *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := util.GetAccountParameterId(r)
	if err != nil {
		return err
	}

	err = a.store.DeleteAccount(id)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusNoContent, map[string]int{"deleted": id})
}

func (a *ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	transferReq := new(model.TransferRequest)

	err := json.NewDecoder(r.Body).Decode(transferReq)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	// account := model.NewAccount(transferReq.FirstName, transferReq.LastName)

	// err = a.store.CreateAccount(account)
	// if err != nil {
	// 	return err
	// }

	return WriteJson(w, http.StatusOK, transferReq)
}
