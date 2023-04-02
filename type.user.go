package main

import "github.com/google/uuid"

type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age int `json:"age"`
	CPF string `json:"cpf"`
}

type User struct {
	ID string `json:"id,omitempty"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age int `json:"age,omitempty"`
	CPF string `json:"cpf,omitempty"`
}

func NewUser(firstName string, LastName string, age int) *User {
	return &User{
		ID: uuid.NewString(),
		FirstName: firstName,
		LastName: LastName,
		Age: age,
	}
}