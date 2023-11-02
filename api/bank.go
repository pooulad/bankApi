package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddr string
}

func NewApiServer(listenAddr string) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
	}
}

func (a *ApiServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", MakeHttpHandleFunc(a.handleAccount))
}

func (a *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
