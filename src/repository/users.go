package repository

import (
	"database/sql"
	"devbook/src/models"
)

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {
	statment, statmentError := repository.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if statmentError != nil {
		return 0, statmentError
	}

	defer statment.Close()
	result, resultError := statment.Exec(user.Name, user.Nick, user.Email, user.Password)
	if resultError != nil {
		return 0, resultError
	}

	lastId, errorLastId := result.LastInsertId()
	if errorLastId != nil {
		return 0, errorLastId
	}

	return uint64(lastId), nil
}
