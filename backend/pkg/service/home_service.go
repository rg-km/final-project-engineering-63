package service

import (
	"go_jwt/helper"
	"go_jwt/model/web"
	"go_jwt/pkg/repository"
)

type HomeServiceSQLite struct {
	homeRepository repository.HomeRepository
}

func NewHomeService(homeRepo repository.HomeRepository) *HomeServiceSQLite {
	return &HomeServiceSQLite{
		homeRepository: homeRepo,
	}
}

func (s *HomeServiceSQLite) GetQuizByCategoryIdWithPagination(categoryId, page, limit int32) ([]web.QuizResponse, error) {
	category, err := s.homeRepository.FindCategoryById(categoryId)
	helper.PanicIfError(err)

	quizRespResult, err := s.homeRepository.FindQuizByCategoryIdWithPagination(
		category, page, limit,
	)
	helper.PanicIfError(err)

	return helper.ToQuizResponses(quizRespResult), nil
}
