package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)


func (s *Database) Init() error {

	err := s.createMigrationsTable()

	if err != nil {
		return err
	}

	err = s.runMigrations()

	if err != nil {
		return err
	}

	return nil

}


func (s *Database) createMigrationsTable() error {

	query := `SELECT EXISTS (
		SELECT FROM 
			pg_tables
		WHERE 
			schemaname = 'public' AND 
			tablename  = 'migrations'
		)`

	var exists bool
	err := s.conn.QueryRow(context.Background(), query).Scan(&exists)
	if (err != nil || exists) {
		log.Println("Migrations table exists")
		return nil
	}

	query = `create table if not exists migrations (
		id varchar,
		migration_name varchar not null primary key,
		created_at timestamp default now()
	)`

	commandTag, err := s.conn.Exec(context.Background(), query)

	if err != nil {
		log.Fatalln("Error creating migrations table")
		return err
	}

	if commandTag.RowsAffected() == 0 {
		log.Println("Migrations table created")
	}

	return nil

}

// TODO
// For now reading from files is OK. But in the future they should be able to be rollbacked aswell
func (s *Database) runMigrations() error {
	path := filepath.Join(os.Getenv("MIGRATIONS_PATH"))

	dir, ioErr := os.ReadDir(path)

	if ioErr != nil {
		log.Fatalln("Migrations folder path not found")
		return ioErr
	 }

	 migrationsRanCounter := 0

	 for _, file := range dir {

		query := fmt.Sprintf(`select exists (select from migrations where migration_name = '%s')`, file.Name())

		var exists bool
		err := s.conn.QueryRow(context.Background(), query).Scan(&exists)
		if (err != nil || exists) {
			log.Println("Migration", file.Name(), "already exists")
			continue
		}

		c, ioErr := os.ReadFile(path+"/"+file.Name())
		if ioErr != nil {
			log.Fatalln(ioErr)
			log.Fatalln("Couldn't read migration file", file.Name())
			return ioErr
		}

		sql	:= string(c)
		_, err = s.conn.Exec(context.Background(), sql)
		if err != nil {
			log.Fatalln("Couldn't run migration file", file.Name())
			return ioErr
		}
		
		insert := fmt.Sprintf(`insert into migrations (id, migration_name, created_at) values ('%s', '%s', now())`, uuid.NewString(), file.Name())
		
		_, err = s.conn.Exec(context.Background(), insert)

		if err != nil {
			log.Fatalln(err)
			return err
		}

		log.Println("Migration", file.Name(), "done")
		migrationsRanCounter++

	 }

	 log.Println("Total migrations ran:", migrationsRanCounter)
	 return nil
}