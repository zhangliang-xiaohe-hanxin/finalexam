package db 

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var Session *sql.DB


func CreateDB(dbName string) error {
	
	session, err := sql.Open("postgres", dbName)
	if err != nil {
		log.Fatal("can't open", err.Error())
	}

	defer session.Close()
	
	
	stmt, err := session.Prepare(
		`CREATE TABLE IF NOT EXISTS customers(
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}