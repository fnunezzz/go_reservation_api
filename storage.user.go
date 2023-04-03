package main

import (
	"context"
	"fmt"
)

func (s *Database) CreateUser(user *User) error {

	err := s.getUserByCPF(user.CPF)

	if err != nil {
		return err
	}

	err = s.getUserByEmail(user.Email)

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


func (s *Database) getUserByCPF(cpf string) error {

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


func (s *Database) getUserByEmail(email string) error {

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