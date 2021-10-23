package database

import (
	"database/sql"
	"devbook/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, connectionError := sql.Open("mysql", config.StringConnection)
	if connectionError != nil {
		return nil, connectionError
	}

	if connectionError = db.Ping(); connectionError != nil {
		db.Close()
		return nil, connectionError
	}

	return db, nil
}
