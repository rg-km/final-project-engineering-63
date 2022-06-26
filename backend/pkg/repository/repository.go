package repository

import (
	"database/sql"
	"go_jwt/model/domain"
	"go_jwt/model/web"
)

type AuthRepository interface {
	Save(user domain.UserDomain, email string) (domain.UserDomain, error)
	FindUser(email, password string) (domain.UserDomain, error)
}

type DashboardRepository interface {
	SaveQuiz(quizDomain domain.QuizDomain, quizReq web.QuizRequest) (domain.QuizDomain, error)
	FindAllAdminQuestions() ([]web.QuizResponse, error)
	DeleteAdminQuestionById(questionId int32) (bool, error)
}

type HomeRepository interface {
	FindCategoryByName(name string) (domain.CategoryDomain, error)
	FindCategoryById(categoryId int32) (domain.CategoryDomain, error)
	FindQuizByCategoryIdWithPagination(category domain.CategoryDomain, page, limit int32) ([]web.QuizResponse, error)
}

type Repository struct {
	AuthRepository
	DashboardRepository
	HomeRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		AuthRepository:      NewAuthRepository(db),
		DashboardRepository: NewDashboardRepository(db),
		HomeRepository:      NewHomeRepository(db),
	}
}
