package main

import "github.com/google/uuid"

type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age int `json:"age"`
	CPF string `json:"cpf"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID string `json:"id,omitempty"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age int `json:"age,omitempty"`
	CPF string `json:"cpf,omitempty"`
	Email string `json:"email"`
	Password string `json:"password,omitempty"`
}

func NewUser(firstName string, LastName string, age int, cpf string, email string, password string) *User {
	return &User{
		ID: uuid.NewString(),
		FirstName: firstName,
		LastName: LastName,
		Age: age,
		CPF: cpf,
		Email: email,
		Password: password,
	}
}