package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage interface {
	CreateUser(*User) error
}


// Don't plan on using anything but PostgresSQL, but in case multiple databases becomes a feature pgx docs have it: https://github.com/jackc/pgx/wiki/Getting-started-with-pgx-through-database-sql#getting-started-with-pgx-through-databasesql
type Database struct {
	conn *pgxpool.Pool
}


func newPostgresConn() (*Database, error) {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	if err := dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &Database{
		conn: dbpool,
	}, nil
}
