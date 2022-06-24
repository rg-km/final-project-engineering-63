package repository

import (
	"context"
	"database/sql"
	"go_jwt/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.UserDomain) (domain.UserDomain, error)
	GetUser(ctx context.Context, email, password string) (domain.UserDomain, error)
	Logout(ctx context.Context, username string) (bool, error)
}

type Repository struct {
	UserRepository
}
type Quiz struct {
	ID          int64  `db:"id"`
	Category    string `db:"category"`
	QuizName	string `db:"quiz_name"`
}

type QuizItem struct {
	ID          int64  `db:"id"`
	Category    string `db:"category"`
	QuizID   	int64  `db:"product_id"`
	QuizName	string `db:"product_name"`
	Quantity    int    `db:"quantity"`
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(db),
	}
}
