package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/paemuri/brdoc"
)

type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age int `json:"age"`
	CPF string `json:"cpf"`
	Email string `json:"email"`
	Password string `json:"password"`
}


type UserResponse struct {
	ID string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

type User struct {
	ID string `json:"id,omitempty"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age int `json:"age,omitempty"`
	CPF string `json:"cpf,omitempty"`
	Email string `json:"email"`
	Password string `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func NewUser(firstName string, LastName string, age int, cpf string, email string, password string) (*User, error) {

	if firstName == "" {
		return nil, fmt.Errorf("first name must be informed")
	}

	if LastName == "" {
		return nil, fmt.Errorf("last name must be informed")
	}

	if age < 18 {
		return nil, fmt.Errorf("age must be informed and be 18 or above")
	}

	if cpf == "" {
		return nil, fmt.Errorf("CPF must be informed")
	} else if !brdoc.IsCPF(cpf) {
		return nil, fmt.Errorf("CPF must be valid")
	}
	
	if email == "" {
		return nil, fmt.Errorf("email must be informed")
	}
	
	if password == "" {
		return nil, fmt.Errorf("password must be informed")
	} else if err := CheckPasswordStrength(password); err != nil {
		return nil, err
	}

	var hash string 

	if h, err := HashPassword(password); err != nil {
		return nil, fmt.Errorf("error saving user")
	} else {
		hash = h
	}

	return &User{
		ID: uuid.NewString(),
		FirstName: strings.ToUpper(firstName),
		LastName: strings.ToUpper(LastName),
		Age: age,
		CPF: cpf,
		Email: email,
		Password: hash,
	}, nil
}

func LoadUser(firstName string, LastName string, age int, cpf string, email string, password string, createdAt time.Time) *User {
	return &User{
		ID: uuid.NewString(),
		FirstName: firstName,
		LastName: LastName,
		Age: age,
		CPF: cpf,
		Email: email,
		Password: password,
		CreatedAt: createdAt,
	}
}