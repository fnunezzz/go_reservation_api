package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *ApiServer) UserHandler(router *mux.Router) {
	router.HandleFunc("/user", makeHttpHandleFunc(s.handleUser)).Methods("POST", "PUT", "DELETE")
	router.HandleFunc("/user/{id}", makeHttpHandleFunc(s.handleGetUser)).Methods("GET")
}

func (s *ApiServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "PUT" {
		return s.handleCreateUser(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteUser(w, r)
	}
	
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *ApiServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	mock := NewUser("Filipe", "nunez", 27)

	return WriteJson(w, http.StatusOK, mock)

}

func (s *ApiServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserRequest := &CreateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(createUserRequest); err != nil {
		return err
	}

	log.Println(createUserRequest)
	return nil
}

func (s *ApiServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}