package service

import (
	"context"
	"go_jwt/model/web"
	"go_jwt/pkg/repository"
)

type AuthService interface {
	Create(request web.RegisterCreateRequest, email string) (web.RegisterResponse, error)
	GenerateToken(email, password string) (web.LoginResponse, error)
	ParseToken(ctx context.Context, token string) (int32, string, error)
}

type DashboardService interface {
	CreateQuiz(request web.QuizRequest) (web.QuizResponse, error)
	GetAdminQuestions() ([]web.QuizResponse, error)
	DeleteAdminQuestionById(questionId int32) (bool, error)
}

type HomeService interface {
	GetQuizByCategoryIdWithPagination(categoryId, page, limit int32) ([]web.QuizResponse, error)
}

type Service struct {
	AuthService
	DashboardService
	HomeService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthService: NewAuthService(repo.AuthRepository),
		DashboardService: NewDashboardService(
			repo.DashboardRepository,
			repo.HomeRepository,
		),
		HomeService: NewHomeService(repo.HomeRepository),
	}
}
