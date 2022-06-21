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

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(db),
	}
}
