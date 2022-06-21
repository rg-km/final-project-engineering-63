package service

import (
	"context"
	"go_jwt/model/web"
	"go_jwt/pkg/repository"
)

type UserService interface {
	Create(ctx context.Context, request web.RegisterCreateRequest) (web.RegisterResponse, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (string, error)
	Logout(ctx context.Context, username string) (bool, error)
}

type Service struct {
	UserService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService: NewAuthService(repo.UserRepository),
	}
}
