package main

import (
	"context"
	"fmt"
)

func (s *Database) GetUser(email string, password string) (*User, error) {

	sql := `
		select id, first_name, last_name, password from users where email = $1
	`
	user := User{}

	err := s.conn.QueryRow(
		context.Background(), 
		sql, 
		email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password)

	if err != nil {
		return nil, fmt.Errorf("email or password are wrong")
	}


	if valid := CheckPasswordHash(password, user.Password); !valid {
		return nil, fmt.Errorf("email or password are wrong")
	}


	return &user, nil
}

func (s *Database) CreateUser(user *User) error {

	err := s.checkIfUserExistsByCpf(user.CPF)

	if err != nil {
		return err
	}

	err = s.checkIfUserExistsByEmail(user.Email)

	if err != nil {
		return err
	}

	sql := `
	insert into users 
	(id, first_name, last_name, age, cpf, email, password, created_at) 
	values 
	($1, $2, $3, $4, $5, $6, $7, now())
	`	
	_, err = s.conn.Query(
		context.Background(), 
		sql, 
		user.ID, 
		user.FirstName, 
		user.LastName, 
		user.Age, 
		user.CPF, 
		user.Email, 
		user.Password)

	if err != nil {
		return err
	}

	return nil
}


func (s *Database) checkIfUserExistsByCpf(cpf string) error {

	sql := `
	select id from users where cpf = $1
	`	

	user := User{}

	err := s.conn.QueryRow(
		context.Background(), 
		sql, 
		cpf).Scan(&user.ID)

	if err == nil {
		return fmt.Errorf("user already exists")
	}

	return nil
}


func (s *Database) checkIfUserExistsByEmail(email string) error {

	sql := `
	select id from users where email = $1
	`	

	user := User{}

	err := s.conn.QueryRow(
		context.Background(), 
		sql, 
		email).Scan(&user.ID)

	if err == nil {
		return fmt.Errorf("user already exists")
	}

	return nil
}