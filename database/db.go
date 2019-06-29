package db 

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"context"
	"log"
)

func GetSession(c context.Context) (*sql.DB, error) {
	session, ok := c.Value("session").(*sql.DB)
	if !ok {
		return nil, errors.New("cannot get session")
	}
	return session, nil
}

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
		log.Println("----------create failed", err)
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println("----------create failed", err)
		return err
	}
	log.Println("----------create Success")
	return nil
}