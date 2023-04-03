package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *ApiServer) UserHandler(router *mux.Router) {
	router.HandleFunc("/user", makeHttpHandleFunc(s.handleUser)).Methods("POST", "PUT", "DELETE")
	router.HandleFunc("/user/{id}", makeHttpHandleFunc(s.handleGetUser)).Methods("GET")
}

func (s *ApiServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleLoginUser(w, r)
	}
	if r.Method == "PUT" {
		return s.handleCreateUser(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteUser(w, r)
	}
	
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *ApiServer) handleLoginUser(w http.ResponseWriter, r *http.Request) error {
	loginUserRequest := &LoginUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(loginUserRequest); err != nil {
		return err
	}


	user, err := s.storage.GetUser(loginUserRequest.Email, loginUserRequest.Password);

	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, &UserResponse{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName})

}

func (s *ApiServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	fmt.Println(id)

	return WriteJson(w, http.StatusCreated, &UserResponse{ID: id}) // mock

}

func (s *ApiServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserRequest := &CreateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(createUserRequest); err != nil {
		return err
	}


	user, err := NewUser(createUserRequest.FirstName, createUserRequest.LastName, createUserRequest.Age, createUserRequest.CPF, createUserRequest.Email, createUserRequest.Password)


	if err != nil {
		return err
	}


	if err := s.storage.CreateUser(user); err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, &UserResponse{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName})
}

func (s *ApiServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}