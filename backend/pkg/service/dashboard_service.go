package service

import (
	"go_jwt/helper"
	"go_jwt/model/domain"
	"go_jwt/model/web"
	"go_jwt/pkg/repository"
)

type DashboardServiceSQLite struct {
	dashboardRepository repository.DashboardRepository
	homeRepository      repository.HomeRepository
}

func NewDashboardService(dashboardRepo repository.DashboardRepository, homeRepo repository.HomeRepository) *DashboardServiceSQLite {
	return &DashboardServiceSQLite{
		dashboardRepository: dashboardRepo,
		homeRepository:      homeRepo,
	}
}

func (s *DashboardServiceSQLite) CreateQuiz(quizReq web.QuizRequest) (web.QuizResponse, error) {
	category, err := s.homeRepository.FindCategoryByName(quizReq.Category)
	helper.PanicIfError(err)

	quizDomain := domain.QuizDomain{
		Question:   quizReq.Question,
		CategoryId: category.Id,
	}

	quizDomainResult, err := s.dashboardRepository.SaveQuiz(
		quizDomain, quizReq,
	)
	helper.PanicIfError(err)

	return helper.ToQuizResponse(quizDomainResult, category, quizReq), nil
}

func (s *DashboardServiceSQLite) GetAdminQuestions() ([]web.QuizResponse, error) {
	quizResponse, err := s.dashboardRepository.FindAllAdminQuestions()
	helper.PanicIfError(err)

	return helper.ToQuizResponses(quizResponse), nil
}

func (s *DashboardServiceSQLite) DeleteAdminQuestionById(questionId int32) (bool, error) {
	return s.dashboardRepository.DeleteAdminQuestionById(questionId)
}
