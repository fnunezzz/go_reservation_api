package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage interface {
	CreateUser(*User) error
}


type Postgres struct {
	conn *pgxpool.Pool
}


func newPostgresConn() (*Postgres, error) {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	if err := dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}


	defer dbpool.Close()
	var greeting string

	fmt.Println(greeting)

	return &Postgres{
		conn: dbpool,
	}, nil
}


func (s *Postgres) CreateUser(*User) error {

	return nil
}