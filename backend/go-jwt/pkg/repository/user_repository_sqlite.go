package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_jwt/helper"
	"go_jwt/model/domain"
)

type UserRepositorySQLite struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositorySQLite {
	return &UserRepositorySQLite{
		DB: db,
	}
}

func (repo *UserRepositorySQLite) Save(ctx context.Context, user domain.UserDomain) (domain.UserDomain, error) {
	query := `
		INSERT INTO users (username, email, password, role, is_login)
		VALUES (?, ?, ?, ?, ?)
	`

	done := make(chan bool)
	go func(ch chan<- bool) error {
		defer close(ch)
		result, err := repo.DB.ExecContext(
			ctx, query,
			user.Username, user.Email, user.Password,
			user.Role, user.IsLogin,
		)
		if err != nil {
			ch <- false
			return errors.New("error when saving user")
		}

		id, errId := result.LastInsertId()
		if errId != nil {
			ch <- false
			return errors.New("error getting last insert id")
		}

		user.Id = uint32(id)
		ch <- true

		return nil
	}(done)

	if helper.OK(done) {
		return user, nil
	}

	return domain.UserDomain{}, errors.New("error when saving user")
}

func (repo *UserRepositorySQLite) GetUser(ctx context.Context, email, password string) (domain.UserDomain, error) {
	querySelect := `
		SELECT id, username, email, password, role, is_login
		FROM users
		WHERE email = ? AND password = ?
	`
	var user domain.UserDomain
	stmtSelect, err := repo.DB.PrepareContext(ctx, querySelect)
	if err != nil {
		return user, err
	}
	defer stmtSelect.Close()

	err = stmtSelect.QueryRowContext(ctx, email, password).Scan(
		&user.Id, &user.Username, &user.Email, &user.Password,
		&user.Role, &user.IsLogin,
	)
	if err != nil {
		return user, err
	}

	queryUpdate := `UPDATE users SET is_login = true WHERE email = ?`
	stmtUpdate, err := repo.DB.PrepareContext(ctx, queryUpdate)
	if err != nil {
		return user, err
	}
	defer stmtUpdate.Close()

	_, err = stmtUpdate.ExecContext(ctx, email)
	if err != nil {
		return user, err
	}

	if user.IsLogin {
		return user, errors.New("user already logged in")
	}

	return user, nil
}

func (repo *UserRepositorySQLite) Logout(ctx context.Context, username string) (bool, error) {
	query := `UPDATE users SET is_login = false WHERE username = ?`

	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		_, err := repo.DB.ExecContext(ctx, query, username)
		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)

	if helper.OK(done) {
		return true, nil
	}

	return false, errors.New("error when logout")
}
