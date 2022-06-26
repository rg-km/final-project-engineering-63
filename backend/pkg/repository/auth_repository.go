package repository

import (
	"database/sql"
	"errors"
	"go_jwt/helper"
	"go_jwt/model/domain"
)

type AuthRepositorySQLite struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepositorySQLite {
	return &AuthRepositorySQLite{
		DB: db,
	}
}

func (repo *AuthRepositorySQLite) Save(userParam domain.UserDomain, email string) (domain.UserDomain, error) {
	var user domain.UserDomain

	query := `SELECT id, email FROM users WHERE email = ?`
	row := repo.DB.QueryRow(query, email)
	err := row.Scan(&user.Id, &user.Email)
	if err == nil {
		return user, errors.New("email already exists")
	}

	query = `
	INSERT INTO users (username, email, password, role) 
	VALUES (?, ?, ?, ?)`

	result, err := repo.DB.Exec(query,
		userParam.Username, email, userParam.Password, userParam.Role,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int32(id)

	return userParam, nil
}

func (repo *AuthRepositorySQLite) FindUser(email, password string) (domain.UserDomain, error) {
	query := `
		SELECT id, username, email, password, role
		FROM users
		WHERE email = ? AND password = ?
	`
	var user domain.UserDomain

	row := repo.DB.QueryRow(query, email, password)
	err := row.Scan(
		&user.Id, &user.Username, &user.Email,
		&user.Password, &user.Role,
	)
	if err != nil {
		return user, errors.New("login failed")
	}

	return user, nil
}
