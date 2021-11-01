package repository

import (
	"database/sql"
	"devbook/src/models"
	"fmt"
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

func (repository Users) GetAll(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)
	rows, requestError := repository.db.Query("SELECT id, name, email, nick, created_at FROM users WHERE name LIKE ? OR nick LIKE ?", nameOrNick, nameOrNick)
	if requestError != nil {
		return nil, requestError
	}

	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		if requestError = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Nick,
			&user.CreatedAt,
		); requestError != nil {
			return nil, requestError
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) GetBy(userID uint64) (models.User, error) {
	rows, requestError := repository.db.Query("SELECT id, name, email, nick, created_at FROM users WHERE id = ?", userID)
	if requestError != nil {
		return models.User{}, requestError
	}

	defer rows.Close()
	var user models.User
	if rows.Next() {
		if requestError = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Nick,
			&user.CreatedAt,
		); requestError != nil {
			return models.User{}, requestError
		}
	}

	return user, nil
}

func (repository Users) Update(userId uint64, user models.User) error {
	statment, statmentError := repository.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")
	if statmentError != nil {
		return statmentError
	}

	defer statment.Close()
	_, resultError := statment.Exec(user.Name, user.Nick, user.Email, userId)
	if resultError != nil {
		return resultError
	}

	return nil
}

func (repository Users) Delete(userId uint64) error {
	statment, statmentError := repository.db.Prepare("DELETE FROM users WHERE id = ?")
	if statmentError != nil {
		return statmentError
	}

	defer statment.Close()
	_, resultError := statment.Exec(userId)
	if resultError != nil {
		return resultError
	}

	return nil
}

func (repository Users) GetByEmail(userEmail string) (models.User, error) {
	rows, requestError := repository.db.Query("SELECT id, password FROM users WHERE email = ?", userEmail)
	if requestError != nil {
		return models.User{}, requestError
	}

	defer rows.Close()
	var user models.User
	if rows.Next() {
		if requestError = rows.Scan(
			&user.ID,
			&user.Password,
		); requestError != nil {
			return models.User{}, requestError
		}
	}

	return user, nil
}

func (repository Users) Follow(userID, followID uint64) error {
	statment, statmentError := repository.db.Prepare("INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)")
	if statmentError != nil {
		return statmentError
	}

	defer statment.Close()
	_, resultError := statment.Exec(userID, followID)
	if resultError != nil {
		return resultError
	}

	return nil
}

func (repository Users) UnFollow(userID, followID uint64) error {
	statment, statmentError := repository.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")
	if statmentError != nil {
		return statmentError
	}

	defer statment.Close()
	_, resultError := statment.Exec(userID, followID)
	if resultError != nil {
		return resultError
	}

	return nil
}

func (repository Users) GetFollowers(userID uint64) ([]models.User, error) {
	rows, requestError := repository.db.Query("SELECT u.id, u.name, u.nick, u.email, u.created_at FROM users u INNER JOIN followers s ON u.id = s.follower_id WHERE s.user_id = ?", userID)
	if requestError != nil {
		return nil, requestError
	}

	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		if requestError = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); requestError != nil {
			return nil, requestError
		}

		users = append(users, user)
	}
	return users, nil
}

func (repository Users) GetFollowing(userID uint64) ([]models.User, error) {
	rows, requestError := repository.db.Query("SELECT u.id, u.name, u.nick, u.email, u.created_at FROM users u INNER JOIN followers s ON u.id = s.user_id WHERE s.follower_id = ?", userID)
	if requestError != nil {
		return nil, requestError
	}

	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		if requestError = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); requestError != nil {
			return nil, requestError
		}

		users = append(users, user)
	}
	return users, nil
}

func (repository Users) GetPassword(userID uint64) (string, error) {
	rows, requestError := repository.db.Query("SELECT password FROM users WHERE id = ?", userID)
	if requestError != nil {
		return "", requestError
	}

	defer rows.Close()
	var user models.User
	if rows.Next() {
		if requestError = rows.Scan(&user.Password); requestError != nil {
			return "", requestError
		}
	}

	return user.Password, nil
}

func (repository Users) UpdatePassword(userID uint64, password string) error {
	statment, statmentError := repository.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if statmentError != nil {
		return statmentError
	}

	defer statment.Close()
	_, resultError := statment.Exec(password, userID)
	if resultError != nil {
		return resultError
	}

	return nil
}
